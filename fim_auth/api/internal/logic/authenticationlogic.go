package logic

import (
	"context"
	"errors"
	"fmt"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	"fim_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	if token == "" {
		err = errors.New("认证失败")
		return
	}
	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}

	_, err = l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("logout_%d", payload.UserID)).Result()
	if err == nil {
		err = errors.New("认证失败")
		return
	}
	resp = &types.Response{
		Data: "认证通过",
	}
	err = nil
	return
}
