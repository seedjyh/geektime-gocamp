package rollingnumber

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRollingNumber_Increase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	bucketSize := time.Millisecond * 300
	bucketCount := 3
	clock := NewMockClock(mockCtrl)
	startTime := time.Date(1984, 1, 22, 0, 0, 0, 0, time.Local)
	r := newRollingNumberWithClock(bucketSize, bucketCount, clock)
	// 全新的时间窗口
	clock.EXPECT().Now().Return(startTime).Times(1)
	assert.Equal(t, int64(0), r.SumUp())
	// 一个桶
	clock.EXPECT().Now().Return(startTime).Times(2)
	r.Increase(3)
	assert.Equal(t, int64(3), r.SumUp())
	clock.EXPECT().Now().Return(startTime).Times(2)
	r.Increase(5)
	assert.Equal(t, int64(8), r.SumUp())
	// 时间过了1个窗口
	clock.EXPECT().Now().Return(startTime.Add(bucketSize)).Times(2)
	r.Increase(7)
	assert.Equal(t, int64(15), r.SumUp())
	// 时间过了所有窗口范围
	clock.EXPECT().Now().Return(startTime.Add(bucketSize * time.Duration(bucketCount*4))).Times(1)
	assert.Equal(t, int64(0), r.SumUp())
}
