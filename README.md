# 客户端
在根文件的 client目录下
# 服务端
地址如下
>rpc/item/main.go


# etcd环境
华为云服务器，docker安装的etcd
版本信息如下
```shell
bitnami/etcd         latest    2f44853ef5e6   3 days ago     219MB
```
# 问题
```shell
2024/03/04 14:22:24.272515 default_server_handler.go:242: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:57554, error=default codec read failed: read tcp 127.0.0.1:8890->127.0.0.1:57554: wsarecv: An existing connection was forcibly closed by the remote host.
2024/03/04 14:22:24.273034 default_server_handler.go:242: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:57554, error=default codec read failed: read tcp 127.0.0.1:8890->127.0.0.1:57554: wsarecv: An existing connection was forcibly closed by the remote host.

```