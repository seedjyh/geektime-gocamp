# Change Log

[[TOC]]

## unreleased (2022-02-20)

### Changed

- 修改了`package dao`的接口签名，取消了用于表示「是否找到目标数据」的`bool`返回项。如果没找到返回数据，使用`code.NotFound`错误。

- 业务通用的错误码定义在`package code`。
