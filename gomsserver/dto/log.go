package dto

import (
	"gomsserver/model"
)


type LogAdmUserLogin struct {
	Id uint32 `json:"id"`

	Type uint8 `json:"type"`

	Account string `json:"account"`
	Time uint32 `json:"time"`
	Status uint8 `json:"status"`
	Ip string `json:"ip"`
	Address string `json:"address"`
}


func ToLogLoginAdmUser(user *model.LogAdmUserLogin) *LogAdmUserLogin {
	admUser := &LogAdmUserLogin{
		Id: user.ID,

		Type: user.Type,

		Time: user.Time,
		Status: user.Status,
		Ip: user.Ip,
		Address: user.Address,
	}

	if user.AdmUser != nil {
		admUser.Account = user.AdmUser.Account
	}

	return admUser
}

func ToLogLoginAdmUsers(users []*model.LogAdmUserLogin) (dtoUsers []*LogAdmUserLogin) {
	for _, user := range users {
		dtoUser := ToLogLoginAdmUser(user)
		dtoUsers = append(dtoUsers, dtoUser)
	}

	return
}
