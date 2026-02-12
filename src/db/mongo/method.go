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

/************************************************
* 通用操作方法
*************************************************/

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

// First 获取集合中的第一条记录
func (x *DB[T]) First() (T, error) {
	return x.Filter(nil, 1) // 使用 1 表示正序，获取第一条
}

// Last 获取集合中的最后一条记录
func (x *DB[T]) Last() (T, error) {
	return x.Filter(nil, -1) // 使用 -1 表示倒序，获取最后一条
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

// 批量插入数据
func (x *DB[T]) AddBatch(data []T) error {
	if len(data) == 0 {
		return nil
	}
	collection := GetCollection(x.tableName())
	// 转换为 []interface{} 类型
	documents := make([]interface{}, len(data))
	for i := range data {
		documents[i] = data[i]
	}
	_, err := collection.InsertMany(context.Background(), documents)
	return err
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

// UpdateBy 根据自定义过滤条件更新单条记录
/*
filter := bson.M{
	"_id":       model.ID,
	"is_delete": bson.M{"$ne": true},
}
modified, err := db.NewOrder().UpdateBy(filter, model)
*/
func (x *DB[T]) UpdateBy(filter bson.M, data T) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}
	collection := GetCollection(x.tableName())
	update := bson.M{
		"$set": data,
	}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

// UpdateManyBy 根据自定义过滤条件更新多条记录
/*
filter := bson.M{"status": 1}
data := bson.M{"status": 2}
modified, err := db.NewUser().UpdateManyBy(filter, data)
*/
func (x *DB[T]) UpdateMany(filter bson.M, data T) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}
	collection := GetCollection(x.tableName())
	update := bson.M{
		"$set": data,
	}
	res, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

// UpdateManyRaw 根据自定义过滤条件和原始更新文档更新多条记录
/*
filter := bson.M{"order_no": "XSCKD2026", "status": "scanned"}
update := bson.M{
	"$set": bson.M{
		"status":    "outstocked",
		"update_at": time.Now(),
	},
}
modified, err := db.NewQrCode().UpdateManyRaw(filter, update)
*/
func (x *DB[T]) UpdateManyRaw(filter bson.M, update bson.M) (int64, error) {
	if filter == nil {
		filter = bson.M{}
	}
	collection := GetCollection(x.tableName())
	res, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
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
func (x *DB[T]) List(filter primitive.M, index, limit, desc int64) ([]T, error) {
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
		return results, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := new(T)
		err := cursor.Decode(result)
		if err != nil {
			return results, err
		}
		results = append(results, *result)
	}

	if err := cursor.Err(); err != nil {
		return results, err
	}

	return results, err
}

