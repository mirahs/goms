package model

import (
	"gomsserver/util"
	"gomsserver/util/time"
	"gorm.io/gorm"
)


type AdmUser struct {
	Model

	Account  	string `gorm:"type:varchar(32); not null; comment:帐号; uniqueIndex"`
	Password 	string `gorm:"type:char(72); not null; comment:密码"`
	Type     	uint8  `gorm:"not null; comment:类型 1:管理员|2:客服"`

	Locked 		uint8  `gorm:"not null; default:0; comment:是否被锁定 0:否|1:是"`
	Remark   	string `gorm:"type:text; not null; comment:备注"`

	LoginAt		uint32 `gorm:"not null; default:0; comment:登录时间; index"`
	LoginIp		string `gorm:"type:varchar(15); not null; default:''; comment:登录IP"`
	LoginTimes	uint32 `gorm:"not null; default:0; comment:登录次数"`
}

// TableName 设置表名, gorm 默认是复数形式
func (admUser *AdmUser) TableName() string {
	return "adm_user"
}


// 登录成功后调用
func AULoginAfter(user *AdmUser, ip string) {
	data := map[string]interface{}{
		"login_at": uint32(time.UnixTime()),
		"login_ip": ip,
		"login_times": user.LoginTimes + 1,
	}

	Db.Model(user).Updates(data)
}


func AUGet(id uint32) (*AdmUser, error) {
	var au AdmUser

	err := Db.First(&au, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &au, nil
}

func AUGetByAccount(account string) (*AdmUser, error) {
	var au AdmUser

	err := Db.Where("account=?", account).First(&au).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &au, nil
}

func AUAdd(data map[string]interface{}) error {
	admUser := AdmUser{
		Account: data["account"].(string),
		Type: data["type"].(uint8),
		Remark: data["remark"].(string),
		Password: data["password"].(string),
	}

	return Db.Create(&admUser).Error
}

func AUEdit(id uint32, data map[string]interface{}) error {
	return Db.Model(&AdmUser{Model: Model{ID: id}}).Updates(data).Error
}

func AUDelete(id uint32) error {
	return Db.Delete(&AdmUser{Model: Model{ID: id}}).Error
}

func AULock(id uint32, locked uint8) error {
	data := map[string]interface{}{"locked": locked}

	return Db.Model(&AdmUser{Model: Model{ID: id}}).Updates(data).Error
}

func AUPassword(id uint32, password string) error {
	data := map[string]interface{}{"password": util.BCryptHash(password)}

	return Db.Model(&AdmUser{Model: Model{ID: id}}).Updates(data).Error
}

func AUHasById(id uint32) (bool, error) {
	var au AdmUser

	err := Db.First(&au, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return au.ID > 0, nil
}

func AUHasByAccount(account string) (bool, error) {
	var au AdmUser

	err := Db.Unscoped().Where("account=?", account).First(&au).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return au.ID > 0, nil
}
