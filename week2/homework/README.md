# Week 2 Homework

Go进阶训练营（第7期）第2周作业。

## Problems

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中： https://jinshuju.net/f/Om7xH6
作业提交截止时间为 1 月 30 日（周日）晚 24:00 分前。

## Answers

Wrap 了这个 error 并上抛，使用全局错误`code.ErrNotFound`表示「没找到」，在errors.Wrap的时候将底层错误字符串化，从而避免高层和sql底层耦合。
