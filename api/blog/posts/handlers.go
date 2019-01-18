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
	c.POST("/:id/comments", a.NewComment())
	c.DELETE("/:id/comments/:commentID", a.DeleteComment())
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

func (a *PostsAPI) NewComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postID := ctx.Param("id")

		var request struct {
			Text string `json:"text"`
			Name string `json:"name"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}

		comment, err := a.postService.AddComment(ctx, postID, request.Text, request.Name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": comment,
		})
	}
}

func (a *PostsAPI) DeleteComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postID := ctx.Param("id")
		commentID := ctx.Param("commentID")

		err := a.postService.DeleteComment(ctx, postID, commentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
