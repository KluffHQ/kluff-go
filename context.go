package kluff

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const _sdkKey = "kluff_sdk"

type Context struct {
	*gin.Context
	SDK *Interactor
}

type ApiError struct {
	status int
	reason interface{}
}

func (a *ApiError) Error() string {
	return fmt.Sprintf("STATUS: %d REASON: %v", a.status, a.reason)
}

func (c *Context) Set(key string, value any) {
	if key != _sdkKey {
		return
	}
	c.Set(key, value)
}

func (c *Context) Error(status int, reason interface{}) error {
	return &ApiError{
		status: status,
		reason: reason,
	}
}
