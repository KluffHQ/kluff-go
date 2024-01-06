package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	api    *gin.RouterGroup
}

type APIHandler func(*Context)

func NewRouter() *Router {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(requestValidator())
	api := engine.Group("/api")
	return &Router{
		engine: engine,
		api:    api,
	}
}

func (r *Router) Start() error {
	return r.engine.Run(":3000")
}
func (r *Router) handleApi(method, path string, handlers []APIHandler) error {
	_h := []gin.HandlerFunc{}
	for _, h := range handlers {
		_h = append(_h, func(c *gin.Context) {
			ctx, err := buildContext(c)
			if err != nil {
				return
			}
			h(ctx)
		})
	}
	r.api.Handle(method, path, _h...)
	return nil
}

func (r *Router) GET(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodGet, path, handlers)
}

func (r *Router) POST(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodPost, path, handlers)
}

func (r *Router) PUT(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodPut, path, handlers)
}

func (r *Router) PATCH(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodPatch, path, handlers)
}

func (r *Router) DELETE(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodDelete, path, handlers)
}

func (r *Router) CONNECT(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodConnect, path, handlers)
}

func (r *Router) HEAD(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodHead, path, handlers)
}

func (r *Router) OPTIONS(path string, handlers ...APIHandler) error {
	return r.handleApi(http.MethodOptions, path, handlers)
}
