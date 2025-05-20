package models

import (
	"fim_server/common/models"
	"fim_server/common/models/ctype"
)

type GroupModel struct {
	models.Model
	Title                string                      `json:"title"`                //群名
	Abstract             string                      `json:"abstract"`             //简介
	Avatar               string                      `json:"avatar"`               //群头像
	Creator              uint                        `json:"creator"`              //群主
	IsSearch             bool                        `json:"isSearch"`             //是否可以被搜索
	Verification         int8                        `json:"verification"`         //群验证 0 不允许任何人添加 1 允许任何人添加 2 需要验证消息 3需要回答对应问题才可以添加好友 4需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"` //验证问题，为3和4时有效
	IsInvite             bool                        `json:"isInvite"`             //是否可以邀请好友加入
	IsTmporarySession    bool                        `json:"isTmporarySession"`    //是否为临时会话
	IsProhibition        bool                        `json:"isProhibition"`        //是否开启全员禁言
	Size                 int                         `json:"size"`                 //群规模 20 100 200 1000
}
