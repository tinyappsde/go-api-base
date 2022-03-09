package routing

import (
	"tinyapps/api-base/handlers"
	"tinyapps/api-base/handlers/auth"
	"tinyapps/api-base/handlers/company"
	"tinyapps/api-base/handlers/lorem"

	"github.com/gin-gonic/gin"
)

func Setup(env *handlers.Env, r *gin.Engine) {
	lorem.Register(env, r)
	company.Register(env, r)
	auth.Register(env, r)
}
