package group_models

import "fim_server/common/models"

type GroupMemberModel struct {
	models.Model
	GroupID        uint       `json:"groupID"` //群id
	GroupModel     GroupModel `gorm:"foreignKey:GroupID" json:"-"`
	UserID         uint       `json:"userID"`                        //用户id
	MemberNickName string     `gorm:"size:32" json:"memberNickName"` //群昵称
	Role           int        `json:"role"`                          //1 群主 2管理员 3群员
	Prohibition    *int       `json:"prohibition"`                   //禁言时间
}
