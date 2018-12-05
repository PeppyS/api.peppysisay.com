package posts

import (
	"github.com/gin-gonic/gin"
)

type PostsAPI struct {
	RouterGroup *gin.RouterGroup
}

func NewAPI(r *gin.RouterGroup) *PostsAPI {
	return &PostsAPI{r}
}
