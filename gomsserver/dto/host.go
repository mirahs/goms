package dto

import (
	"gomsserver/model"
)


type Host struct {
	Id uint32 `json:"id"`
	CreatedAt uint32 `json:"created_at"`

	Name string `json:"name"`

	SshPort uint16 `json:"ssh_port"`
	SshUsername string `json:"ssh_username"`
	SshPassword string `json:"ssh_password"`

	Remark string `json:"remark"`
}


func ToHost(user *model.Host) *Host {
	return &Host{
		Id: user.ID,
		CreatedAt: user.CreatedAt,

		Name: user.Name,

		SshPort: user.SshPort,
		SshUsername: user.SshUsername,
		SshPassword: user.SshPassword,

		Remark: user.Remark,
	}
}

func ToHosts(users []*model.Host) (dtoUsers []*Host) {
	for _, user := range users {
		dtoUser := ToHost(user)
		dtoUsers = append(dtoUsers, dtoUser)
	}

	return
}
