package dto

import (
	"gomsserver/model"
)


type AdmUser struct {
	ID       uint32 `json:"id"`
	Account  string `json:"account"`
	Type uint8 `json:"type"`

	Locked bool `json:"locked"`
	Remark string `json:"remark"`

	CreatedAt uint32 `json:"created_at"`

	LoginAt uint32 `json:"login_at"`
	LoginIp string `json:"login_ip"`
	LoginTimes uint32 `json:"login_times"`
}


func ToAdmUser(user *model.AdmUser) *AdmUser {
	admUser := &AdmUser{
		ID:       user.ID,
		Account:  user.Account,
		Type: 		user.Type,

		Locked:     user.Locked == 1,
		Remark:       user.Remark,

		CreatedAt: user.CreatedAt,

		LoginAt: user.LoginAt,
		LoginIp:       user.LoginIp,
		LoginTimes:    user.LoginTimes,
	}

	return admUser
}

func ToAdmUsers(users []*model.AdmUser) (dtoUsers []*AdmUser) {
	for _, user := range users {
		dtoUser := ToAdmUser(user)
		dtoUsers = append(dtoUsers, dtoUser)
	}

	return
}
