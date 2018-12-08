package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostsAPI struct {
	postService *PostService
}

func NewAPI(ps *PostService) *PostsAPI {
	return &PostsAPI{ps}
}

func (a *PostsAPI) SetupHandlers(rg *gin.RouterGroup) {
	c := rg.Group("/posts")

	c.GET("/", a.GetAll())
	c.GET("/:id", a.GetByID())
}

func (a *PostsAPI) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, err := a.postService.GetAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": posts,
		})
	}
}

func (a *PostsAPI) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postID := ctx.Param("id")

		post, err := a.postService.GetByID(ctx, postID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": post,
		})
	}
}