/*
	获取列表

filter:过滤条件. eg: bson.M{"name": "zs"}
orderby:自定义字段排序，必须使用 bson.D 类型
  - 单字段排序: bson.D{{Key: "_id", Value: -1}}
  - 多字段排序: bson.D{{Key: "field1", Value: 1}, {Key: "field2", Value: -1}}
    注意：MongoDB 多字段排序必须使用 bson.D 而不是 bson.M，因为 map 不保证顺序

index:分页索引
limit:分页大小
*/
func (x *DB[T]) ListBy(filter primitive.M, orderby bson.D, index, limit int64) ([]T, error) {
	collection := GetCollection(x.tableName())
	if filter == nil {
		filter = bson.M{}
	}
	if len(orderby) == 0 {
		orderby = bson.D{{Key: "_id", Value: -1}}
	}
	results := []T{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	options := options.Find().SetSort(orderby).SetLimit(limit).SetSkip(limit * index)
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return results, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := new(T)
		err := cursor.Decode(result)
		if err != nil {
			return results, err
		}
		results = append(results, *result)
	}

	if err := cursor.Err(); err != nil {
		return results, err
	}

	return results, err
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
		Data: []T{},
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

/*
	分页查询（支持自定义排序）

filter:过滤条件. eg: bson.M{"name": "zs"}
orderby:自定义字段排序，必须使用 bson.D 类型
  - 单字段排序: bson.D{{Key: "_id", Value: -1}}
  - 多字段排序: bson.D{{Key: "field1", Value: 1}, {Key: "field2", Value: -1}}
    注意：MongoDB 多字段排序必须使用 bson.D 而不是 bson.M，因为 map 不保证顺序

index:分页索引
limit:分页大小
*/
func (x *DB[T]) PagingBy(filter primitive.M, orderby bson.D, index, limit int64) (Page, error) {

	page := Page{
		Data: []T{},
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

/*
pipeline:聚合管道
返回:聚合结果

example:

	filter := bson.M{"reply_content": "兜底回复"}
	pipeline := []bson.M{
		{"$match": filter},
		{"$group": bson.M{
			"_id":   "$belong_to",
			"count": bson.M{"$sum": 1},
		}},
	}
	results, err := db.NewChatMessage().Aggregate(pipeline)
	if err != nil {
		log.Printf("[err]=>%s\n", err.Error())
	}

	for _, v := range results {
		fmt.Println(v["_id"], v["count"])
	}

	fmt.Printf("%s\n", tools.ToBytes(results))

----------------------------------------------------------------

聚合结果处理方式:
// 方式 1: 定义具体的结构体（类型安全）

	type AggregateResult struct {
		ID          string  `bson:"_id"`
		TotalAmount float64 `bson:"total_amount"`
		Count       int     `bson:"count"`
		AvgAmount   float64 `bson:"avg_amount"`
	}

// 方式 2: 使用 bson.M (map[string]interface{}) - 最灵活
// var results []bson.M

// 方式 3: 使用 bson.D (有序的键值对) - 保持字段顺序
// var results []bson.D

// 方式 4: 使用 interface{} - 完全动态
// var results []interface{}
*/
func (x *DB[T]) Aggregate(pipeline interface{}) ([]bson.M, error) {
	collection := GetCollection(x.tableName())
	ctx := context.Background()
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	results := []bson.M{}
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// 指定删除某个字段
func (x *DB[T]) FieldDrop(filter bson.M, field string) error {
	collection := GetCollection(x.tableName())
	if filter == nil {
		filter = bson.M{}
	}

	//  更新多个字段
	update := bson.M{
		"$unset": bson.M{
			field: "",
		},
	}

	_, err := collection.UpdateMany(context.Background(), filter, update)
	return err
}

// 统一生成索引名称
// field: 字段名
// sort: 排序方式 1:正序,-1:倒序
func GetIndexName(field string, sort int) string {
	if sort == 1 {
		return field + "_1"
	}
	return field + "_-1"
}

// 创建索引
// https://kb.objectrocket.com/mongo-db/how-to-create-an-index-using-the-golang-driver-for-mongodb-455
func (x *DB[T]) IndexCreate(field string, sort int /*1 自然排序(默认方式) || -1 倒序*/, unique bool) error {
	collection := GetCollection(x.tableName())
	log.Printf("index: %s->%s\n", collection.Name(), field)
	// db.members.createIndex( { "SOME_FIELD": 1 }, { unique: true } )
	// MongoDB默认索引名称格式: 字段名_排序方向
	indexName := GetIndexName(field, sort)
	mod := mongo.IndexModel{
		Keys: bson.M{
			field: sort, // 1 自然排序(默认方式) || -1 倒序
		}, Options: options.Index().SetUnique(unique).SetName(indexName),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), mod)
	return err
}

// 创建带过期时间的索引
// field: 索引字段
// sort: 排序方式 1:正序,-1:倒序
// expireAfterSeconds: 过期时间(秒)
func (x *DB[T]) IndexCreateWithExpiry(field string, sort int, expireAfterSeconds int32) error {
	collection := GetCollection(x.tableName())
	log.Printf("index with expiry: %s->%s, expireAfterSeconds: %d\n", collection.Name(), field, expireAfterSeconds)

	mod := mongo.IndexModel{
		Keys: bson.M{
			field: sort,
		},
		Options: options.Index().SetExpireAfterSeconds(expireAfterSeconds),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), mod)
	return err
}

// 检查指定索引是否存在（通过完整索引名）
func (x *DB[T]) IndexExists(indexName string, sort int) (bool, error) {

	indexName = GetIndexName(indexName, sort)

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
		// 将点号替换为下划线，例如 db.edgetype -> db_edgetype
		x.table = strings.ReplaceAll(x.table, ".", "_")
	}
	return x.table
}

// 分页
type Page struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
