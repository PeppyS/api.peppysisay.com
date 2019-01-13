package blog

import (
	"github.com/gin-gonic/gin"

	"github.com/PeppyS/api.peppysisay.com/api/blog/posts"
)

type BlogAPI struct {
	PostsAPI *posts.PostsAPI
}

func NewAPI(p *posts.PostsAPI) *BlogAPI {
	return &BlogAPI{p}
}

func (b *BlogAPI) SetupHandlers(r *gin.RouterGroup) {
	bg := r.Group("/blog")

	b.PostsAPI.SetupHandlers(bg)
}
