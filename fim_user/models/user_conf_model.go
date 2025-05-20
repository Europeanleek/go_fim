package models

import (
	"fim_server/common/models"
	"fim_server/common/models/ctype"
)

// 用户设置表
type UserConfModel struct {
	models.Model
	UserId               uint                        `json:"user_id"`
	RecallMessage        *string                     `json:"recall_message"`       // 撤回消息设置
	FriendOnline         bool                        `json:"friendOnline"`         // 好友上线通知
	Sound                bool                        `json:"sound"`                // 好友上线是否发出提醒声
	SecureLink           bool                        `json:"secureLink"`           // 安全连接
	SavePwd              bool                        `json:"savePwd"`              //是否记住密码
	SearchUser           int                         `json:"searchUser"`           //别人查找到我的方式，0不允许别人找到我，1 通过用户号找到我 2可以通过昵称搜索到我
	Verification         int8                        `json:"Verification"`         //好友验证 0 不允许任何人添加 1 允许任何人添加 2 需要验证消息 3需要回答对应问题才可以添加好友 4需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"` //验证问题 只有FriendVerification为3或者4时才需要
	Online               bool                        `json:"online"`               //是否在线
}
