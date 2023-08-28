package mongo

import (
	"context"
	"jimu/src/config"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/shzy2012/common/log"
)

const timeout = 30 * time.Second

// mongodb client
var mongoClient *mongo.Client
var mux sync.Mutex

func init() {

	var err error
	//初始化 MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//mongodb的用户名和密码是基于特定数据库的，而不是基于整个系统的。所有所有数据库db都需要设置密码
	//mongodb://youruser2:yourpassword2@localhost/yourdatabase

	connectString := config.MongoDB.URI
	log.Infoln("[mongo URI]=>", connectString)
	if config.MongoDB.Username != "" {
		// 设置登录凭证
		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(connectString).SetAuth(options.Credential{
			Username: config.MongoDB.Username,
			Password: config.MongoDB.Password,
		}))
	} else {
		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	}
	if err != nil {
		log.Fatalf("[mongo]=>%s\n", err)
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	if err == nil {
		log.Println("[mongo]=>int ok.")
	} else {
		log.Fatalf("[mongo]=>int fail. %s\n", err.Error())
	}
}

// Mongo Collection
var (
	mongoDB *mongo.Database
)

// GetCollection 获取 Collection
func GetCollection(name string) *mongo.Collection {
	mux.Lock()
	if mongoDB == nil {
		//初始化链接数据库
		mongoDB = mongoClient.Database(config.MongoDB.Database)
	}
	defer mux.Unlock()
	return mongoDB.Collection(name)
}

// ID=>Hex
func FromHex(id string) primitive.ObjectID {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Infof("id=>[%s] is %s\n", id, err.Error())
	}
	return ID
}

// New Hex
func GetHex() string {
	return primitive.NewObjectID().Hex()
}

// New Hex
func GetObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

// 初始化索引
func InitIndexes() {
	createIndexes()
}

// createIndexes 创建索引
func createIndexes() {

	/************************************
	*  创建 xxx 索引
	*************************************/

	// Declare an index model object to pass to CreateOne()
	// db.members.createIndex({"SOME_FIELD":1},{unique:true})
	mod := mongo.IndexModel{
		Keys: bson.M{
			"code": 1, // 1=> index in ascending order  | -1=> index in descending order
		}, Options: options.Index().SetUnique(true),
	}

	// Create an Index using the CreateOne() method
	collection := GetCollection("Collection_Name")
	ind, err := collection.Indexes().CreateOne(context.TODO(), mod)
	// Check if the CreateOne() method returned any errors
	if err != nil {
		log.Fatalf("ERROR:%s\n", err)
	} else {
		// API call returns string of the index name
		log.Println("Index:", ind, reflect.TypeOf(ind))
	}
}
