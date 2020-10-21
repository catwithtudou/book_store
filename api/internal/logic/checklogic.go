package logic

import (
	"context"

	"book_store/api/internal/svc"
	"book_store/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx context.Context
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	return &types.CheckResp{}, nil
}
