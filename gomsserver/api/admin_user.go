package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gomsserver/common"
	"gomsserver/constant"
	"gomsserver/dto"
	"gomsserver/model"
	"gomsserver/util"
	"gomsserver/util/page"
	"gomsserver/util/token"
	"gomsserver/vm"
	"gomsserver/vo"
	"log"
	"net/http"
)


func AdmUserLogin(ctx *gin.Context) {
	var voLogin vo.AULogin

	err := util.GinBindValid(ctx, &voLogin)
	if err != nil {
		log.Printf("admin_user.AdmUserLogin GinBindValid err:%v\n", err)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	ip := util.GinIP(ctx)
	address := util.IpAddress(ip)

	vmUser := vm.AdminUser{Account: voLogin.Account}

	user, err := vmUser.GetByAccount()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if user.ID == 0 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUAccountNotExist)
		return
	}
	if user.Locked == 1 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAULocked)
		return
	}
	if !model.LogAdmUserFailedCheck(user.ID) {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUPasswordFailedMax)
		return
	}

	if !util.BCryptCompare(user.Password, voLogin.Password) {
		model.LogAdmUserLoginAddFailed(user.ID, ip, address)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrAUPasswordNotSame)
		return
	}

	tokenMake := token.Make(user.ID, false)

	model.AULoginAfter(user, ip)
	model.LogAdmUserLoginAddSuccess(1, user.ID, ip, address)

	common.RespSuccess(ctx, gin.H{"token": tokenMake, "user": dto.ToAdmUser(user)})
}

func AdmUserLoginToken(ctx *gin.Context) {
	user:= ctx.MustGet(constant.KeyAdmUser).(*model.AdmUser)

	ip := util.GinIP(ctx)
	address := util.IpAddress(ip)

	model.AULoginAfter(user, ip)
	model.LogAdmUserLoginAddSuccess(2, user.ID, ip, address)

	common.RespSuccess(ctx, dto.ToAdmUser(user))
}

func AdmUserInfo(ctx *gin.Context) {
	user := ctx.MustGet(constant.KeyAdmUser).(*model.AdmUser)
	common.RespSuccess(ctx, dto.ToAdmUser(user))
}

func AdmUserLogout(ctx *gin.Context) {
	tokenX := ctx.GetHeader(constant.KeyHttpHeaderXToken)
	if tokenX != "" {
		token.Destroy(tokenX)
	}
	common.RespSuccess(ctx, nil)
}


// 获取所有用户
func AdmUserGetAll(ctx *gin.Context) {
	var users []*model.AdmUser

	count, err := page.Page(&page.Param{Ctx: ctx, Db: model.Db, Datas: &users})
	if err != nil {
		log.Printf("admin_user.AdmUserGetAll Page err:%v\n", err)
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, gin.H{"count": count, "datas": dto.ToAdmUsers(users)})
}

// 添加用户
func AdmUserAdd(ctx *gin.Context) {
	var voAdd vo.AUAdd

	err := util.GinBindValid(ctx, &voAdd)
	if err != nil {
		log.Printf("admin_user.AdmUserAdd GinBindValid err:%v\n", err)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	vmAU := vm.AdminUser{Account: voAdd.Account, Type: voAdd.Type, Remark: voAdd.Remark}
	exist, err := vmAU.HasAccount()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if exist {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUAccountExist)
		return
	}

	password := util.RandStr(20)
	vmAU.Password = password

	err = vmAU.Add()
	if err != nil {
		log.Printf("admin_user.AdmUserAdd vmAU.Add err:%v\n", err)
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, password)
}

// 更新用户
func AdmUserEdit(ctx *gin.Context) {
	id := com.StrTo(ctx.Params.ByName("id")).MustInt64()
	voEdit := vo.AUEdit{ID: uint32(id)}

	err := util.GinBindValid(ctx, &voEdit)
	if err != nil {
		log.Printf("admin_user.AdmUserEdit GinBindValid err:%v\n", err)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	vmUser := &vm.AdminUser{Id: voEdit.ID, Type: voEdit.Type, Remark: voEdit.Remark}

	exist, err := vmUser.HasId()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if !exist {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUIdNotExist)
		return
	}

	err = vmUser.Edit()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, nil)
}

// 删除用户
func AdmUserDelete(ctx *gin.Context) {
	uid := ctx.MustGet(constant.KeyAdmUserId).(uint32)

	id := uint32(com.StrTo(ctx.Params.ByName("id")).MustInt64())
	if id <= 0 {
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	if uid == id {
		common.RespFailed(ctx, http.StatusForbidden, common.ErrAUOperateSelf)
		return
	}

	vmUser := &vm.AdminUser{Id: id}

	exist, err := vmUser.HasId()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if !exist {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUIdNotExist)
		return
	}

	err = vmUser.Delete()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, nil)
}

// 锁住用户
func AdmUserLock(ctx *gin.Context) {
	id := uint32(com.StrTo(ctx.Params.ByName("id")).MustInt64())
	if id <= 0 {
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	uid := ctx.MustGet(constant.KeyAdmUserId).(uint32)
	if id == uid {
		common.RespFailed(ctx, http.StatusForbidden, common.ErrAUOperateSelf)
		return
	}

	vmUser := vm.AdminUser{Id: id}

	user, err := vmUser.Get()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if user.ID == 0 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUIdNotExist)
		return
	}

	isLocked := uint8(util.If(user.Locked == 0, 1, 0).(int))
	vmUser.IsLocked = isLocked

	err = vmUser.Lock()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	common.RespSuccess(ctx, gin.H{"is_locked": isLocked})
}

// 重置用户密码
func AdmUserReset(ctx *gin.Context) {
	id := uint32(com.StrTo(ctx.Params.ByName("id")).MustInt64())
	if id <= 0 {
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	uid := ctx.MustGet(constant.KeyAdmUserId).(uint32)
	if id == uid {
		common.RespFailed(ctx, http.StatusForbidden, common.ErrAUOperateSelf)
		return
	}

	vmUser := vm.AdminUser{Id: id}

	user, err := vmUser.Get()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if user.ID == 0 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrAUIdNotExist)
		return
	}

	password := util.RandStr(20)
	vmUser.Password = password

	err = vmUser.Reset()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	common.RespSuccess(ctx, password)
}

// 用户更改密码
func AdmUserPassword(ctx *gin.Context) {
	uid := ctx.MustGet(constant.KeyAdmUserId).(uint32)
	voPassword := vo.AUPassword{ID: uid}

	err := util.GinBindValid(ctx, &voPassword)
	if err != nil {
		log.Printf("admin_user.AdmUserPassword GinBindValid err:%v\n", err)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	vmUser := vm.AdminUser{Id: voPassword.ID, Password: voPassword.Password}
	err = vmUser.PasswordFun()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	common.RespSuccess(ctx, nil)
}
