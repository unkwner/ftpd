session [![Build Status](https://drone.gitea.com/api/badges/tango/session/status.svg)](https://drone.gitea.com/tango/session) [![](http://gocover.io/_badge/gitea.com/tango/session)](http://gocover.io/gitea.com/tango/session)
======

Session is a session middleware for [Tango](https://gitea.com/lunny/tango).

## Backend Supports

Currently session support some backends below:

* Memory - memory as a session store, this is the default store
* [nodb](http://gitea.com/tango/session-nodb) - nodb as a session store
* [redis](http://gitea.com/tango/session-redis) - redis server as a session store
* [ledis](http://gitea.com/tango/session-ledis) - ledis server as a session store
* [ssdb](http://gitea.com/tango/session-ssdb) - ssdb server as a session store

## Installation

    go get gitea.com/tango/session

## Simple Example

```Go
package main

import (
    "gitea.com/lunny/tango"
    "gitea.com/tango/session"
)

type SessionAction struct {
    session.Session
}

func (a *SessionAction) Get() string {
    a.Session.Set("test", "1")
    return a.Session.Get("test").(string)
}

func main() {
    o := tango.Classic()
    o.Use(session.New(session.Options{
        MaxAge:time.Minute * 20,
        }))
    o.Get("/", new(SessionAction))
}
```

## Getting Help

- [API Reference](https://godoc.org/gitea.com/tango/session)

## License

This project is under BSD License. See the [LICENSE](LICENSE) file for the full license text.
