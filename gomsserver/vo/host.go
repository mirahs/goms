package vo


// 如果是 application/json, 不需要设置 tag

type HostAdd struct {
	Name string `json:"name" validate:"min=2,max=30"`

	SshPort uint16 `json:"ssh_port" validate:"omitempty,gte=1,lte=65535"`
	SshUsername string `json:"ssh_username" validate:"omitempty,max=30"`
	SshPassword string `json:"ssh_password" validate:"omitempty,max=50"`

	Remark string `json:"remark" validate:"min=2"`
}

type HostEdit struct {
	ID       uint32 `json:"id" validate:"gt=0"`

	Name string `json:"name" validate:"min=2,max=30"`

	SshPort uint16 `json:"ssh_port" validate:"omitempty,gte=1,lte=65535"`
	SshUsername string `json:"ssh_username" validate:"omitempty,max=30"`
	SshPassword string `json:"ssh_password" validate:"omitempty,max=50"`

	Remark string `json:"remark" validate:"min=2"`
}
