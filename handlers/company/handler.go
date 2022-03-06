package company

import (
	"net/http"
	"tinyapps/api-base/handlers"
	"tinyapps/api-base/models"

	"github.com/gin-gonic/gin"
)

func Register(env *handlers.Env, r *gin.Engine) {
	v1 := r.Group("/companies")
	{
		v1.GET("", func(c *gin.Context) {
			getCompanies(env, c)
		})
		v1.POST("", func(c *gin.Context) {
			createCompany(env, c)
		})
		v1.GET("/dummy", func(c *gin.Context) {
			getDummyCompany(env, c)
		})
	}
}

func getCompanies(env *handlers.Env, c *gin.Context) {
	companies := models.GetCompanies(env, 5)
	c.JSON(http.StatusOK, companies)
}

func getDummyCompany(env *handlers.Env, c *gin.Context) {
	companies := []models.Company{models.DummyCompany()}
	c.JSON(http.StatusOK, companies)
}

func createCompany(env *handlers.Env, c *gin.Context) {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.SaveCompany(env, company)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
