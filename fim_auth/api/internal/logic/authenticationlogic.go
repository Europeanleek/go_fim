package logic

import (
	"context"
	"errors"
	"fmt"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	comparelist "fim_server/utils/compare_list"
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

func (l *AuthenticationLogic) Authentication(req *types.AuthRequest) (resp *types.AuthResponse, err error) {
	// todo: add your logic here and delete this line
	token := req.Token
	origin_url := req.OriginUrl
	fmt.Println(token, origin_url)
	is_find_url := comparelist.ContainsStringByRegex(l.svcCtx.Config.WhiteList, origin_url)
	if is_find_url {
		return nil, nil
	}

	if token == "" {
		err = errors.New("token为空，认证失败")
		return
	}
	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("parse错误，认证失败")
		return
	}

	_, err = l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("logout_%d_%s", payload.UserID, token)).Result()
	if err == nil {
		err = errors.New("redis中存在注销标记，认证失败")
		return
	}
	resp = &types.AuthResponse{
		UserID: payload.UserID,
		Role:   int(payload.Role),
	}
	err = nil
	return
}
