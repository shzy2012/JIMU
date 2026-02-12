package user

import (
	"tolo/src/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// Routes initialize routing information
func Routes(route *gin.Engine) {

	r := route.Group("/user")
	r.Use(jwt.JWTAuth())
	{
		api := API{}
		r.GET("", api.Get)
		r.POST("/list", api.List)
		r.DELETE("", api.Del)
	}
}
