package vo


// 如果是 application/json, 不需要设置 tag

type HostAdd struct {
	Name string `form:"name" validate:"min=2,max=30"`

	SshPort uint16 `form:"ssh_port" validate:"omitempty,gte=1,lte=65535"`
	SshUsername string `form:"ssh_username" validate:"omitempty,max=30"`
	SshPassword string `form:"ssh_password" validate:"omitempty,max=50"`

	Remark string `form:"remark" validate:"min=2"`
}

type HostEdit struct {
	ID       uint32 `form:"id" validate:"gt=0"`

	Name string `form:"name" validate:"min=2,max=30"`

	SshPort uint16 `form:"ssh_port" validate:"omitempty,gte=1,lte=65535"`
	SshUsername string `form:"ssh_username" validate:"omitempty,max=30"`
	SshPassword string `form:"ssh_password" validate:"omitempty,max=50"`

	Remark string `form:"remark" validate:"min=2"`
}
