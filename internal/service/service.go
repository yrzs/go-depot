package service

import (
	"context"
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
	svc.dao = dao.New(global.DB)
	return svc
}
