package routes

import (
	"net/http"

	"github.com/PeppyS/api.peppysisay.com/api/routes/blog"
	"github.com/gin-gonic/gin"
)

type RootAPI struct {
	BlogAPI *blog.BlogAPI
}

func NewAPI(blogAPI *blog.BlogAPI) *RootAPI {
	return &RootAPI{blogAPI}
}

func (a *RootAPI) SetupHandlers(rg *gin.RouterGroup) {
	rg.GET("/", a.Get())

	a.BlogAPI.SetupHandlers(rg.Group("/blog"))
}

func (a *RootAPI) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"description": "Welcome to Peppy's API 👋🏾",
			"experience": []gin.H{
				gin.H{
					"title":       "Platform Engineer",
					"company":     "Dave.com",
					"location":    "Los Angeles, CA",
					"date_range":  "August 2019 - Present",
					"description": `Helping milions of people thrive financially - https://dave.com/careers`,
				},
				gin.H{
					"title":       "Senior Software Engineer",
					"company":     "Mothership",
					"location":    "Los Angeles, CA",
					"date_range":  "July 2018 - July 2019",
					"description": `Helping bring business shipping to the modern age! Building out freight aggregation technology to connect businesses with nearby trucks and vans with extra space to allow for the fastest shipping at the lowest rates.`,
				},
				gin.H{
					"title":       "Software Engineer",
					"company":     "Tradesy",
					"location":    "Los Angeles, CA",
					"date_range":  "January 2016 - June 2018",
					"description": `Full stack engineer on the Shopping Experience scrum team. Worked with Product & Design team to build out APIs & front-end implementations to support various product features & engineering initiatives`,
				},
				gin.H{
					"title":       "Software Engineer",
					"company":     "Universy of Southern California",
					"location":    "Los Angeles, CA",
					"date_range":  "December 2014 - January 2016",
					"description": `Feature development on inventory management web app for USC art students`,
				},
				gin.H{
					"title":       "Intern Software Engineer",
					"company":     "AT&T",
					"location":    "Seattle, CA",
					"date_range":  "June 2015 - August 2015",
					"description": `Joined a team developing an iPad app for the company’s sales team`,
				},
			},
			"education": []gin.H{
				gin.H{
					"school":     "University of Southern California",
					"degree":     "Bachelor of Science, Computer Science",
					"date_range": "2012 - 2016",
				},
			},
		})
	}
}
