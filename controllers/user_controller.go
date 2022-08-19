package controllers

import (
	"mysqlgin/models"
	"mysqlgin/repositories"
	"mysqlgin/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	reposit repositories.UserRepository = repositories.UserRepository{}
	serve   services.UserService        = *services.NewUserService(&reposit)
)

func ShowUse(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	result, err := serve.ShowUser(int(newid))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func ShowAllUse(c *gin.Context) {

	result, err := serve.ShowAllUsers()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func CreateUse(c *gin.Context) {
	var use models.User

	err := c.ShouldBindJSON(&use)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	result, err := serve.CreateUser(use)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}
	c.JSON(200, result)
}
func UpdateUse(c *gin.Context) {
	var use models.User

	err := c.ShouldBindJSON(&use)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = serve.UpdateUser(use)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func DeleteUse(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})

		return
	}

	result, err := serve.DeleteUser(int(newid))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, result)
}
