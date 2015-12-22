jaderender
=========

Package jaderender is a template renderer that can be used with the Gin web
framework https://github.com/gin-gonic/gin it uses the jaderender template library
https://github.com/zdebeer99/gojade

Usage
-----

To use jaderender you need to set your router.HTMLRenderer to a new renderer
instance, this is done after creating the Gin router when the Gin application
starts up. You can use jaderender.Default() to create a new renderer with
default options, this assumes templates will be located in the "views"
directory, or you can use jaderender.New() to specify a custom location.

To render templates from a route, call c.HTML just as you would with
regular Gin templates.

Basic Example
-------------

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/flosch/pongo2"
    "github.com/robvdl/pongo2gin"
)

func main() {
    router := gin.Default()

    // Use jaderender.Default() for default options or jaderender.New()
    // if you need to use custom RenderOptions.
    router.HTMLRender = jaderender.Default()

    router.GET("/", func(c *gin.Context) {
        c.HTML(200, "hello.html", gin.H{"name": "world"})
    })

    router.Run(":8080")
}
```

RenderOptions
-------------

When calling jaderender.New() instead of jaderender.Default() you can use these
custom RenderOptions:

```go
type RenderOptions struct {
    TemplateDir string  // location of the template directory
    Beautify    bool // beautify the resulting HTML
}
```
