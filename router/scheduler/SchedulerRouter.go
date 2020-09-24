package scheduler

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func JobJson(c *gin.Context) {
	// returns a map[string]interface{} that can be marshalled as JSON
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "Status.html", jobrunner.StatusPage())

}

func SetSchedulerRouter(router *gin.Engine) {

	router.GET("api/v1API/jobrunner/json", JobJson)

	router.LoadHTMLGlob("html/template/Status.html")
	router.GET("/jobrunner/html", JobHtml)
}
