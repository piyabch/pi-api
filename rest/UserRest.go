package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piyabch/pi-api/model"
	"github.com/piyabch/pi-api/service"
)

// Route the user endpoints to the handlers.
func InitUserRest(e *gin.Engine) {
	e.POST("/users", createUser)
	e.GET("/users/:id", findUserById)
	e.GET("/users/search", findUserByName)
	e.PUT("/users", updateUser)
}

// Create a new user from the JSON string.
// The required fields are as follows.
//   - Firstname
//   - Lastname
//   - Email
//
// Return the JSON string of the newly created user
// with a given ID.
func createUser(c *gin.Context) {
	// check authentication
	if !CheckAuth(c) {
		return
	}
	// parse the input
	var inputData model.User
	err := c.BindJSON(&inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	// validate required fields
	if inputData.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Firstname is required"})
		return
	}
	if inputData.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "LastName is required"})
		return
	}
	if inputData.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is required"})
		return
	}
	// call the service
	createdUser, err := service.CreateUser(&inputData)
	if createdUser != nil && createdUser.ID <= 0 && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Create user failed"})
		return
	}
	// build the response
	c.JSON(http.StatusOK, createdUser)
}

// Find a user from the ID string.
// The required parameters are as follows.
//   - id
//
// Return the JSON string of the user
// matched with a input ID.
func findUserById(c *gin.Context) {
	// check authentication
	if !CheckAuth(c) {
		return
	}
	// parse the input
	id := c.Param("id")
	// validate required fields
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}
	userId, err := strconv.Atoi(id)
	if userId <= 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is invalid"})
		return
	}
	// call the service
	user, err := service.FindUserByID(userId)
	if user == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found", "ID": id})
		return
	}
	// build the response
	c.JSON(http.StatusOK, user)
}

// Find a user from the name string
// by searching wildcard on the Firstname field.
// The required parameters are as follows.
//   - name
//
// Return the JSON string of the users
// founded with a input name.
func findUserByName(c *gin.Context) {
	// check authentication
	if !CheckAuth(c) {
		return
	}
	// parse the input
	name := c.Query("name")
	// validate required fields
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name is required"})
		return
	}
	// call the service
	users, err := service.FindUsersByName(name)
	if len(users) == 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Users not found", "Name": name})
		return
	}
	// build the response
	c.JSON(http.StatusOK, users)
}

// Update user information from the JSON string.
// The required fields are as follows.
//   - ID
//   - Firstname
//   - Lastname
//   - Email
//
// Return the JSON string of updated user
// with a given ID.
func updateUser(c *gin.Context) {
	// check authentication
	if !CheckAuth(c) {
		return
	}
	// parse the input
	var inputData model.User
	err := c.BindJSON(&inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	// validate required fields
	if inputData.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}
	if inputData.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Firstname is required"})
		return
	}
	if inputData.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "LastName is required"})
		return
	}
	if inputData.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email is required"})
		return
	}
	// call the service
	updatedUser, err := service.UpdateUser(&inputData)
	if updatedUser != nil && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update user failed"})
		return
	}
	// build the response
	c.JSON(http.StatusOK, updatedUser)
}
