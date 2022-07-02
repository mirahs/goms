package model

import "gorm.io/gorm"

type Host struct {
	Model

	Name 		string `gorm:"type:char(30); not null; comment:主机名(hostname); uniqueIndex"`

	SshPort 	uint16	`gorm:"not null; default:0; comment:ssh端口"`
	SshUsername string  `gorm:"type:char(30); not null; default:''; comment:ssh账号"`
	SshPassword string  `gorm:"type:char(50); not null; default:''; comment:ssh密码"`

	Remark   	string `gorm:"type:text; not null; comment:备注"`
}

// TableName 设置表名, gorm 默认是复数形式
func (admUser *Host) TableName() string {
	return "ms_host"
}


func HostGet(id uint32) (*Host, error) {
	var au Host

	err := Db.First(&au, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &au, nil
}

func HostGetByName(name string) (*Host, error) {
	var au Host

	err := Db.Where("name=?", name).First(&au).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &au, nil
}

func HostAdd(data map[string]interface{}) error {
	admUser := Host{
		Name: data["name"].(string),

		SshPort: data["ssh_port"].(uint16),
		SshUsername: data["ssh_username"].(string),
		SshPassword: data["ssh_password"].(string),

		Remark: data["remark"].(string),
	}

	return Db.Create(&admUser).Error
}

func HostEdit(id uint32, data map[string]interface{}) error {
	return Db.Model(&Host{Model: Model{ID: id}}).Updates(data).Error
}

func HostDelete(id uint32) error {
	return Db.Delete(&Host{Model: Model{ID: id}}).Error
}
