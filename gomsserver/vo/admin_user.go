package vo


// 如果是 application/json, 不需要设置 tag

type AULogin struct {
	Account string `form:"account" validate:"min=2"`
	Password string `form:"password" validate:"min=5"`
}

type AUAdd struct {
	Account string `form:"account" validate:"min=2,max=10"`
	Type     uint8 `form:"type" validate:"gte=1,lte=3"` //类型 1:管理员|2:客服
	Remark string `form:"remark" validate:"min=2"`
}

type AUEdit struct {
	ID       uint32 `form:"id" validate:"gt=0"`
	Type     uint8 `form:"type" validate:"gte=1,lte=3"` //类型 1:管理员|2:客服
	Remark string `form:"remark" validate:"min=2"`
}

type AUPassword struct {
	ID       uint32
	Password string `form:"password" validate:"min=5"`
}

type AULock struct {
	ID       uint32 `form:"id" validate:"gt=0"`
	Password string `form:"password" validate:"min=5"`
}
