package vo


// 如果是 application/json, 不需要设置 tag

type LogLoginSearch struct {
	Account string `form:"account" validate:"omitempty"`
}
