package controllers

import (
	"mysqlgin/models"
	"mysqlgin/repositories"
	"mysqlgin/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	repoes repositories.EmployeeRepository = repositories.EmployeeRepository{}
	servic services.EmployeeService        = *services.NewEmployeeService(&repoes)
)

func ShowEmpl(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	result, err := servic.ShowEmployee(int(newid))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func ShowAllEmpl(c *gin.Context) {

	result, err := servic.ShowAllEmployee()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func CreateEmpl(c *gin.Context) {
	var empl models.Employee

	err := c.ShouldBindJSON(&empl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	result, err := servic.CreateEmployee(empl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}
	c.JSON(200, result)
}

func UpdateEmpl(c *gin.Context) {
	var empl models.Employee

	err := c.ShouldBindJSON(&empl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = servic.UpdateEmployee(empl)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func DeleteEmpl(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})

		return
	}

	result, err := servic.DeleteEmployee(int(newid))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, result)
}
