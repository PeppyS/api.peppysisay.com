package app

import (
	"net/http"
	"time"

	"github.com/PeppyS/api.peppysisay.com/app/middleware"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
}

func NewAPI() *API {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	p := r.Group("/post")

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

	return &API{r}
}

func (a *API) Run(port string) {
	a.router.Run(port)
}
