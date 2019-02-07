package api

import (
	"github.com/PeppyS/api.peppysisay.com/api/routes"

	"github.com/gin-gonic/gin"
)

type API struct {
	Router  *gin.Engine
	RootAPI *routes.RootAPI
	Opts
}

type Opts struct {
	Version string
}

func New(router *gin.Engine, rootAPI *routes.RootAPI, opts Opts) *API {
	router.Use(enableCORS())

	rg := router.Group("/")

	rootAPI.SetupHandlers(rg)

	return &API{router, rootAPI, opts}
}

func enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Session-ID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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
