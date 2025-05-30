package logic

import (
	"context"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login() (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
