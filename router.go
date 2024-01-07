package kluff

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ActionMeta struct {
	Route  string `json:"route"`
	Object string `json:"object"`
	ID     string `json:"id"`
}

type BathMeta struct {
}

type Router struct {
	engine *gin.Engine
	api    *gin.RouterGroup
	action *gin.RouterGroup

	// this is used to keep track of all metadata
	actions map[string]ActionMeta
}

type APIHandler func(*Context)

const (
	apiPath    = "/api"
	actionPath = "/action"
	metaPath   = "/_meta"
)

func NewRouter() *Router {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return &Router{
		engine:  engine,
		api:     engine.Group(apiPath, requestValidator()),
		action:  engine.Group(actionPath, requestValidator()),
		actions: map[string]ActionMeta{},
	}
}

func (r *Router) Start() error {
	r.engine.GET(metaPath, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"actions": r.actions,
		})
	})
	return r.engine.Run(":3000")
}

/*
API HANDLERS
*/
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

/*
 Actions
*/

type ActionHandler func(*Interactor, *Record) error

type Action struct {
	// the id must be unique for every action registered
	ID      string
	Name    string
	Object  string
	Handler ActionHandler
}

func (a Action) validate() error {
	if a.ID == "" {
		return errors.New("`ID` is required")
	}
	if a.Name == "" {
		return errors.New("`Name` is required")
	}
	if a.Handler == nil {
		return errors.New("invalid action handler")
	}
	if a.Object == "" {
		return errors.New("object is required")
	}
	return nil
}

func (a Action) parseID() string {
	return strings.ReplaceAll(strings.ToLower(a.ID), " ", "")
}

func (a Action) getRoute() string {
	return fmt.Sprintf("%s/%s", actionPath, a.parseID())
}

func (r *Router) handleAction(a Action) error {
	meta := ActionMeta{
		Route:  a.getRoute(),
		Object: a.Object,
		ID:     a.parseID(),
	}
	r.actions[a.parseID()] = meta
	r.action.POST(fmt.Sprintf("/%s", a.parseID()), func(ctx *gin.Context) {
		sdk, err := parseSdk(ctx)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var rec map[string]any
		err = ctx.BindJSON(&rec)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		r := newRecord(sdk.cl, a.Object, rec)
		err = a.Handler(sdk, r)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
	})
	return nil
}

func (r *Router) RegisterAction(a Action) error {
	if err := a.validate(); err != nil {
		log.Fatalf("%s while creating action [%s]", err.Error(), a.Name)
	}
	_, ok := r.actions[a.parseID()]
	if ok {
		log.Fatalf("duplicate action id [%s]", a.ID)
	}
	return r.handleAction(a)
}
