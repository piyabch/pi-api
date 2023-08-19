package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/piyabch/pi-api/config"
)

var engine *gin.Engine

func Init() {
	engine = gin.Default()
	initDefaultPage(engine)
	InitUserRest(engine)
}

func Start() {
	engine.Run(config.App.WebAddress)
}

func initDefaultPage(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.String(200, "pi-api is running")
	})
}
