package vm

import (
	"gomsserver/model"
)


type Host struct {
	Id uint32

	Name string

	SshPort uint16
	SshUsername string
	SshPassword string

	Remark string
}


func (au *Host) Get() (*model.Host, error) {
	return model.HostGet(au.Id)
}

func (au *Host) GetByName() (*model.Host, error) {
	return model.HostGetByName(au.Name)
}

func (au *Host) Add() error {
	data := map[string]interface{}{
		"name": au.Name,

		"ssh_port": au.SshPort,
		"ssh_username": au.SshUsername,
		"ssh_password": au.SshPassword,

		"remark": au.Remark,
	}
	return model.HostAdd(data)
}

func (au *Host) Edit() error {
	data := map[string]interface{}{
		"name": au.Name,

		"ssh_port": au.SshPort,
		"ssh_username": au.SshUsername,
		"ssh_password": au.SshPassword,

		"remark": au.Remark,
	}
	return model.HostEdit(au.Id, data)
}

func (au *Host) Delete() error {
	return model.HostDelete(au.Id)
}
