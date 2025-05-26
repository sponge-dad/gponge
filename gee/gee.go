package gee

import (
	"log"
	"net/http"
	"strings"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // support middleware
	engine      *Engine       // all groups share Engine instance
}

type Engine struct {
	*RouterGroup
	router *router
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
}

func cleanPath(path string) string {
	if path == "" {
		return ""
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	// 替换多个连续的 / 为一个
	for strings.Contains(path, "//") {
		path = strings.ReplaceAll(path, "//", "/")
	}
	// 去除末尾的 /
	if len(path) > 1 && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	return path
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	cleanPrefix := cleanPath(prefix)
	newGroup := &RouterGroup{
		prefix: g.prefix + cleanPrefix,
		engine: engine,
	}
	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}
func (g *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	g.addRoute("DELETE", pattern, handler)
}

func (g *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	g.addRoute("PUT", pattern, handler)
}

func (g *RouterGroup) Run(address string) error {
	return http.ListenAndServe(address, g)
}

func (g *RouterGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	g.engine.router.handle(ctx)
}
