package posts

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *PostsAPI) SetupHandlers() {
	c := a.RouterGroup.Group("/posts")

	c.GET("/:id", a.Get())
}

func (a *PostsAPI) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"likes_count": 0,
				"comments": []gin.H{
					gin.H{
						"post_id":     postID,
						"created_at":  time.Now().String(),
						"first_name":  "Bob",
						"last_name":   "Smith",
						"text":        "Great write-up! Keep up the good work.",
						"likes_count": "0",
					},
				},
			},
		})
	}
}
