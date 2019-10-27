renders [![Build Status](https://drone.gitea.com/api/badges/tango/renders/status.svg)](https://drone.gitea.com/tango/renders)
 [![](http://gocover.io/_badge/gitea.com/tango/renders)](http://gocover.io/gitea.com/tango/renders)
======

Middleware renders is a go template render middlewaer for [Tango](https://gitea.com/lunny/tango). 

## Version

   v0.2.0510 Added RenderBytes for Renderer and simplifed codes.

## Installation

    go get gitea.com/tango/renders

## Simple Example

```Go
type RenderAction struct {
    renders.Renderer
}

func (x *RenderAction) Get() {
    x.Render("test.html", renders.T{
        "test": "test",
    })
}

func main() {
    t := tango.Classic()
    t.Use(renders.New(renders.Options{
        Reload: true, // if reload when template is changed
        Directory: "./templates", // Directory to load templates
        Funcs: template.FuncMap{
            "test": func() string {
                    return "test"
            },
        },
        // Vars is a data map for global
        Vars: renders.T{
            "var": var,
        }
        Charset: "UTF-8", // Appends the given charset to the Content-Type header. Default is UTF-8
        // Allows changing of output to XHTML instead of HTML. Default is "text/html"
        HTMLContentType: "text/html",
        DelimsLeft:"{{",
        DelimsRight:"}}", // default Delims is {{}}, if it conflicts with your javascript template such as angluar, you can change it.
    }))
}
```

## License

This project is under BSD License. See the [LICENSE](LICENSE) file for the full license text.