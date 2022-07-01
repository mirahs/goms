package page

import (
	"github.com/gin-gonic/gin"
	"gomsserver/util"
	"gorm.io/gorm"
)


type Param struct {
	Ctx *gin.Context
	Db  *gorm.DB

	Datas interface{}
	Fields []string

	Preloads [][]interface{}

	Distincts []interface{}

	Wheres [][]interface{}
	Ors    [][]interface{}

	Groups  []string
	Havings [][]interface{}

	Orders []string
}


func Page(param *Param) (int64, error) {
	db := param.Db

	if len(param.Fields) > 0 {
		db = db.Select(param.Fields)
	}

	for _, preload := range param.Preloads {
		if len(preload) >= 1 {
			db = db.Preload(preload[0].(string), preload[1:]...)
		}
	}

	if len(param.Distincts) > 0 {
		db = db.Distinct(param.Distincts...)
	}

	for _, where := range param.Wheres {
		if len(where) >= 1 {
			db = db.Where(where[0], where[1:]...)
		}
	}

	for _, or := range param.Ors {
		if len(or) >= 1 {
			db = db.Or(or[0], or[1:]...)
		}
	}

	for _, group := range param.Groups {
		db = db.Group(group)
	}

	for _, having := range param.Havings {
		if len(having) >= 1 {
			db = db.Having(having[0], having[1:]...)
		}
	}

	for _, order := range param.Orders {
		db = db.Order(order)
	}


	page, pageSize := util.PageInfo(param.Ctx)

	var count int64

	err := db.Model(param.Datas).Count(&count).Offset((page - 1) * pageSize).Limit(pageSize).Find(param.Datas).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return count, err
}
