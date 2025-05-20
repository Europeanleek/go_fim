package user_models

import "fim_server/common/models"

//好友表

type FriendModel struct {
	models.Model
	SendUserID    uint      `json:"sendUserID"`                     //发起验证的用户方
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"` //发起验证方
	RevUserID     uint      `json:"revUserID"`                      //接收服务的验证方
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID" json:"-"`  //接受验证方
	Notice        string    `gorm:"size:128" json:"notice"`         //加好友时的验证消息
}
