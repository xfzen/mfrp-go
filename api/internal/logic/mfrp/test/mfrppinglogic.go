package test

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mfrp/api/internal/svc"
)

type MfrpPingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMfrpPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MfrpPingLogic {
	return &MfrpPingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MfrpPingLogic) MfrpPing() error {
	// todo: add your logic here and delete this line

	return nil
}
