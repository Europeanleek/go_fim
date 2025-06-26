package logic

import (
	"context"
	"errors"
	"fmt"

	"fim_server/fim_auth/api/internal/svc"
	"fim_server/fim_auth/api/internal/types"
	"fim_server/fim_auth/auth_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"
	"fim_server/utils/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "nick_name = ?", req.NickName).Error
	if err == nil && user.NickName != "" {
		err = errors.New("用户已经注册")
		return
	}
	res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
		NickName:           req.NickName,
		Password:           pwd.HashPwd(req.Pwd),
		Role:               2,
		Avatar:             "",
		OpenId:             "",
		RegistrationSource: "offical_register",
	})
	if err != nil {
		return nil, errors.New("注册失败")
	}
	data := fmt.Sprintf("账号厨厕成功，用户id为%d", res.UserId)
	return &types.Response{
		Data: data,
	}, nil
}
