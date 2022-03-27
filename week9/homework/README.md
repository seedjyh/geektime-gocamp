# Gocamp Week 9 Homework

## 作业内容

> 1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
> 2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

## 总结socket解包方式

### fix length

通过socket收发的所有消息都有固定长度。这样一来，接收者只要将收到的字节数和该长度比较，就能确定是否收到了完整的消息。

因为所有消息长度固定，如果协议里的不同类型消息所需字节数差异很大的协议，会导致浪费。

下面是长度固定为4的协议的例子。

字节下标 | 字节含义 -|- 0 | 消息类型码（0~255） 1 | 消息参数1 （0~255） 2 | 消息参数2 （0~255） 3 | 消息参数3 （0~255）

### delimiter based

以特定标识符表示消息结束。接收者需要持续收取并依次检查每个字节，直到看到结束符。

如果正文里需要用到结束符，则需要转义。

典型的 delimiter based 协议：HTTP的 header 以`\r\n\r\n`结尾。

### length field based frame decoder

在消息开头用一个域来表示整个消息或消息体的长度。

接收者先按照 fix length 收取帧头，并从帧头中解析出帧体的长度。

典型的 length field based frame decoder 例子：

字节下标 | 字节含义 -|- 0 | 帧长度LB（lower bits） 1 | 帧长度UB（upper bits）
`2` ~ `2+(UB<<8|LB)-1` | 帧体，有`(UB << 8 | LB)`个字节)

HTTP的请求体的长度由HTTP头里的`Content-Length`描述，所以也是 length field based frame decoder 类型的。

## goim 解码器

见 ./tcp-demo/internal/protocol/goim/parser.go

参考 goim/api/protocol/protocol.go
