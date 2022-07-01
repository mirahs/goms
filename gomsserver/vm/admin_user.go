package vm

import (
	"gomsserver/model"
	"gomsserver/util"
)

type AdminUser struct {
	Id uint32
	Account string
	Password string
	Type     uint8 //类型 1:管理员
	Remark string
	IsLocked uint8

	CardNum	uint16
}


func (au *AdminUser) Get() (*model.AdmUser, error) {
	return model.AUGet(au.Id)
}

func (au *AdminUser) GetByAccount() (*model.AdmUser, error) {
	return model.AUGetByAccount(au.Account)
}

func (au *AdminUser) Add() error {
	data := map[string]interface{}{
		"account": au.Account,
		"type": au.Type,
		"remark": au.Remark,
		"password": util.BCryptHash(au.Password),
	}
	return model.AUAdd(data)
}

func (au *AdminUser) Edit() error {
	data := map[string]interface{}{
		"type": au.Type,
		"remark": au.Remark,
	}
	return model.AUEdit(au.Id, data)
}

func (au *AdminUser) Delete() error {
	return model.AUDelete(au.Id)
}

func (au *AdminUser) Lock() error {
	return model.AULock(au.Id, au.IsLocked)
}

func (au *AdminUser) Reset() error {
	return model.AUPassword(au.Id, au.Password)
}

func (au *AdminUser) PasswordFun() error {
	return model.AUPassword(au.Id, au.Password)
}


func (au *AdminUser) HasId() (bool, error) {
	return model.AUHasById(au.Id)
}

func (au *AdminUser) HasAccount() (bool, error) {
	return model.AUHasByAccount(au.Account)
}
