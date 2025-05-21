package main

import (
	"fim_server/core"
	"fim_server/fim_chat/chat_models"
	"fim_server/fim_group/group_models"
	"fim_server/fim_user/user_models"
	"flag"
	"fmt"
)

type Options struct {
	DB bool
}

func main() {
	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()
	if opt.DB {
		db := core.InitGorm("root:12345678@tcp(127.0.0.1:33069)/fim_server_db?charset=utf8mb4&parseTime=true")
		err := db.AutoMigrate(
			&user_models.UserModel{},         //用户表
			&user_models.FriendModel{},       //好友表
			&user_models.FriendVerifyModel{}, //好友验证表
			&user_models.UserConfModel{},     //用户配置表
			&chat_models.ChatModel{},         //对话表
			&group_models.GroupModel{},
			&group_models.GroupMemberModel{},
			&group_models.GroupMsgModel{},
			&group_models.GroupVerifyModel{},
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")
	} else {
		fmt.Println("请输入参数")
	}
	// fmt.Println("hello")
}
