package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/piyabch/pi-api/config"
)

var engine *gin.Engine

// Start the rest server.
// The server startup is made by following steps.
//   - Create the default gin engine.
//   - Configure the rest endpoints.
//   - Start web server on the pre-configuration port.
func Start() {
	engine = gin.Default()
	// configure endpoints
	initDefaultPage(engine)
	InitUserRest(engine)
	// start sever on
	engine.Run(config.App.WebAddress)
}

// Default handler on the root path
func initDefaultPage(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		// reply status ok and show the server status
		c.String(200, "pi-api is running")
	})
}
