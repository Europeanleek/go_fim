package group_models

import (
	"fim_server/common/models"
	"fim_server/common/models/ctype"
)

// 群消息
type GroupMsgModel struct {
	models.Model
	GroupID    uint             `json:"groupID"` //群id
	GroupModel GroupModel       `gorm:"foreignKey:GroupID" json:"-"`
	SendUserID uint             `json:"sendUserId"`                //发送者id
	MsgType    int              `json:"msgType"`                   //消息类型 1文本类型 2图片消息 3视频消息 4文件消息 5语音消息 6语音通话 7视频通话 8撤回消息 9回复消息 10引用消息 11@用户消息
	Msgpreview string           `gorm:"size:64" json:"msgpreview"` //消息预览
	Msg        ctype.Msg        `json:"msg"`                       //消息内容
	SystemMsg  *ctype.SystemMsg `json:"systemMsg"`                 //系统提示
}
