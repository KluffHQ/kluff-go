package kluff

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type actionMeta struct {
	Route  string `json:"route"`
	Name   string `json:"name"`
	Object string `json:"object"`
	ID     string `json:"id"`
}

type meta struct {
	Actions  map[string]actionMeta  `json:"actions"`
	Triggers map[string]triggerMeta `json:"triggers"`
}

type Router struct {
	engine  *gin.Engine
	api     *gin.RouterGroup
	action  *gin.RouterGroup
	trigger *gin.RouterGroup

	// this is used to keep track of all metadata
	meta meta
}

type APIHandler func(*Context)

const (
	apiPath     = "/api"
	actionPath  = "/action"
	triggerPath = "/trigger"
	metaPath    = "/_meta"
)

func NewRouter() *Router {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return &Router{
		engine:  engine,
		api:     engine.Group(apiPath, requestValidator()),
		action:  engine.Group(actionPath, requestValidator()),
		trigger: engine.Group(triggerPath, requestValidator()),
		meta: meta{
			Actions:  map[string]actionMeta{},
			Triggers: map[string]triggerMeta{},
		},
	}
}

func (r *Router) Start() error {
	r.engine.GET(metaPath, func(ctx *gin.Context) {
		ctx.JSON(200, r.meta)
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

func (a Action) getRoute() string {
	return fmt.Sprintf("%s/%s", actionPath, parseID(a.ID))
}

func (r *Router) handleAction(a Action) error {
	meta := actionMeta{
		Route:  a.getRoute(),
		Object: a.Object,
		Name:   a.Name,
		ID:     parseID(a.ID),
	}
	r.meta.Actions[parseID(a.ID)] = meta
	r.action.POST(fmt.Sprintf("/%s", parseID(a.ID)), func(ctx *gin.Context) {
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
	_, ok := r.meta.Actions[parseID(a.ID)]
	if ok {
		log.Fatalf("duplicate action id [%s]", a.ID)
	}
	return r.handleAction(a)
}

/* ================================================
*  Triggers
* ================================================== */

type TriggerAction string

const (
	ON_CREATE TriggerAction = "ON_SAVE"
	ON_UPDATE TriggerAction = "ON_UPDATE"
	ON_DELETE TriggerAction = "ON_DELETE"
)

type TriggerHandler func(ctx *Context, data map[string]any) error

type triggerMeta struct {
	Object string
	Type   TriggerAction
	Route  string
}

type Trigger struct {
	ID      string
	Action  TriggerAction
	Object  string
	Handler TriggerHandler
}

func (t Trigger) getRoute() string {
	return fmt.Sprintf("%s/%s", triggerPath, parseID(t.ID))
}

func (t Trigger) validate() error {
	if t.ID == "" {
		return errors.New("trigger ID required")
	}
	if t.Action == ON_CREATE || t.Action == ON_DELETE || t.Action == ON_UPDATE {
		return nil
	}
	return errors.New("invalid trigger action")
}

func (r *Router) RegisterTrigger(trigger Trigger) {
	if err := trigger.validate(); err != nil {
		log.Fatal(err)
	}
	id := parseID(trigger.ID)
	_, ok := r.meta.Triggers[id]
	if ok {
		log.Fatal("failed to register multiple trigger with the same id")
	}
	meta := triggerMeta{
		Object: trigger.Object,
		Type:   trigger.Action,
		Route:  trigger.getRoute(),
	}
	r.meta.Triggers[id] = meta
	r.trigger.POST(id, func(ctx *gin.Context) {
		sdk, err := parseSdk(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var data map[string]any
		if err := ctx.BindJSON(&data); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		context := &Context{
			Context: ctx,
			Inter:   sdk,
		}
		trigger.Handler(context, data)
	})
}
