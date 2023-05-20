package db

import "jimu/src/db/mongo"

// Init 初始化
func Init() {
	mongo.InitIndexes()
}
