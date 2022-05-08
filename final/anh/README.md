# anh

## 术语

telA, telX, telB = numbers in AXB-binding mode. xbr = x-binding-register

## 启动方式

进入`anh`根目录，依次执行下面的指令。

```shell
# 启动xbr-service
go run cmd/service/xbr/main.go -app_config=configs/default/service/xbr/config.ini -log_config=configs/default/service/xbr/seelog.xml

# 启动user-interface
go run cmd/interface/user/main.go -app_config=configs/default/interface/user/config.ini -log_config=configs/default/interface/user/seelog.xml
```