// Package jaderender is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the gojade template
// library https://github.com/zdebeer99/gojade

package jaderender

import (
	"net/http"
	"github.com/Lnd-stoL/gojade"
	"github.com/gin-gonic/gin/render"
)


// RenderOptions is used to configure the renderer.
type RenderOptions struct {
	TemplateDir string
	Beautify    bool
	CacheSize   int
}


// JadeRender is a custom Gin template renderer using gojade.
type JadeRender struct {
	Template     *gojade.Engine
	Context      interface{}
	TemplateName string

	cache        *renderCache
}


// New creates a new JadeRender instance with custom Options.
func New(options RenderOptions) *JadeRender {
	this := &JadeRender{
		Template: gojade.New(),
	}
	this.Template.ViewPath = options.TemplateDir
	this.Template.Beautify = options.Beautify

	if options.CacheSize > 0 {
		this.cache = newRenderCache(options.CacheSize)
	}

	return this
}

// Default creates a JadeRender instance with default options.
func Default() *JadeRender {
	return New(RenderOptions{
		TemplateDir: "views",
		Beautify: false,
		CacheSize: 128,
	})
}

// Instance should return a new JadeRender struct per request
func (this *JadeRender) Instance(templateName string, data interface{}) render.Render {
	return JadeRender{
		Template: this.Template,
		Context:  data,
		TemplateName: templateName,
		cache: this.cache,
	}
}

// Render should render the template to the response.
func (this JadeRender) Render(w http.ResponseWriter) error {
	writeContentType(w, []string{"text/html; charset=utf-8"})

	// the cache is disabled in JadeRender options
	if this.cache == nil {
		this.Template.RenderFileW(w, this.TemplateName, this.Context)
		return nil
	}

	rendered, existsInCache := this.cache.Get(this.TemplateName, this.Context)
	if !existsInCache {
		rendered = this.Template.RenderFile(this.TemplateName, this.Context).Bytes()
		this.cache.Add(this.TemplateName, this.Context, rendered)
	}

	w.Write(rendered)
	return nil
}

// writeContentType is also in the gin/render package but it has not been made
// pubic so is repeated here, maybe convince the author to make this public.
func (this JadeRender) WriteContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
