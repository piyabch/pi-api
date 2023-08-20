package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piyabch/pi-api/model"
	"github.com/piyabch/pi-api/service"
)

func InitAuthRest(e *gin.Engine) {
	e.POST("/auth", authHandler)
}

// Validation handler to filter API requests
func CheckAuth(c *gin.Context) bool {
	// get token from header
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication required"})
		return false
	}
	// check the validity of the token
	if !service.IsAuthTokenValid(authToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return false
	}
	return true
}

// Authorize to get the token for calling REST APIs.
// The required parameters are as follows.
//   - Email
//   - Password
//
// Return the authorized token
func authHandler(c *gin.Context) {
	// parse the auth data
	var authData model.AuthData
	err := c.BindJSON(&authData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid auth data"})
		return
	}
	// validate required fields
	if authData.Email == "" || authData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid auth data"})
		return
	}
	// call the service
	authResult, err := service.Authorize(&authData)
	if authResult == nil || authResult.Token == "" || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Authentication failed"})
		return
	}
	// build the response
	c.JSON(http.StatusOK, authResult)
}
