package user

import (
	"tolo/src/tools"

	"github.com/gin-gonic/gin"
	"github.com/shzy2012/common/log"
)

type API struct {
}

func (x *API) AddOrUpdate(c *gin.Context) {

	param := ReqDto{}
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(412, err.Error())
		return
	}
	log.Printf("[param]=>%+v\n", param)

	c.JSON(200, "")
}

func (x *API) Get(c *gin.Context) {
	id := c.Query("id")
	log.Printf("[param]=>%s\n", id)
	if tools.IsEmpty(id) {
		c.JSON(412, "")
		return
	}
	c.JSON(200, "")
}

func (x *API) Del(c *gin.Context) {
	id := c.Query("id")
	log.Printf("[param]=>%s\n", id)
	if tools.IsEmpty(id) {
		c.JSON(412, "id不能为空")
		return
	}

	c.JSON(200, "")
}

func (x *API) List(c *gin.Context) {
	param := Params{}
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(412, err.Error())
		return
	}
	log.Printf("[param]=>%+v\n", param)

	if param.Page < 1 {
		param.Page = 1
	}

	if param.Size > 100 {
		param.Size = 100
	}

	c.JSON(200, "")
}
