package comments

import (
	"github.com/gin-gonic/gin"
)

func (a *CommentsAPI) SetupHandlers() {
	c := a.RouterGroup.Group("/comments")

	c.POST("/", a.New())
}

func (a *CommentsAPI) New() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
