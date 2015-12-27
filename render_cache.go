// Package jaderender is a template renderer that can be used with the Gin
// web framework https://github.com/gin-gonic/gin it uses the gojade template
// library https://github.com/zdebeer99/gojade

package jaderender

import (
	"github.com/hashicorp/golang-lru"
    "fmt"
)


// Some package-internal structure helping to implemlement render cache
type renderCache struct {
	cache *lru.Cache
}

func newRenderCache(maxCacheEnties int) *renderCache {
	arcCache, _ := lru.New(maxCacheEnties)

	return &renderCache{
		cache: arcCache,
	}
}

func (c *renderCache) Get(templateName string, context interface{}) ([]byte, bool) {
	rendered, alreadyInCache := c.cache.Get(c.generateCacheKey(templateName, context))

	if alreadyInCache {
		return rendered.([]byte), true
	} else {
		return nil, false
	}
}

func (c *renderCache) Add(templateName string, context interface{}, rendered []byte) {
	c.cache.Add(c.generateCacheKey(templateName, context), rendered)
}

// Probabbly not the best cache key in terms of performance but
// that Go has no way of defining custom hashable types so no elegant simple solution is avaliable here
func (c *renderCache) generateCacheKey(templateName string, context interface{}) string {
	return fmt.Sprintf("%s %v", templateName, context)
}
