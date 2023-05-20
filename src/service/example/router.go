package example

import (
	"github.com/gin-gonic/gin"
)

// Routes initialize routing information
func Routes(route *gin.RouterGroup) {
	r := route.Group("/example")
	r.Use()
	{
		api := API{}
		r.GET("", api.Get)
		r.POST("", api.AddOrUpdate)
		r.POST("/list", api.List)
		r.DELETE("", api.Del)
	}
}
