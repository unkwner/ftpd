Tango [简体中文](README_CN.md)
=======================

[![Build Status](https://drone.gitea.com/api/badges/lunny/tango/status.svg)](https://drone.gitea.com/lunny/tango) [![](http://gocover.io/_badge/gitea.com/lunny/tango)](http://gocover.io/gitea.com/lunny/tango)
[![](https://goreportcard.com/badge/gitea.com/lunny/tango)](https://goreportcard.com/report/gitea.com/lunny/tango)
[![Join the chat at https://img.shields.io/discord/323705316027924491.svg](https://img.shields.io/discord/323705316027924491.svg)](https://discord.gg/7Ckxjwu)

![Tango Logo](logo.png)

Package tango is a micro & pluggable web framework for Go.

##### Current version: v0.5.0   [Version History](https://gitea.com/lunny/tango/releases)

## Getting Started

To install Tango:

    go get gitea.com/lunny/tango

A classic usage of Tango below:

```go
package main

import (
    "errors"

    "gitea.com/lunny/tango"
)

type Action struct {
    tango.JSON
}

func (Action) Get() interface{} {
    if true {
        return map[string]string{
            "say": "Hello tango!",
        }
    }
    return errors.New("something error")
}

func main() {
    t := tango.Classic()
    t.Get("/", new(Action))
    t.Run()
}
```

Then visit `http://localhost:8000` on your browser. You will get
```
{"say":"Hello tango!"}
```

If you change `true` after `if` to `false`, then you will get
```
{"err":"something error"}
```

This code will automatically convert returned map or error to a json because we has an embedded struct `tango.JSON`.

## Features

- Powerful routing & Flexible routes combinations.
- Directly integrate with existing services.
- Easy to plugin features with modular design.
- High performance dependency injection embedded.

## Middlewares

Middlewares allow you easily plugin features for your Tango applications.

There are already many [middlewares](https://gitea.com/tango) to simplify your work:

- [recovery](https://gitea.com/lunny/tango/wiki/Recovery) - recover after panic
- [compress](https://gitea.com/lunny/tango/wiki/Compress) - Gzip & Deflate compression
- [static](https://gitea.com/lunny/tango/wiki/Static) - Serves static files
- [logger](https://gitea.com/lunny/tango/wiki/Logger) - Log the request & inject Logger to action struct
- [param](https://gitea.com/lunny/tango/wiki/Params) - get the router parameters
- [return](https://gitea.com/lunny/tango/wiki/Return) - Handle the returned value smartlly
- [context](https://gitea.com/lunny/tango/wiki/Context) - Inject context to action struct
- [session](https://gitea.com/tango/session) - [![CircleCI](https://circleci.com/gh/tango-contrib/session/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/session/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/session/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/session) Session manager, [session-redis](http://gitea.com/tango/session-redis), [session-nodb](http://gitea.com/tango/session-nodb), [session-ledis](http://gitea.com/tango/session-ledis), [session-ssdb](http://gitea.com/tango/session-ssdb)
- [xsrf](https://gitea.com/tango/xsrf) - [![CircleCI](https://circleci.com/gh/tango-contrib/xsrf/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/xsrf/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/xsrf/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/xsrf) Generates and validates csrf tokens
- [binding](https://gitea.com/tango/binding) - [![CircleCI](https://circleci.com/gh/tango-contrib/binding/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/binding/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/binding/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/binding) Bind and validates forms
- [renders](https://gitea.com/tango/renders) - [![CircleCI](https://circleci.com/gh/tango-contrib/renders/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/renders/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/renders/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/renders) Go template engine
- [dispatch](https://gitea.com/tango/dispatch) - [![CircleCI](https://circleci.com/gh/tango-contrib/dispatch/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/dispatch/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/dispatch/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/dispatch) Multiple Application support on one server
- [tpongo2](https://gitea.com/tango/tpongo2) - [![CircleCI](https://circleci.com/gh/tango-contrib/tpongo2/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/tpongo2/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/tpongo2/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/tpongo2) [Pongo2](https://github.com/flosch/pongo2) teamplte engine support
- [captcha](https://gitea.com/tango/captcha) - [![CircleCI](https://circleci.com/gh/tango-contrib/captcha/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/captcha/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/captcha/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/captcha) Captcha
- [events](https://gitea.com/tango/events) - [![CircleCI](https://circleci.com/gh/tango-contrib/events/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/events/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/events/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/events) Before and After
- [flash](https://gitea.com/tango/flash) - [![CircleCI](https://circleci.com/gh/tango-contrib/flash/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/flash/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/flash/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/flash) Share data between requests
- [debug](https://gitea.com/tango/debug) - [![CircleCI](https://circleci.com/gh/tango-contrib/debug/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/debug/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/debug/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/debug) show detail debug infomaton on log
- [basicauth](https://gitea.com/tango/basicauth) - [![CircleCI](https://circleci.com/gh/tango-contrib/basicauth/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/basicauth/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/basicauth/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/basicauth) basicauth middleware
- [authz](https://gitea.com/tango/authz) - [![Build Status](https://travis-ci.org/tango-contrib/authz.svg?branch=master)](https://travis-ci.org/tango-contrib/authz) [![Coverage Status](https://coveralls.io/repos/github/tango-contrib/authz/badge.svg?branch=master)](https://coveralls.io/github/tango-contrib/authz?branch=master) manage permissions via ACL, RBAC, ABAC
- [cache](https://gitea.com/tango/cache) - [![CircleCI](https://circleci.com/gh/tango-contrib/cache/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/cache/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/cache) cache middleware - cache-memory, cache-file, [cache-ledis](https://gitea.com/tango/cache-ledis), [cache-nodb](https://gitea.com/tango/cache-nodb), [cache-mysql](https://gitea.com/tango/cache-mysql), [cache-postgres](https://gitea.com/tango/cache-postgres), [cache-memcache](https://gitea.com/tango/cache-memcache), [cache-redis](https://gitea.com/tango/cache-redis)
- [rbac](https://gitea.com/tango/rbac) - [![CircleCI](https://circleci.com/gh/tango-contrib/rbac/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/rbac/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/rbac/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/rbac) rbac control

## Documentation

- [Manual](http://gobook.io/read/github.com/go-tango/manual-en-US/), And you are welcome to contribue for the book by git PR to [gitea.com/tango/manual-en-US](https://gitea.com/tango/manual-en-US)
- [操作手册](http://gobook.io/read/github.com/go-tango/manual-zh-CN/)，您也可以访问 [github.com/go-tango/manual-zh-CN](https://gitea.com/tango/manual-zh-CN)为本手册进行贡献
- [API Reference](https://godoc.org/gitea.com/lunny/tango)

## Discuss

- [Google Group - English](https://groups.google.com/forum/#!forum/go-tango)
- QQ Group - 简体中文 #369240307

## Cases

- [GopherTC](https://github.com/jimmykuu/gopher/tree/2.0) - China Discuss Forum
- [Wego](https://gitea.com/tango/wego) - Discuss Forum
- [dbweb](https://gitea.com/xorm/dbweb) - DB management web UI
- [Godaily](http://godaily.org) - [gitea](https://gitea.com/godaily/news)
- [Pugo](https://github.com/go-xiaohei/pugo) - A pugo blog
- [Gos](https://gitea.com/tango/gos) - Static web server
- [GoFtpd](https://github.com/goftp/ftpd) - Pure Go cross-platform ftp server

## License

This project is under BSD License. See the [LICENSE](LICENSE) file for the full license text.
