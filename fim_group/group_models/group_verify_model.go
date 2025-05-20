package group_models

import (
	"fim_server/common/models"
	"fim_server/common/models/ctype"
)

type GroupVerifyModel struct {
	models.Model
	GroupID              uint                        `json:"groupID"` //群ID
	GroupModel           GroupModel                  `gorm:"foreignKey:GroupID" json:"-"`
	UserID               uint                        `json:"userID"`                            //用户ID
	Status               int8                        `json:"status"`                            //状态0 未操作 1 同意 2 拒绝 3忽略
	AdditionalMessages   string                      `gorm:"size:32" json:"additionalMessages"` //附加信息
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"`              //验证问题 为3和4的时候需要
	Type                 int8                        `json:"type"`                              //类型 1 加群 2退群
}
