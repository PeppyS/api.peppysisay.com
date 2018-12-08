package api

import (
	"net/http"

	"github.com/PeppyS/api.peppysisay.com/api/blog"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router  *gin.Engine
	BlogAPI *blog.BlogAPI
	Opts
}

type Opts struct {
	Version string
}

func New(router *gin.Engine, blogAPI *blog.BlogAPI, opts Opts) *API {
	router.Use(enableCORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"version": opts.Version,
		})
	})

	blogAPI.SetupHandlers(&router.RouterGroup)

	return &API{router, blogAPI, opts}
}

func enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (a *API) Run(port string) {
	a.Router.Run(port)
}
