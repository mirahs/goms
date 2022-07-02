package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gomsserver/common"
	"gomsserver/dto"
	"gomsserver/model"
	"gomsserver/util"
	"gomsserver/util/page"
	"gomsserver/vm"
	"gomsserver/vo"
	"log"
	"net/http"
)


// HostGetAll 获取所有主机
func HostGetAll(ctx *gin.Context) {
	var hosts []*model.Host

	count, err := page.Page(&page.Param{Ctx: ctx, Db: model.Db, Datas: &hosts})
	if err != nil {
		log.Printf("host.HostGetAll Page err:%v\n", err)
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, gin.H{"count": count, "datas": dto.ToHosts(hosts)})
}

// HostAdd 添加主机
func HostAdd(ctx *gin.Context) {
	var voAdd vo.HostAdd

	err := util.GinBindValid(ctx, &voAdd)
	if err != nil {
		log.Printf("host.HostAdd GinBindValid err:%v\n", err)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	vmAU := vm.Host{Name: voAdd.Name, SshPort: voAdd.SshPort, SshUsername: voAdd.SshUsername, SshPassword: voAdd.SshPassword, Remark: voAdd.Remark}

	host, err := vmAU.GetByName()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if host.ID != 0 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrHostExist)
		return
	}

	err = vmAU.Add()
	if err != nil {
		log.Printf("host.HostAdd vmAU.Add err:%v\n", err)
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, nil)
}

// HostEdit 更新主机
func HostEdit(ctx *gin.Context) {
	id := com.StrTo(ctx.Params.ByName("id")).MustInt64()
	voEdit := vo.HostEdit{ID: uint32(id)}

	err := util.GinBindValid(ctx, &voEdit)
	if err != nil {
		log.Printf("host.HostEdit GinBindValid err:%v\n", err)
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}
	fmt.Printf("voEdit:%v\n", voEdit)

	vmUser := vm.Host{Id: voEdit.ID, Name: voEdit.Name, SshPort: voEdit.SshPort, SshUsername: voEdit.SshUsername, SshPassword: voEdit.SshPassword, Remark: voEdit.Remark}

	host, err := vmUser.Get()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if host.ID == 0 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrHostIdNotExist)
		return
	}

	err = vmUser.Edit()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, nil)
}

// HostDelete 删除主机
func HostDelete(ctx *gin.Context) {
	id := uint32(com.StrTo(ctx.Params.ByName("id")).MustInt64())
	if id <= 0 {
		common.RespFailed(ctx, http.StatusBadRequest, common.ErrParamFailed)
		return
	}

	vmUser := &vm.Host{Id: id}

	host, err := vmUser.Get()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}
	if host.ID == 0 {
		common.RespFailed(ctx, http.StatusNotFound, common.ErrHostIdNotExist)
		return
	}

	err = vmUser.Delete()
	if err != nil {
		common.RespFailed(ctx, http.StatusInternalServerError, common.ErrDb)
		return
	}

	common.RespSuccess(ctx, nil)
}
