package comments

import (
	"github.com/gin-gonic/gin"
)

type CommentsAPI struct {
	RouterGroup *gin.RouterGroup
}

func NewAPI(r *gin.RouterGroup) *CommentsAPI {
	return &CommentsAPI{r}
}
