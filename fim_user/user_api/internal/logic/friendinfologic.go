package logic

import (
	"context"
	"encoding/json"
	"errors"

	"fim_server/fim_user/user_api/internal/svc"
	"fim_server/fim_user/user_api/internal/types"
	"fim_server/fim_user/user_models"
	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	// todo: add your logic here and delete this line
	//
	var friend user_models.FriendModel
	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("他不是你的好友")
	}
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.FriendID),
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	var friendUser user_models.UserModel
	json.Unmarshal(res.Data, &friendUser)
	response := types.FriendInfoResponse{
		UserID:   friendUser.ID,
		Nickname: friendUser.NickName,
		Abstract: friendUser.Abstract,
		Avatar:   friendUser.Avatar,
		Notice:   friend.GetUserNotice(req.UserID),
	}
	return &response, nil
}
