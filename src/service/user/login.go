package user

import (
	"fmt"
	"jimu/src/db"
	"jimu/src/tools"

	"github.com/gin-gonic/gin"
	"github.com/shzy2012/common/log"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginHandle(c *gin.Context) {

	param := ReqDto{}
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(412, err.Error())
		return
	}
	log.Printf("[param]=>phone:%s\n", param.Phone)
	param.Phone = tools.Trim(param.Phone)

	if tools.IsEmpty(param.Phone) {
		msg := "请填写手机号码"
		log.Errorln(msg)
		c.JSON(412, msg)
		return
	}

	// 1.判断手机号码
	user, err := db.NewUser().Filter(bson.M{"phone": param.Phone})
	if err != nil || user.ID.IsZero() {
		msg := fmt.Sprintf("手机号码【%s】不存在\n", param.Phone)
		log.Errorln(msg)
		c.JSON(412, msg)
		return
	}

	// 2.是否锁定
	if user.LoginC > 60 {
		log.Printf("%s登录次数: %v\n", param.Phone, user.LoginC)
		c.JSON(412, "多次输入错误的密码已锁定，请联系管理员解锁。")
		return
	}

	// 3.判断密码
	if user.Passwd != param.Passwd {
		user.LoginC = user.LoginC + 1
		if err := user.Update(user.ID, user); err != nil {
			log.Printf("%s\n", err.Error())
		}

		log.Errorf("%s\n", "密码不正确")
		c.JSON(412, "密码不正确")
		return
	}

	// 5.更新LoginC
	user.LoginC = 0
	err = user.Update(user.ID, user)
	if err != nil {
		log.Errorf("%s\n", err.Error())
		c.JSON(412, err.Error())
		return
	}

	u := LoginResp(user)
	log.Printf("[LoginHandle]=>%s\n", tools.ToBytes(u))
	c.JSON(200, u)
}

// 统一登录响应
func LoginResp(user db.User) RespDto {

	token, err := tools.GenerateToken(user.ID.Hex())
	if err != nil {
		log.Errorf("[LoginResp]=>%s\n", err.Error())
	}

	return RespDto{
		ID:    user.ID.Hex(),
		Phone: user.Phone,
		Role:  string(user.Role),
		Token: token,
	}
}
