# ftpd

[中文](README_CN.md)

A FTP server based on [gitea.com/goftp/server](http://gitea.com/goftp/server).

Full documentation for the package is available on [godoc](http://godoc.org/goftp.io/ftpd)

# Installation

## From binary releases

Download the binaries from [https://gitea.com/goftp/ftpd/releases](https://gitea.com/goftp/ftpd/releases).

You can also build the binary yourself. After you clone the repository,

    go generate ./...
    go build -tags=bindata -mod=vendor

## From Source

    go get goftp.io/ftpd

Then run it:

    $GOPATH/bin/ftpd

And finally, connect to the server with any FTP client and the following
details:

    host: 127.0.0.1
    port: 2121
    username: admin
    password: 123456

More features, you can copy config.ini to the ftpd directory and modify it.