package db

import (
	"time"

	"jimu/src/db/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Example struct {
	mongo.DB[Example] `json:"-" bson:"-"`
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`    //系统ID
	Name              string             `json:"name" bson:"name"`           //Name
	Desc              string             `json:"desc" bson:"desc"`           //设备描述
	CreateAt          time.Time          `json:"create_at" bson:"create_at"` //创建时间
	UpdateAt          time.Time          `json:"update_at" bson:"update_at"` //修改时间
}

func NewExample() *Example {
	return &Example{}
}

func InitExampleIndex() {
	// 创建索引
	indexNames := []string{"name", "create_at"}
	for i := 0; i < len(indexNames); i++ {
		indexName := indexNames[i]
		exist, _ := NewExample().IndexExists(indexName, 1)
		if exist {
			continue
		}

		NewExample().IndexCreate(indexName, 1, false)
	}
}
