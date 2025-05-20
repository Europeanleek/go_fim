package chat_models

import (
	"fim_server/common/models"
	"fim_server/common/models/ctype"
)

type ChatModel struct {
	models.Model
	SendUserID uint             `json:"sendUserId"`
	RevUserID  uint             `json:"revUserID"`
	MsgType    int              `json:"msgType"`                   //消息类型 1文本类型 2图片消息 3视频消息 4文件消息 5语音消息 6语音通话 7视频通话 8撤回消息 9回复消息 10引用消息
	Msgpreview string           `gorm:"size:64" json:"msgpreview"` //消息预览
	Msg        ctype.Msg        `json:"msg"`                       //消息内容
	SystemMsg  *ctype.SystemMsg `json:"systemMsg"`                 //系统提示
}
