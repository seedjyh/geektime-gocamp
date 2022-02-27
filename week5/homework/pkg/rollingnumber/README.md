# RollingNumber

## 简介

`RollingNumber`是参考[Hystrix](https://github.com/Netflix/Hystrix)的 HystrixRollingNumber 实现的滑动窗口计数器。

一个`RollingNumber`只支持一种指标的采样。如果需要对多种指标进行采样，则需要创建多个`RollingNumber`。

## 使用方法

```go
# 创建一个计数器，有5个采样窗口，每个采样窗口的宽度是1分钟。即，合计5分钟。
r := NewRollingNumber(time.Minute, 5)

# 计数器增加3。这个操作会将过期的采样桶移除。
r.Increase(3)

# 计算当前采样窗口里的总数。这个操作会将过期的采样桶移除。
fmt.Println(r.SumUp())
```

## 实现方法

`RollingNumber`内置了一个固定长度的采样桶数组`buckets`，这个数组滚动使用，是为了避免可能的slice拷贝。

`RollingNumber`内部有一个读写锁控制`buckets`的操作：
- 在增加计数器时，会先尝试用**读锁**获取需要的桶并操作。如果能获取到，则直接操作得到的桶。如果发现`buckets`不符合要求，则使用**写锁**更新`buckets`再操作。
- 在统计总数时，会先用**读锁**检查`buckets`数组是否有过期数据，如果有，则使用**写锁**去除过期数据，然后再用**读锁**统计。

采样桶`Bucket`内部用原子操作`atomic.AddInt64`来增加计数。

