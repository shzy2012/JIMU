package db

import (
	"fmt"
	"time"

	"jimu/src/db/mongo"
	"jimu/src/tools"

	"github.com/shzy2012/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const DEFAULT_PASSWD = "@123"

type Role string

const (
	Admin    Role = "admin"
	Operator Role = "operator"
	Common   Role = "common"
)

// 用户信息
type User struct {
	mongo.DB[User] `json:"-" bson:"-"`
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`      //系统ID
	Phone          string             `json:"phone" bson:"phone"`           //手机号码
	Passwd         string             `json:"passwd" bson:"passwd"`         //密码
	Role           Role               `json:"role" bson:"role"`             //手机号码
	Token          string             `json:"token" bson:"token"`           //安全 token（忽略）
	LoginC         int                `json:"loginc" bson:"loginc"`         //登录次数.超过60次锁定,只有管理员才能解锁
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"` //创建时间
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"` //更新时间
}

func NewUser() *User {
	return &User{}
}

// 初始化数据
func initUser() {
	log.Infoln("初始化 user 数据")
	ids := []string{"admin"}
	for _, ele := range ids {
		c, _ := NewUser().Count(bson.M{"phone": ele})
		if c <= 0 {
			model := User{
				Phone:     ele,
				Role:      Admin,
				Passwd:    tools.Sha256(fmt.Sprintf("%s%s", ele, DEFAULT_PASSWD)),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			model.Add(model)
		}
	}

	log.Info("user data is ok")
}
