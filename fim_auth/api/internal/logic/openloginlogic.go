package logic

import (
	"context"
	"errors"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	"fim_server/fim_auth/auth_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils/jwts"
	"fim_server/utils/open_login"

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

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	type OpenInfo struct {
		Nickname string
		Avatar   string
		OpenID   string
	}
	var info OpenInfo
	switch req.Flag {
	case "qq":
		qqInfo, openError := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		info = OpenInfo{
			Nickname: qqInfo.Nickname,
			Avatar:   qqInfo.Avatar,
			OpenID:   qqInfo.OpenID,
		}
		err = openError
	default:
		err = errors.New("暂时不支持该方式登录")
	}
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id=?", info.OpenID).Error
	if err != nil {
		//注册服务
		logx.Info("用户不存在，进行注册调用注册服务")
		res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
			NickName:           info.Nickname,
			Password:           "",
			Role:               2,
			Avatar:             info.Avatar,
			OpenId:             info.OpenID,
			RegistrationSource: "qq",
		})
		if err != nil {
			logx.Error(err)
			return nil, errors.New("登录失败")
		}
		user.Model.ID = uint(res.UserId)
		user.Role = 2
		user.NickName = info.Nickname
	}
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   user.Model.ID,
		NickName: user.NickName,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		logx.Error(err)
		err = errors.New("服务器内部错误")
		return
	}

	return &types.LoginResponse{
		Token: token,
	}, nil
}
