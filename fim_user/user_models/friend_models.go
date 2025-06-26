package user_models

import (
	"fim_server/common/models"
	"fmt"

	"gorm.io/gorm"
)

//好友表

type FriendModel struct {
	models.Model
	SendUserID     uint      `json:"sendUserID"`                     //发起验证的用户方
	SendUserModel  UserModel `gorm:"foreignKey:SendUserID" json:"-"` //发起验证方
	RevUserID      uint      `json:"revUserID"`                      //接收服务的验证方
	RevUserModel   UserModel `gorm:"foreignKey:RevUserID" json:"-"`  //接受验证方
	SendUserNotice string    `gorm:"size:128" json:"senUserNotice"`  //发送方备注
	RevUserNotice  string    `gorm:"size:128" json:"revUserNotice"`  //接收方备注
}

func (f FriendModel) IsFriend(db *gorm.DB, A, B uint) bool {
	err := db.Take(&f, "(send_user_id=? and rev_user_id=?) or (send_user_id=? and rev_user_id=?)", A, B, B, A).Error
	if err == nil {
		return true
	}
	return false
}

func (f *FriendModel) GetUserNotice(userID uint) string {
	fmt.Println(userID)
	fmt.Println(f.SendUserID)
	fmt.Println(f.RevUserID)
	if userID == f.SendUserID {
		return f.SendUserNotice
	}
	if userID == f.RevUserID {
		return f.RevUserNotice
	}
	return ""
}
