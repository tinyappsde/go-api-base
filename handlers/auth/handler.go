package auth

import (
	"net/http"
	authHelper "tinyapps/api-base/auth"
	"tinyapps/api-base/handlers"

	"github.com/gin-gonic/gin"
)

func Register(env *handlers.Env, r *gin.Engine) {
	r.POST("/auth", func(c *gin.Context) {
		auth(env, c)
	})
	r.GET("/auth/user", func(c *gin.Context) {
		getAuthenticatedUser(env, c)
	})
}

func auth(env *handlers.Env, c *gin.Context) {
	type Credentials struct {
		Email		string `json:"email"`
		Password	string `json:"password"`
	}

	var credentials Credentials

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := authHelper.CheckCredentials(env, credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken := authHelper.IssueToken(user)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"accessToken": accessToken,
	})
}

func getAuthenticatedUser(env *handlers.Env, c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	user, err := authHelper.ValidateAccessToken(env, authHeader)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user": user,
	})
}
