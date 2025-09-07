// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsImagesDao is the data access object for the table goods_images.
type GoodsImagesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GoodsImagesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GoodsImagesColumns defines and stores column names for the table goods_images.
type GoodsImagesColumns struct {
	Id      string //
	Url     string // 七牛云url
	GoodsId string // 商品ID
	FileId  string // 文件ID（关联file_info）
	Sort    string // 排序
}

// goodsImagesColumns holds the columns for the table goods_images.
var goodsImagesColumns = GoodsImagesColumns{
	Id:      "id",
	Url:     "url",
	GoodsId: "goods_id",
	FileId:  "file_id",
	Sort:    "sort",
}

// NewGoodsImagesDao creates and returns a new DAO object for table data access.
func NewGoodsImagesDao(handlers ...gdb.ModelHandler) *GoodsImagesDao {
	return &GoodsImagesDao{
		group:    "default",
		table:    "goods_images",
		columns:  goodsImagesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GoodsImagesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GoodsImagesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GoodsImagesDao) Columns() GoodsImagesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GoodsImagesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GoodsImagesDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *GoodsImagesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
