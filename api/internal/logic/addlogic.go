package logic

import (
	"context"

	"book_store/api/internal/svc"
	"book_store/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx context.Context
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	return &types.AddResp{}, nil
}
