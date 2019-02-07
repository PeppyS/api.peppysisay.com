package blog

import (
	"github.com/gin-gonic/gin"

	"github.com/PeppyS/api.peppysisay.com/api/routes/blog/posts"
)

type BlogAPI struct {
	PostsAPI *posts.PostsAPI
}

func NewAPI(p *posts.PostsAPI) *BlogAPI {
	return &BlogAPI{p}
}

func (b *BlogAPI) SetupHandlers(rg *gin.RouterGroup) {
	b.PostsAPI.SetupHandlers(rg.Group("/posts"))
}
