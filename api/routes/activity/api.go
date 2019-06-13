package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/PeppyS/api.peppysisay.com/internal/pkg/service"
)

type ActivityAPI struct {
	activityService *service.ActivityService
}

func NewAPI(activityService *service.ActivityService) *ActivityAPI {
	return &ActivityAPI{activityService}
}

func (a *ActivityAPI) SetupHandlers(rg *gin.RouterGroup) {
	rg.GET("/code", a.GetCode())
}

func (a *ActivityAPI) GetCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": a.activityService.GetRecentCodeCommits()
		})
	}
}
