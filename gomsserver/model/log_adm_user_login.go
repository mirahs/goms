package model

import (
	"gomsserver/constant"
	"gomsserver/util/time"
)


type LogAdmUserLogin struct {
	ID      uint32

	Type	uint8 `gorm:"not null; comment:登录类型 1:正常登录|2:token登录"`

	AdmUserID uint32 `gorm:"not null; comment:管理员ID; index"`
	AdmUser *AdmUser
	Time    uint32 `gorm:"not null; comment:时间; index"`
	Status  uint8  `gorm:"not null; comment:状态 0:失败(密码错误)|1:成功"`
	Ip      string `gorm:"type:varchar(15); not null; comment:IP"`
	Address string `gorm:"type:varchar(128); not null; comment:地址"`
}

// TableName 设置表名, gorm 默认是复数形式
func (user *LogAdmUserLogin) TableName() string {
	return "log_adm_user_login"
}


func LogAdmUserLoginAddSuccess(typeLogin uint8, admUserId uint32, ip, address string) {
	logAdmUserLoginAdd(typeLogin, admUserId, ip, address, 1)
}

func LogAdmUserLoginAddFailed(admUserId uint32, ip, address string) {
	logAdmUserLoginAdd(1, admUserId, ip, address, 0)
}

func LogAdmUserDelete(id uint32) {
	Db.Delete(&LogAdmUserLogin{ID: id})
}

func LogAdmUserFailedCheck(id uint32) bool {
	timeBegin, timeEnd := time.TodayBeginAdnEndTime()

	var count int64

	Db.Model(&LogAdmUserLogin{}).Where("`time` BETWEEN ? AND ?", timeBegin, timeEnd).Where("`adm_user_id`=?", id).Where("`status`=0").Count(&count)

	return count < int64(constant.LoginFailedMax)
}


func logAdmUserLoginAdd(typeLogin uint8, admUserId uint32, ip, address string, status int) {
	Db.Create(&LogAdmUserLogin{
		Type: typeLogin,

		AdmUserID: admUserId,
		Time:    uint32(time.UnixTime()),
		Status:  uint8(status),
		Ip:      ip,
		Address: address,
	})
}
