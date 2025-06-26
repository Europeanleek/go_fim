package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	"fim_server/utils/jwts"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	if token == "" {
		err = errors.New("请传入token")
		return
	}
	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	now := time.Now()
	expiration := payload.ExpiresAt.Time.Sub(now)
	key := fmt.Sprintf("logout_%d_%s", payload.UserID, token)
	l.svcCtx.Redis.SetNX(l.ctx, key, "", expiration)
	return
}
