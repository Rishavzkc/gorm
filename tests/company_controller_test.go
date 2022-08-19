package tests

import (
	"bytes"
	"encoding/json"
	"mysqlgin/config"
	"mysqlgin/controllers"
	"mysqlgin/database"
	"mysqlgin/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestShowAllComp(t *testing.T) {
	config.Init()
	database.StartDatabase()

	r := gin.Default()

	r.GET("/company/", controllers.ShowAllComp)
	req, err := http.NewRequest("GET", "/company/", nil)

	if err != nil {
		t.Fatalf("Not able to create request %s/n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []models.Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateComp(t *testing.T) {

	config.Init()
	database.StartDatabase()
	r := gin.Default()

	r.POST("/company/", controllers.CreateComp)

	company := models.Company{
		Name:     "Hp",
		CompanyID:       6,
		Location: "Blr",
	}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/company/", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateComp(t *testing.T) {
	config.Init()
	database.StartDatabase()
	r := gin.Default()
	r.PUT("/company/", controllers.UpdateComp)
	comp := models.Company{
		Name:     "Info",
		CompanyID:       1,
		Location: "Pune",
	}
	jsonValue, _ := json.Marshal(comp)
	reqFound, err := http.NewRequest("PUT", "/company/", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, 204, w.Code)

}

func TestShowComp(t *testing.T) {
	config.Init()
	database.StartDatabase()
	r := gin.Default()
	r.GET("/company/:id", controllers.ShowComp)

	req, err := http.NewRequest("GET", "/company/1", nil)

	if err != nil {
		t.Fatalf("Not able to get request %s/n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []models.Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeleteComp(t *testing.T) {
	config.Init()
	database.StartDatabase()

	r := gin.Default()

	r.DELETE("/company/:id", controllers.DeleteComp)

	req, err := http.NewRequest("DELETE", "/company/5", nil)

	if err != nil {
		t.Fatalf("Not able to delete request %s/n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []models.Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
}
