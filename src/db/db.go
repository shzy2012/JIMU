package db

// Init 初始化
func Init() {
	// 初始化用户索引
	InitUserIndex()

	// 初始化用户数据
	InitUserDate()
}
