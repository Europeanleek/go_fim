package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"
	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(req.UserID)
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
		UserId: uint32(req.UserID),
	})
	fmt.Println("res==", res)
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("res==", res)
		logx.Error(err)
		return nil, err
	}
	var user user_models.UserModel
	err = json.Unmarshal(res.Data, &user)
	fmt.Println("user==", user)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据错误")
	}
	resp = &types.UserInfoResponse{
		UserID:   user.ID,
		Nickname: user.NickName,
		Abstract: user.Abstract,
		Avatar:   user.Avatar,
	}
	if user.UserConfModel != nil {
		resp.FriendOnline = user.UserConfModel.FriendOnline
		resp.Sound = user.UserConfModel.Sound
		resp.SecureLink = user.UserConfModel.SecureLink
		resp.SavePwd = user.UserConfModel.SavePwd
		resp.SearchUser = int8(user.UserConfModel.SearchUser)
		resp.Verification = user.UserConfModel.Verification
		if user.UserConfModel.RecallMessage != nil {
			resp.RecallMessage = user.UserConfModel.RecallMessage
		}
		if user.UserConfModel.VerificationQuestion != nil {
			resp.VerificationQuestion = &types.VerificationQuestion{
				Problem1: user.UserConfModel.VerificationQuestion.Problem1,
				Problem2: user.UserConfModel.VerificationQuestion.Problem2,
				Problem3: user.UserConfModel.VerificationQuestion.Problem3,
				Answer1:  user.UserConfModel.VerificationQuestion.Answer1,
				Answer2:  user.UserConfModel.VerificationQuestion.Answer2,
				Answer3:  user.UserConfModel.VerificationQuestion.Answer3,
			}
		}
	}

	return resp, nil
}
