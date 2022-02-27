// Package rollingnumber 是参考 Hystrix 的 HystrixRollingNumber 实现的滑动窗口计数器。
package rollingnumber

import (
	"sync"
	"time"
)

type Clock interface {
	Now() time.Time
}

// RollingNumber 是一个滑动时间窗口计数器。
// 循环使用固定长度的桶数组。
type RollingNumber struct {
	bucketDuration    time.Duration // 每个bucket的时间长度
	maxBucketCount    int           // 最大bucket数
	nowBucketCount    int           // 当前bucket数
	oldestBucketIndex int           // 最早的bucket的下标
	latestBucketIndex int           // 最晚的bucket的下标
	buckets           []*Bucket     // bucket数组，滚动使用
	mutex             sync.RWMutex  // 整个对象的读写锁
	clock             Clock         // 时钟接口，一般直接使用系统时钟，但测试的时候会使用其他时钟。
}

// NewRollingNumber 创建一个新的时间窗口计数器
func NewRollingNumber(bucketDuration time.Duration, maxBucketCount int) *RollingNumber {
	return newRollingNumberWithClock(bucketDuration, maxBucketCount, SystemClock{})
}

// newRollingNumberWithClock 创建一个带时钟的时间窗口计数器，主要用于测试。
func newRollingNumberWithClock(bucketDuration time.Duration, maxBucketCount int, clock Clock) *RollingNumber {
	return &RollingNumber{
		bucketDuration:    bucketDuration,
		maxBucketCount:    maxBucketCount,
		nowBucketCount:    0,
		oldestBucketIndex: 0,
		latestBucketIndex: 0,
		buckets:           make([]*Bucket, maxBucketCount),
		mutex:             sync.RWMutex{},
		clock:             clock,
	}
}

// Increase 对当前bucket增加 v 个计数
func (r *RollingNumber) Increase(v int64) {
	now := r.Now()
	if b, ok := r.getCurrentBucket(now); ok {
		b.Increase(v)
	}
}

// SumUp 统计所有窗口的所有采样值
func (r *RollingNumber) SumUp() int64 {
	now := r.Now()
	_, _ = r.getCurrentBucket(now) // 将过期的 Bucket 刷掉
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	var total int64 = 0
	for i := 0; i < r.nowBucketCount; i++ {
		index := (r.oldestBucketIndex + i) % r.maxBucketCount
		total += r.buckets[index].Value()
	}
	return total
}

// getCurrentBucket 返回符合 now 时刻的 Bucket
// 如果 now 在已有 buckets 里，会返回匹配的 Bucket 。
// 如果 now 早于最早的 Bucket 则返回 nil, false
// 如果 now 比最晚的 Bucket 更晚，会调整现有的buckets数组，确保匹配的 Bucket 存在。这个过程可能移除过于老旧的 Bucket 。
func (r *RollingNumber) getCurrentBucket(now time.Time) (*Bucket, bool) {
	// 先期望匹配的 Bucket 已存在
	if b, ok := r.tryLatestBucket(now); ok {
		return b, true
	}
	return r.assureCurrentBucket(now)
}

// tryLatestBucket 简单地试着返回最新的 Bucket 。
// 使用读锁。
// 如果匹配，则返回该 Bucket 和 true
// 否则，返回 nil, false
func (r *RollingNumber) tryLatestBucket(t time.Time) (*Bucket, bool) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.nowBucketCount == 0 {
		return nil, false
	}
	b := r.buckets[r.latestBucketIndex]
	if b.Match(t) {
		return b, true
	}
	return nil, false
}

// assureCurrentBucket 会通过滑动窗口来能得到适当的 Bucket 。
// 使用写锁。
// 但如果 now 过于老旧，则视为找不到。
func (r *RollingNumber) assureCurrentBucket(now time.Time) (*Bucket, bool) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.removeExpiredBuckets(now)
	// 如果目前 buckets 是空的
	if r.nowBucketCount == 0 {
		newBucketIndex := 0
		r.oldestBucketIndex = newBucketIndex
		r.latestBucketIndex = newBucketIndex
		r.nowBucketCount = 1
		r.buckets[newBucketIndex] = NewBucket(now, r.bucketDuration)
		return r.buckets[newBucketIndex], true
	}
	// 逐个检查现有 buckets
	// TODO: 利用 Bucket.Compare 二分，可能会更快。
	for i := 0; i < r.nowBucketCount; i++ {
		b := r.buckets[(r.oldestBucketIndex+i)%r.maxBucketCount]
		if b.Match(now) {
			return b, true
		}
	}
	// 如果现有 buckets 都太老旧
	for !r.buckets[r.latestBucketIndex].Match(now) {
		b := NewBucket(r.buckets[r.latestBucketIndex].finishTime, r.bucketDuration)
		newIndex := (r.latestBucketIndex + 1) % r.maxBucketCount
		r.nowBucketCount++
		r.latestBucketIndex = newIndex
		r.buckets[newIndex] = b
	}
	return r.buckets[r.latestBucketIndex], true
}

// removeExpiredBuckets 移除 buckets 数组里已经过期的 Bucket 。
// 这里过期的 Bucket 是指，如果保留这个 Bucket ，会导致 now 所属的 Bucket 无法进入 buckets 数组。
// 此函数必须在写锁已经锁定时调用。
func (r *RollingNumber) removeExpiredBuckets(now time.Time) {
	maxDuration := r.bucketDuration * time.Duration(r.maxBucketCount)
	for r.nowBucketCount > 0 {
		if r.buckets[r.oldestBucketIndex].startTime.Add(maxDuration).Sub(now) <= 0 {
			r.oldestBucketIndex = (r.oldestBucketIndex + 1) % r.maxBucketCount
			r.nowBucketCount--
		} else {
			break
		}
	}
}

func (r *RollingNumber) Now() time.Time {
	return r.clock.Now()
}
