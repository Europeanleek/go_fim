package model

import (
	"fim_server/common/models"
	"time"
)

type ChatModel struct {
	models.Model
	SendUserID uint      `json:"sendUserId"`
	RevUserID  uint      `json:"revUserID"`
	MsgType    int       `json:"msgType"`    //消息类型 1文本类型 2图片消息 3视频消息 4文件消息 5语音消息 6语音通话 7视频通话 8撤回消息 9回复消息 10引用消息
	Msgpreview string    `json:"msgpreview"` //消息预览
	Msg        Msg       `json:"msg"`        //消息内容
	SystemMsg  SystemMsg `json:"systemMsg"`  //系统提示
}

type SystemMsg struct {
	Type int8 `json:"type"` //违规类型 1 涉黄 2 涉恐 3 涉政 不正当言论
}
type Msg struct {
	Type         int8          `json:"type"` //回复消息类型 1文本类型 2图片消息 3视频消息 4文件消息 5语音消息 6语音通话 7视频通话 8撤回消息 9回复消息 10引用消息
	Content      *string       `json:"content"`
	ImageMsg     *ImageMsg     `json:"imageMsg"`
	VideoMsg     *VideoMsg     `json:"videoMsg"`
	FileMsg      *FileMsg      `json:"fileMsg"`
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"`
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"`
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`
	ReplyMsg     *ReplyMsg     `json:"replyMsg"`
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`
}
type ImageMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type VideoMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  string `json:"time"` //时长
}
type FileMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` //文件大小
	Type  string `json:"type"` //文件类型
}
type VoiceMsg struct {
	Src  string `json:"src"`
	Time int    `json:"time"` //时长（秒）
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"` //开始时间
	EndTime   time.Time `json:"endTime"`   //结束时间
	EndReason int8      `json:"endReason"` //结束原因 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}
type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"` //开始时间
	EndTime   time.Time `json:"endTime"`   //结束时间
	EndReason int8      `json:"endReason"` //结束原因 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}
type WithdrawMsg struct {
	Content   string `json:"content"` //撤回的提示词
	OriginMsg *Msg   `json:"-"`       //原消息
}
type ReplyMsg struct {
	MsgID   uint   `json:"msgID"`   //消息ID
	Content string `json:"content"` //回复的文本消息，目前只能回复文本消息
	Msg     *Msg   `json:"msg"`
}
type QuoteMsg struct {
	MsgID   uint   `json:"msgID"`   //消息ID
	Content string `json:"content"` //回复的文本消息，目前只能回复文本消息
	Msg     *Msg   `json:"msg"`
}
