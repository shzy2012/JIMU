package jwt

import (
	"fmt"
	"jimu/src/tools"

	"github.com/gin-gonic/gin"
	"github.com/shzy2012/common/log"
)

// Log 中间件
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		request := fmt.Sprintf("remote=>%s host=>%s method=>%s url=>%s content_length=>%v", r.RemoteAddr, r.Host, r.Method, r.URL, r.ContentLength)
		log.Infoln(request)

		c.Next()
	}
}

// JWTAuth 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("token")
		log.Println("Request Token: ", token)
		if token == "" {
			// 类似于403 Forbidden，401语义即“未认证”，即用户没有必要的凭据。
			c.JSON(401, "无权限访问")
			c.Abort()

			log.Info("请求未携带token，无权限访问")
			return
		}

		// 解析token中包含的相关信息(有效载荷)
		claims, err := tools.ParseToken(token)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			log.Info(err.Error())
			return
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)
	}
}
