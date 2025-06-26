package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/internal/svc"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *UserInfoLogic) UserInfo(in *user_rpc.UserInfoRequest) (*user_rpc.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	var user user_models.UserModel
	fmt.Println("查询的用户id为", in.UserId)
	err := l.svcCtx.DB.Preload("UserConfModel").Take(&user, in.UserId).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	fmt.Println("查询成功")
	byteData, _ := json.Marshal(user)
	fmt.Println(byteData)
	fmt.Println(string(byteData))
	return &user_rpc.UserInfoResponse{Data: byteData}, nil
}
