// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileInfoDao is the data access object for the table file_info.
type FileInfoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FileInfoColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FileInfoColumns defines and stores column names for the table file_info.
type FileInfoColumns struct {
	Id         string // 文件ID
	Name       string // 文件名字
	Url        string // 七牛云URL
	UploaderId string // 上传者ID
	CreatedAt  string //
	DeletedAt  string //
}

// fileInfoColumns holds the columns for the table file_info.
var fileInfoColumns = FileInfoColumns{
	Id:         "id",
	Name:       "name",
	Url:        "url",
	UploaderId: "uploader_id",
	CreatedAt:  "created_at",
	DeletedAt:  "deleted_at",
}

// NewFileInfoDao creates and returns a new DAO object for table data access.
func NewFileInfoDao(handlers ...gdb.ModelHandler) *FileInfoDao {
	return &FileInfoDao{
		group:    "default",
		table:    "file_info",
		columns:  fileInfoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FileInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FileInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FileInfoDao) Columns() FileInfoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FileInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FileInfoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FileInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
