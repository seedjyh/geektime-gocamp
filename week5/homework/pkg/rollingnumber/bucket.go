package rollingnumber

import (
	"sync/atomic"
	"time"
)

// Bucket 是一个采样窗口
type Bucket struct {
	startTime  time.Time // 窗口开始时刻。闭区间（包括）。
	finishTime time.Time // 窗口结束时刻。开区间（不包括）。
	value      int64     // 计数器
}

// NewBucket 创建一个新的采样窗口 Bucket 。
// 采样时间为 [startTime, startTime + size) 的半闭半开区间
func NewBucket(startTime time.Time, duration time.Duration) *Bucket {
	return &Bucket{
		startTime:  startTime,
		finishTime: startTime.Add(duration),
		value:      0,
	}
}

// Increase 原子地将计数器增加v
func (b *Bucket) Increase(v int64) {
	atomic.AddInt64(&b.value, v)
}

type CompareResult int

const (
	BeforeWindow CompareResult = iota
	MatchWindow
	AfterWindow
)

func (b *Bucket) Compare(t time.Time) CompareResult {
	if t.Sub(b.startTime) < 0 { // t 早于窗口开始时间
		return BeforeWindow
	}
	if b.finishTime.Sub(t) <= 0 { // t 晚于或等于窗口结束时间
		return AfterWindow
	}
	return MatchWindow
}

// Match 检查 t 是否在该 Bucket 的范围内
func (b *Bucket) Match(t time.Time) bool {
	return b.Compare(t) == MatchWindow
}

func (b *Bucket) Value() int64 {
	return b.value
}
