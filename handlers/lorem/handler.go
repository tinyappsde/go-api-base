package lorem

import (
	"net/http"
	"tinyapps/api-base/handlers"

	"github.com/gin-gonic/gin"
)

func Register(env *handlers.Env, r *gin.Engine) {
	r.GET("/lorem", getLorem)
	r.POST("/lorem", postLorem)
}

func getLorem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Lorem Ipsum dolor sit amet.",
	})
}

func postLorem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is an example POST endpoint.",
	})
}
