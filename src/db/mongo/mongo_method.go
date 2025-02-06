package mongo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 分页限制
var PageLimit int64 = 30

type DB[T any] struct {
	table string
}

// 根据ID查找
func (x *DB[T]) Find(id primitive.ObjectID) (T, error) {
	model := new(T)
	filter := bson.M{"_id": id}
	collection := GetCollection(x.tableName())
	err := collection.FindOne(context.Background(), filter).Decode(model)
	return *model, err
}

// 自定义过滤条件
func (x *DB[T]) Filter(filter primitive.M, params ...int) (T, error) {
	model := new(T)
	if filter == nil {
		filter = bson.M{}
	}

	option := options.FindOne()
	if len(params) > 0 {
		option.SetSort(bson.M{"_id": params[0]})
	}

	collection := GetCollection(x.tableName())
	err := collection.FindOne(context.Background(), filter /*不能为nil*/, option).Decode(model)
	return *model, err
}

// 新增1条数据
func (x *DB[T]) Add(data T) (primitive.ObjectID, error) {
	collection := GetCollection(x.tableName())
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// 更新1条记录
func (x *DB[T]) Update(id primitive.ObjectID, data T) error {

	collection := GetCollection(x.tableName())
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
	collection := GetCollection(x.tableName())
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

// 删除多条记录
func (x *DB[T]) DelMany(filter primitive.M) error {
	collection := GetCollection(x.tableName())
	_, err := collection.DeleteMany(context.Background(), filter)
	return err
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

/*
	获取列表

filter:过滤条件
index:分页索引
limit:分页大小
desc:排序方式 1:正序,-1:倒序
*/
func (x *DB[T]) List(filter primitive.M, index, limit, desc int64) (*[]T, error) {
	collection := GetCollection(x.tableName())
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

/*
	获取列表

filter:过滤条件. eg: bson.M{"name": "zs"}
orderby:自定义字段排序.eg: bson.M{"_id": -1}
index:分页索引
limit:分页大小
desc:排序方式 1:正序,-1:倒序
*/
func (x *DB[T]) ListBy(filter, orderby primitive.M, index, limit int64) (*[]T, error) {
	collection := GetCollection(x.tableName())
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
	collection := GetCollection(x.tableName())
	return collection.CountDocuments(context.Background(), filter)
}

// 字段去重
func (x *DB[T]) Distinct(fieldName string, filter primitive.M) ([]interface{}, error) {
	if filter == nil {
		filter = bson.M{}
	}
	collection := GetCollection(x.tableName())
	return collection.Distinct(context.Background(), fieldName, filter)
}

// 分页
func (x *DB[T]) Paging(filter bson.M, index, limit, desc int64) (Page, error) {

	page := Page{
		Data: &[]T{},
	}
	if filter == nil {
		filter = bson.M{}
	}

	data, err := x.List(filter, index, limit, desc)
	if err != nil {
		return page, nil
	}
	total, err := x.Count(filter)
	if err != nil {
		return page, nil
	}

	page.Data = data
	page.Total = total

	return page, err
}

func (x *DB[T]) PagingBy(filter, orderby primitive.M, index, limit int64) (Page, error) {

	page := Page{
		Data: &[]T{},
	}
	if filter == nil {
		filter = bson.M{}
	}

	data, err := x.ListBy(filter, orderby, index, limit)
	if err != nil {
		return page, nil
	}
	total, err := x.Count(filter)
	if err != nil {
		return page, nil
	}

	page.Data = data
	page.Total = total

	return page, err
}

// 创建索引
// https://kb.objectrocket.com/mongo-db/how-to-create-an-index-using-the-golang-driver-for-mongodb-455
func (x *DB[T]) IndexCreate(field string, sort int /*1 自然排序(默认方式) || -1 倒序*/, unique bool) error {
	collection := GetCollection(x.tableName())
	log.Printf("index: %s->%s\n", collection.Name(), field)
	// db.members.createIndex( { "SOME_FIELD": 1 }, { unique: true } )
	mod := mongo.IndexModel{
		Keys: bson.M{
			field: sort, // 1 自然排序(默认方式) || -1 倒序
		}, Options: options.Index().SetUnique(unique),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), mod)
	return err
}

// 检查指定索引是否存在
func (x *DB[T]) IndexExists(indexName string) (bool, error) {
	// 获取当前集合的索引列表
	collection := GetCollection(x.tableName())
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		return false, err
	}
	defer cursor.Close(context.TODO())

	// 遍历索引列表
	for cursor.Next(context.TODO()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			return false, err
		}
		// 检查索引的名称是否匹配
		if index["name"] == indexName {
			return true, nil
		}
	}
	return false, nil
}

// 将struct的名称转化为数据库表的名称
func (x *DB[T]) tableName() string {
	if x.table == "" {
		x.table = strings.ToLower(fmt.Sprintf("%T", *new(T)))
	}
	return x.table
}

// 分页
type Page struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
