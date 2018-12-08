package blog

import (
	"github.com/gin-gonic/gin"

	"github.com/PeppyS/api.peppysisay.com/api/blog/comments"
	"github.com/PeppyS/api.peppysisay.com/api/blog/posts"
)

type BlogAPI struct {
	CommentsAPI *comments.CommentsAPI
	PostsAPI    *posts.PostsAPI
}

func NewAPI(c *comments.CommentsAPI, p *posts.PostsAPI) *BlogAPI {
	return &BlogAPI{c, p}
}

func (b *BlogAPI) SetupHandlers(r *gin.RouterGroup) {
	bg := r.Group("/blog")

	b.CommentsAPI.SetupHandlers(bg)
	b.PostsAPI.SetupHandlers(bg)
}
