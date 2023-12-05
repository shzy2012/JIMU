package postgre

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DB[T any] struct {
}

// 查找1条数据
func (x *DB[T]) Find(query string, args ...interface{}) (*T, error) {
	model := new(T)
	if err := PGClient().NewSelect().Model(model).Limit(1).Where(query, args).Scan(context.TODO()); err != nil {
		return nil, err
	}

	return model, nil
}

// "name LIKE '%foo%'"
func (x *DB[T]) Exist(query string, args ...interface{}) bool {
	exists, err := PGClient().NewSelect().Model((*T)(nil)).Where(query, args).Exists(context.TODO())
	if err != nil {
		panic(err)
	}

	return exists
}

// 新增1条数据
func (x *DB[T]) Add(data T) (int64, error) {
	res, err := PGClient().NewInsert().Model(data).Exec(context.TODO())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 根据ID更新
func (x *DB[T]) Update(id int64, data T) (int64, error) {
	res, err := PGClient().NewUpdate().Model(data).Where("id = ?", id).Exec(context.TODO())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 根据主键删除
func (x *DB[T]) Del() (int64, error) {
	res, err := PGClient().NewDelete().Model(x).WherePK().Exec(context.TODO())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// 删除多条记录
func (x *DB[T]) DelMany(query string, args ...interface{}) (int64, error) {
	res, err := PGClient().NewDelete().Model(x).Where(query, args).Exec(context.TODO())
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

/*
	获取列表

filter:过滤条件
index:分页索引
limit:分页大小
desc:排序方式 1:正序,-1:倒序
*/
func (x *DB[T]) List(query string, index, limit int) (*[]T, error) {
	var models []T
	err := PGClient().NewSelect().Model(&models).Where(query).Offset(index).Limit(limit).Order().Scan(context.TODO())
	if err != nil {
		return nil, err
	}
	return &models, nil
}

// 文档数量
func (x *DB[T]) Count(filter primitive.M) (int, error) {
	count, err := PGClient().NewSelect().Model((*T)(nil)).Count(context.TODO())
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 分页
func (x *DB[T]) Paging(query string, index, limit int) (*Page, error) {
	page := Page{}
	var models []T
	total, err := PGClient().NewSelect().Model(&models).Limit(limit).ScanAndCount(context.TODO())
	if err != nil {
		panic(err)
	}

	page.Data = models
	page.Total = total
	return &page, err
}

type Page struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

/************************************************/

func (x *DB[T]) CreateTable() error {
	// Create  table.
	_, err := PGClient().NewCreateTable().IfNotExists().Model((*T)(nil)).Exec(context.TODO())
	return err
}

func (x *DB[T]) DropTable() error {
	_, err := PGClient().NewDropTable().Model((*T)(nil)).Exec(context.TODO())
	return err
}
