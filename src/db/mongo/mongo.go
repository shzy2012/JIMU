package mongo

import (
	"context"
	"sync"
	"time"
	"tolo/src/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/shzy2012/common/log"
)

const timeout = 30 * time.Second

var (
	mux sync.Mutex

	// mongodb client
	mongoClient *mongo.Client

	// mongodb Collection
	mongoDB *mongo.Database
)

func MongoClient() *mongo.Client {

	if mongoClient != nil {
		return mongoClient
	}

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

	return mongoClient
}

// GetCollection 获取 Collection
func GetCollection(name string) *mongo.Collection {
	mux.Lock()
	defer mux.Unlock()
	if mongoDB == nil {
		//初始化链接数据库
		mongoDB = MongoClient().Database(config.MongoDB.Database)
	}
	return mongoDB.Collection(name)
}

// New Hex
func GetHex() string {
	return primitive.NewObjectID().Hex()
}

// ID=>Hex
func FromHex(id string) primitive.ObjectID {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Infof("id=>[%s] is %s\n", id, err.Error())
	}
	return ID
}

// IDs=>Hexs
func FromHexs(ids []string) []primitive.ObjectID {
	res := make([]primitive.ObjectID, 0)
	for _, id := range ids {
		ID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Infof("id=>[%s] is %s\n", id, err.Error())
			continue
		}
		res = append(res, ID)
	}
	return res
}

func ToHexs(ids []primitive.ObjectID) []string {
	res := make([]string, len(ids))
	for i, id := range ids {
		res[i] = id.Hex()
	}
	return res
}

// New Hex
func GetObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}
