# test-seelog

## 学到的知识

- 配置文件的格式串末尾，需要`%n`作为换行符。
- 格式化函数很像中间件，即「返回函数的函数」。
- seelog.CloneLogger 出来的logger调用Info等函数后，需要对clone出来的logger本身进行flush。对原始logger进行flush是没用的。

## 疑难点

- set context是否会导致协程安全问题？即不同的协程同时打印。（这一点似乎github上有注明）。
- 怎么能方便地打印参数dict、error、sessionID等信息呢。
- 怎么让代码不直接依赖特定中间件呢。（或许，用string作为参数传递给一个package？）
- 怎么让代码不直接依赖特定中间件呢。（或许，用string作为参数传递给一个package？）
