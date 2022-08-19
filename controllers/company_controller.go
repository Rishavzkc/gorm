package controllers

import (
	"mysqlgin/models"
	"mysqlgin/repositories"
	"mysqlgin/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	repo    repositories.CompanyRepository = repositories.CompanyRepository{}
	service services.CompanyService        = *services.NewCompanyService(&repo)
)

func ShowComp(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	result, err := service.ShowCompany(int(newid))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func ShowAllComp(c *gin.Context) {

	result, err := service.ShowAllCompanies()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func CreateComp(c *gin.Context) {
	var com models.Company
	var emp models.Employee

	err := c.ShouldBindJSON(&com)
	c.ShouldBindJSON(&emp)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	result, _, err := service.CreateCompany(com, emp)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}
	c.JSON(200, result)
}

func UpdateComp(c *gin.Context) {
	var com models.Company

	err := c.ShouldBindJSON(&com)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateCompany(com)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func DeleteComp(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})

		return
	}

	result, err := service.DeleteCompany(int(newid))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, result)
}

// repo := &repositories.CompanyRepository{}
// service := services.NewCompanyService(repo)
