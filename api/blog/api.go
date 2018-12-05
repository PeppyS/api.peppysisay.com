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

func NewAPI(r *gin.Engine) *BlogAPI {
	b := r.Group("/blog")

	postsAPI := posts.NewAPI(b)
	commentsAPI := comments.NewAPI(b)

	return &BlogAPI{commentsAPI, postsAPI}
}

func (b *BlogAPI) SetupHandlers() {
	b.CommentsAPI.SetupHandlers()
	b.PostsAPI.SetupHandlers()
}
