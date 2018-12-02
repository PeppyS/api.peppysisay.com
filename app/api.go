package app

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)
type API struct {
	router *gin.Engine
}

func NewAPI() *API {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	p := r.Group("/post")
	
	p.GET("/:id/comments", func(c * gin.Context) {
		postID := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": []gin.H{
				gin.H{
					"post_id": postID,
					"created_at": time.Now().String(),
					"first_name": "Bob",
					"last_name": "Smith",
					"text": "Great write-up! Keep up the good work.",
				},
			},
		})
	})

	return &API{r}
}

func (a *API) Run(port string) {
	a.router.Run(port)
}
