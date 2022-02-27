package rollingnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBucket_Increase(t *testing.T) {
	b := NewBucket(time.Now(), time.Second)
	assert.Equal(t, int64(0), b.Value())
	b.Increase(3)
	assert.Equal(t, int64(3), b.Value())
	b.Increase(100)
	assert.Equal(t, int64(103), b.Value())
}

func TestBucket_Match(t *testing.T) {
	now := time.Now()
	duration := time.Second
	b := NewBucket(time.Now(), duration)
	assert.False(t, b.Match(now.Add(-time.Millisecond)))         // 早于开始时间的，视为不匹配
	assert.True(t, b.Match(now))                                 // 开始时间，视为匹配
	assert.True(t, b.Match(now.Add(duration-time.Millisecond)))  // 早于结束时间，视为匹配
	assert.False(t, b.Match(now.Add(duration)))                  // 等于结束时间的，视为不匹配
	assert.False(t, b.Match(now.Add(duration+time.Millisecond))) // 晚于结束时间的，视为不匹配
}

func TestBucket_Compare(t *testing.T) {
	now := time.Now()
	duration := time.Second
	b := NewBucket(time.Now(), duration)
	assert.Equal(t, BeforeWindow, b.Compare(now.Add(-time.Millisecond)))
	assert.Equal(t, MatchWindow, b.Compare(now))
	assert.Equal(t, MatchWindow, b.Compare(now.Add(duration-time.Millisecond)))
	assert.Equal(t, AfterWindow, b.Compare(now.Add(duration)))
	assert.Equal(t, AfterWindow, b.Compare(now.Add(duration+time.Millisecond)))
}
