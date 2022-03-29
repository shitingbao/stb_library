package context

import "github.com/gin-gonic/gin"

type GContext struct {
	*gin.Context
	Username string
	Token    string
}

func New(g *gin.Context, options ...string) *GContext {
	name, token := "", ""
	if len(options) > 0 {
		name = options[0]
	}
	if len(options) > 1 {
		token = options[1]
	}
	return &GContext{g, name, token}
}
