package logic

import (
	"context"
	"errors"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	"fim_server/fim_auth/auth_models"
	"fim_server/utils/jwts"
	"fim_server/utils/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	// fmt.Println(req.UserName, req.Password)
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nick_name = ?", req.UserName).Error
	if err != nil {
		err = errors.New("用户名或者密码错误")
		return
	}
	if !pwd.CheckPwd(user.Pwd, req.Password) {
		err = errors.New("用户名或者密码错误")
		return
	}
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.ID,
		NickName: user.NickName,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		err = errors.New("服务端出现问题,请稍后再试")
		return
	}
	return &types.LoginResponse{
		Token: token,
	}, nil
}
