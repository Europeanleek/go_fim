package user_models

import "fim_server/common/models"

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint      `json:"sendUserID"`                         //发起验证方
	SendUserModel        UserModel `gorm:"foreignKey:SendUserID" json:"-"`     //发起验证方
	RevUserID            uint      `json:"revUserID"`                          //接受验证方
	RevUserModel         UserModel `gorm:"foreignKey:RevUserID" json:"-"`      //接受验证方
	Status               int8      `json:"status"`                             //状态0 未操作 1 同意2拒绝3忽略
	AdditonalMessages    string    `gorm:"size:128" json:"additionalMessages"` //附加消息
	VerificationQuestion string    `json:"verificationQuestion"`               //验证问题
}
