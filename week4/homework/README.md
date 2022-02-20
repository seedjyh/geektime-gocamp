# Week 4 Homework

Go进阶训练营（第7期）第4周作业。

## 介绍

这是一个学生信息管理系统。提供了两个功能：

- 添加学生信息。
- 删除学生信息。

## API使用方法

下面的HTTP请求使用[HTTPie](https://httpie.io)发起。

添加一个学生的信息：

```shell
# http delete 127.0.0.1:9000/students/3020611017
HTTP/1.1 200 OK
Content-Length: 31
Content-Type: application/json; charset=UTF-8
Date: Sun, 20 Feb 2022 13:54:16 GMT

{
    "message": "success"
}
```

删除一个学生的信息：

```shell
# http delete 127.0.0.1:9000/students/3020611017
HTTP/1.1 200 OK
Content-Length: 31
Content-Type: application/json; charset=UTF-8
Date: Sun, 20 Feb 2022 13:55:05 GMT

{
    "code": 0,
    "message": "success"
}
```
