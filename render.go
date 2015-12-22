// Package jaderender is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the gojade template
// library https://github.com/zdebeer99/gojade

package jaderender

import (
	"net/http"
	"github.com/zdebeer99/gojade"
	"github.com/gin-gonic/gin/render"
)

// RenderOptions is used to configure the renderer.
type RenderOptions struct {
	TemplateDir string
	Beautify    bool
}


// JadeRender is a custom Gin template renderer using gojade.
type JadeRender struct {
	Template     *gojade.Engine
	Context      interface{}
	TemplateName string
}

// New creates a new JadeRender instance with custom Options.
func New(options RenderOptions) *JadeRender {
	this := &JadeRender{
		Template: gojade.New(),
	}
	this.Template.ViewPath = options.TemplateDir
	this.Template.Beautify = options.Beautify
	return this
}

// Default creates a JadeRender instance with default options.
func Default() *JadeRender {
	return New(RenderOptions{
		TemplateDir: "views",
		Beautify: false,
	})
}

// Instance should return a new JadeRender struct per request
func (this JadeRender) Instance(templateName string, data interface{}) render.Render {
	return JadeRender{
		Template: this.Template,
		Context:  data,
		TemplateName: templateName,
	}
}

// Render should render the template to the response.
func (this JadeRender) Render(w http.ResponseWriter) error {
	writeContentType(w, []string{"text/html; charset=utf-8"})
	return this.Template.RenderFileW(w, this.TemplateName, this.Context)
}

// writeContentType is also in the gin/render package but it has not been made
// pubic so is repeated here, maybe convince the author to make this public.
func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
