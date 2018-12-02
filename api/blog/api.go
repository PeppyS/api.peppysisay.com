package blog

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupAPI(r *gin.Engine) {
	b := r.Group("/blog")

	p := b.Group("/post")

	p.GET("/:id", func(c *gin.Context) {
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
	})
}
