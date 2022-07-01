package constant


const (
	AdminUserTypeAdmin uint8 = 1 //管理员
	AdminUserTypeCS uint8 = 2 //客服

	LoginTypeNormal uint8 = 1 //正常登录
	LoginTypeToken uint8 = 2 //token登录

	LoginFailedMax uint8 = 3 //登录失败最大次数
)
