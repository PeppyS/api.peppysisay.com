package comments

import (
	"github.com/gin-gonic/gin"
)

type CommentsAPI struct {
	commentService *CommentService
}

func NewAPI(cs *CommentService) *CommentsAPI {
	return &CommentsAPI{cs}
}

func (a *CommentsAPI) SetupHandlers(rg *gin.RouterGroup) {
	c := rg.Group("/comments")

	c.POST("/", a.New())
}

func (a *CommentsAPI) New() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
