package kluff

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseSdk(c *gin.Context) (*Interactor, error) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		bearer, err := c.Cookie("x-kluff-auth")
		if err != nil {
			return nil, err
		}
		if bearer == "" {
			return nil, errors.New("cookie not found")
		}
	}
	return Get(bearer)
}

func buildContext(c *gin.Context) (*Context, error) {
	res, ok := c.Get(_sdkKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return nil, errors.New("sdk not fond on the context")
	}

	sdk, ok := res.(*Interactor)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return nil, errors.New("invalid sdk value")
	}
	return &Context{
		Context: c,
		SDK:     sdk,
	}, nil
}

func requestValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sdk, err := parseSdk(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid request token"})
			return
		}
		ctx.Set(_sdkKey, sdk)
		ctx.Next()
	}
}
