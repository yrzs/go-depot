package service

import (
	"context"
	opentracingGorm "github.com/yrzs/opentracing-gorm"
	"go-depot/global"
	"go-depot/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	// new Dao
	svc.dao = dao.New(
		opentracingGorm.WithContext(svc.ctx, global.DB), // 数据库连接实例的上下文信息注册
	)
	return svc
}
