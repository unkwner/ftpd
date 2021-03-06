# ftpd

[English](README.md)

这是一个基于 [gitea.com/goftp/server](http://gitea.com/goftp/server) 编写的Ftp服务器程序。

文档可以通过 [godoc](http://godoc.org/goftp.io/ftpd) 获取。

# 安装

## 二进制安装

从 [https://gitea.com/goftp/ftpd/releases](https://gitea.com/goftp/ftpd/releases) 下载二进制程序。

你也可以在克隆此仓库后，自己编译二进制程序：

    go generate ./...
    go build -tags=bindata -mod=vendor

## 源代码安装

    go get goftp.io/ftpd

然后运行

    $GOPATH/bin/ftpd

最后，通过FTP客户端连接即可：

    host: 127.0.0.1
    port: 2121
    username: admin
    password: 123456

如需要进一步修改，可以拷贝 config.ini 文件到 ftpd 目录下，然后修改其中的配置

同时，你也可以直接使用环境变量来运行 ftpd，环境变量与配置文件样例中对应，形如 ToUpper(`<section_name>_<item_name>`), 比如：

```
SERVER_PORT=2222 ./ftpd
```

即使 config.ini 存在，这也将会在 2222 端口运行一个ftp服务器。