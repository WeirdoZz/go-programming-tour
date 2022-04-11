package service

import (
	"blog-service/global"
	"blog-service/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

// New 新建一个service服务引擎
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
