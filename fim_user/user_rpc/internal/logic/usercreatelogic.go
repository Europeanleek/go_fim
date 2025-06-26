package logic

import (
	"context"
	"errors"
	"time"

	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, "open_id=? And nick_name=?", in.OpenId, in.NickName).Error
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	switch in.RegistrationSource {
	case "qq":
		user = user_models.UserModel{
			NickName:           in.NickName,
			OpenID:             in.OpenId,
			Avatar:             in.Avatar,
			Role:               int8(in.Role),
			RegistrationSource: in.RegistrationSource,
		}
	case "offical_register":
		user = user_models.UserModel{
			NickName:           in.NickName,
			Role:               int8(in.Role),
			Pwd:                in.Password,
			RegistrationSource: in.RegistrationSource,
		}
	default:
		return nil, errors.New("注册时源错误")
	}
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("创建用户失败")
	}
	user_conf := &user_models.UserConfModel{
		UserID:        user.ID,
		RecallMessage: nil,   //撤回消息需要提示的内容 撤回了一条消息
		FriendOnline:  false, // 关闭好友上线提醒
		Sound:         true,  //声音
		SecureLink:    false, //安全链接
		SavePwd:       false, //是否保存密码
		SearchUser:    2,     //2代表可以通过昵称搜索到我
		Verification:  2,     //2代表添加好友时需要验证消息
		Online:        true,  //在线状态，默认创建用户就是在线的
	}
	user_conf.CreatedAt = time.Now()
	user_conf.UpdateAt = time.Now()
	l.svcCtx.DB.Create(&user_conf)
	return &user_rpc.UserCreateResponse{UserId: int32(user.ID)}, nil
}
