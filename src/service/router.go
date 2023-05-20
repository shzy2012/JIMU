package service

import (
	"jimu/src/config"
	"jimu/src/middleware/cors"
	"jimu/src/middleware/jwt"
	"jimu/src/service/example"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	gin.SetMode(config.Serve.RunMode)
	r := gin.New()
	// 中间件
	r.Use(jwt.Log()) //请求日志
	// 开启跨域
	r.Use(cors.Middleware())
	// 性能
	pprof.Register(r)

	//健康检查
	r.GET("/h", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "i am working",
		})
	})

	// routes
	g1 := r.Group("/api/v1")
	g1.Use(cors.Middleware()) //开启跨域

	example.Routes(g1)

	return r
}
