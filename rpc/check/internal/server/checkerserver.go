// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package server

import (
	"context"

	"book_store/rpc/check/internal/logic"
	"book_store/rpc/check/internal/svc"
	check "book_store/rpc/check/pb"
)

type CheckerServer struct {
	svcCtx *svc.ServiceContext
}

func NewCheckerServer(svcCtx *svc.ServiceContext) *CheckerServer {
	return &CheckerServer{
		svcCtx: svcCtx,
	}
}

func (s *CheckerServer) Check(ctx context.Context, in *check.CheckReq) (*check.CheckResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
