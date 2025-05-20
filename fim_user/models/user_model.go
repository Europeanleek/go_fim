package models

import "fim_server/common/models"

// 用户基本信息表
type UserModel struct {
	models.Model
	// 密码
	Pwd string `json:"pwd"`
	// 昵称
	NickName string `json:"nickname"`
	// 简介
	Abstract string `json:"abstract"`
	// 头像
	Avatar string `json:"avatar"`
	IP     string `json:"ip"`
	Addr   string `json:"addr"`
}
