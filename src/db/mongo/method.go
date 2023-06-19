package mongo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 分页限制
var PageLimit int64 = 30

type DB[T any] struct {
}

// 根据ID查找
func (x *DB[T]) Find(id primitive.ObjectID) (T, error) {
	model := new(T)
	filter := bson.M{"_id": id}
	collection := GetCollection(x.s2n(*model))
	err := collection.FindOne(context.Background(), filter).Decode(model)
	return *model, err
}

// 自定义过滤条件
func (x *DB[T]) Filter(filter primitive.M) (T, error) {
	model := new(T)
	if filter == nil {
		filter = bson.M{}
	}
	collection := GetCollection(x.s2n(*model))
	err := collection.FindOne(context.Background(), filter /*不能为nil*/).Decode(model)
	return *model, err
}

func (x *DB[T]) Exist(filter primitive.M) bool {
	if filter == nil {
		filter = bson.M{}
	}
	c, err := x.Count(filter)
	if err != nil {
		log.Println(err.Error())
	}
	return c > 0
}

// 新增1条数据
func (x *DB[T]) Add(data T) (primitive.ObjectID, error) {
	collection := GetCollection(x.s2n(data))
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// 更新1条记录
func (x *DB[T]) Update(id primitive.ObjectID, data T) error {

	collection := GetCollection(x.s2n(data))
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": data,
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// 删除一条记录
func (x *DB[T]) Del(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	collection := GetCollection(x.s2n(*new(T)))
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

// 删除多条记录
func (x *DB[T]) DelMany(filter primitive.M) error {
	collection := GetCollection(x.s2n(*new(T)))
	_, err := collection.DeleteMany(context.Background(), filter)
	return err
}

// 获取列表
func (x *DB[T]) List(filter primitive.M, index, limit, desc int64) (*[]T, error) {
	collection := GetCollection(x.s2n(*new(T)))
	if filter == nil {
		filter = bson.M{}
	}
	results := []T{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	options := options.Find().SetLimit(limit).SetSkip(limit * index)
	if desc == 1 || desc == -1 {
		options.SetSort(bson.M{"_id": desc})
	}
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return &results, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := new(T)
		err := cursor.Decode(result)
		if err != nil {
			return &results, err
		}
		results = append(results, *result)
	}

	if err := cursor.Err(); err != nil {
		return &results, err
	}

	return &results, err
}

// 获取列表
func (x *DB[T]) ListOrderBy(filter primitive.M, orderby primitive.M, index, limit int64) (*[]T, error) {
	collection := GetCollection(x.s2n(*new(T)))
	if filter == nil {
		filter = bson.M{}
	}
	if orderby == nil {
		orderby = bson.M{"_id": -1}
	}
	results := []T{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	options := options.Find().SetSort(orderby).SetLimit(limit).SetSkip(limit * index)
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return &results, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := new(T)
		err := cursor.Decode(result)
		if err != nil {
			return &results, err
		}
		results = append(results, *result)
	}

	if err := cursor.Err(); err != nil {
		return &results, err
	}

	return &results, err
}

// 文档数量
func (x *DB[T]) Count(filter primitive.M) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}
	collection := GetCollection(x.s2n(*new(T)))
	return collection.CountDocuments(context.Background(), filter)
}

// 分页
func (x *DB[T]) Paging(filter bson.M, index, limit, desc int64) (*Page, error) {
	page := Page{}
	if filter == nil {
		filter = bson.M{}
	}

	data, err := x.List(filter, index, limit, desc)
	if err != nil {
		return &page, nil
	}
	total, err := x.Count(filter)
	if err != nil {
		return &page, nil
	}

	page.Data = data
	page.Total = total

	return &page, err
}

// 将struct的名称转化为数据库表的名称
func (x *DB[T]) s2n(m T) string {
	return strings.ToLower(fmt.Sprintf("%T", m))
}

// 分页
type Page struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}