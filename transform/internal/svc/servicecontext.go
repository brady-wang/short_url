package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"shorturl/transform/internal/config"
	"shorturl/transform/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.ShorturlModel   // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource)), // 手动代码
	}
}
