package svc

import (
	"book_store/rpc/add/internal/config"
	"book_store/rpc/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c config.Config
	Model *model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
	}
}
