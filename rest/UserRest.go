package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piyabch/pi-api/model"
	"github.com/piyabch/pi-api/service"
)

func InitUserRest(e *gin.Engine) {
	e.POST("/users", createUser)
	e.GET("/users/:id", findUserById)
	e.GET("/users/search", findUserByName)
	e.PUT("/users", updateUser)
}

func createUser(c *gin.Context) {
	var inputData model.User
	err := c.BindJSON(&inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	createdUser, err := service.CreateUser(&inputData)
	if createdUser != nil && createdUser.ID <= 0 && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Create user failed"})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func findUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}
	userId, err := strconv.Atoi(id)
	if userId <= 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is invalid"})
		return
	}
	user, err := service.FindUserByID(userId)
	if user == nil || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found", "ID": id})
		return
	}
	c.JSON(http.StatusOK, user)
}

func findUserByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Name is required"})
		return
	}
	users, err := service.FindUsersByName(name)
	if len(users) == 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Users not found", "Name": name})
		return
	}
	c.JSON(http.StatusOK, users)
}

func updateUser(c *gin.Context) {
	var inputData model.User
	err := c.BindJSON(&inputData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return
	}
	if inputData.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}
	updatedUser, err := service.UpdateUser(&inputData)
	if updatedUser != nil && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update user failed"})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
