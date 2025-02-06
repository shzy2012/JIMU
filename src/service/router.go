package service

import (
	"jimu/src/config"
	"jimu/src/middleware/cors"
	"jimu/src/middleware/jwt"
	"jimu/src/service/example"
	"jimu/src/service/user"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	gin.SetMode(config.Server.RunMode)
	r := gin.New()
	// 中间件
	r.Use(jwt.Log()) //请求日志
	// 开启跨域
	r.Use(cors.Middleware())

	// 静态服务
	r.Use(static.Serve("/", static.LocalFile("www", true))) // true:自动返回index.html
	r.NoRoute(func(c *gin.Context) {
		/*
			1.解决前端history路由404(缺失index.html)
			2.前端路由遇到错误会自动重定向到 /
		*/
		accept := c.Request.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			var err error
			var content []byte
			content, err = os.ReadFile("www/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})

	//健康检查
	r.GET("/h", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "i am working",
		})
	})

	// 登录
	r.POST("/login", func(c *gin.Context) {
		user.LoginHandle(c)
	})

	// routes
	g1 := r.Group("/api/v1")
	g1.Use(cors.Middleware()) //开启跨域

	example.Routes(g1)

	return r
}
