flash [![Build Status](https://drone.gitea.com/api/badges/tango/flash/status.svg)](https://drone.gitea.com/tango/flash) [![](http://gocover.io/_badge/gitea.com/tango/flash)](http://gocover.io/gitea.com/tango/flash)
======

Middleware flash is a tool for share data between requests for [Tango](https://gitea.com/lunny/tango). 

## Notice

This is a new version, it stores all data via [session](https://gitea.com/tango/session) not cookie. And it is slightly non-compitable with old version.

## Installation

    go get gitea.com/tango/flash

## Simple Example

```Go

import "gitea.com/tango/session"

type FlashAction struct {
    flash.Flash
}

func (x *FlashAction) Get() {
    x.Flash.Set("test", "test")
}

func (x *FlashAction) Post() {
   x.Flash.Get("test").(string) == "test"
}

func main() {
    t := tango.Classic()
    sessions := session.Sessions()
    t.Use(flash.Flashes(sessions))
    t.Any("/", new(FlashAction))
    t.Run()
}
```

## License

This project is under BSD License. See the [LICENSE](LICENSE) file for the full license text.