# Week 2 Homework

Go进阶训练营（第7期）第2周作业。

## Problems

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中： https://jinshuju.net/f/Om7xH6
作业提交截止时间为 1 月 30 日（周日）晚 24:00 分前。

## Answers

不应该 Wrap 这个 error 并上抛。

因为如果上抛的话，上层的包也会依赖`database/sql`包的实现细节`sql.ErrNoRows`。

其实上层只需要知道是否查询成功，如果成功了，是没有这个数据，还是有这个数据（此时返回该数据）。

为了让返回的error能保持功能的纯净性，将「数据库故障」和「不存在该数据」（可能发生的正常情况）分离，所以采用了这个接口：

```go
type StudentDAO interface {
	// FindByID 根据ID查询一个学生
	// 如果出错了，返回 (nil, false, TheError)
	// 如果没出错，但没找到数据，返回 (nil, false, nil)
	// 如果没出错，且找到了数据，返回 (TheStudent, true, nil)
	FindByID(ctx context.Context, id int) (*Student, bool, error)
}
```
