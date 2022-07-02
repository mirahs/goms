package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomsserver/conf"
	"gomsserver/constant"
	"gomsserver/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


// 全文索引设置 index:,class:FULLTEXT;
type Model struct {
	// 默认使用 ID 或者 ID 作为主键, 不用设置字段标签 primarykey
	// 等同于 ID uint32 和 ID uint32
	ID        uint32 `gorm:"primarykey"`
	// CreatedAt 新增数据会自动赋值
	CreatedAt uint32 `gorm:"autoCreateTime; index; comment:创建时间"`
	// UpdatedAt 新增和更新数据会自动赋值
	UpdatedAt uint32 `gorm:"autoUpdateTime; index; comment:更新时间"`
	// DeletedAt 字段存在为软删除并会自动赋值
	// 软删除后的数据要查询要先调用 Unscoped()
	DeletedAt gorm.DeletedAt `gorm:"comment:删除时间"`
}

type NoUModel struct {
	ID        uint32 `gorm:"primarykey"`
	CreatedAt uint32 `gorm:"autoCreateTime; index; comment:创建时间"`
	// DeletedAt 字段存在为软删除并会自动赋值
	// 软删除后的数据要查询要先调用 Unscoped()
	DeletedAt gorm.DeletedAt `gorm:"comment:删除时间"`
}

type NoUDModel struct {
	ID        uint32 `gorm:"primarykey"`
	CreatedAt uint32 `gorm:"autoCreateTime; index; comment:创建时间"`
}


var Db *gorm.DB


// gorm 初始化
func Init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.App.MysqlUser, conf.App.MysqlPassword, conf.App.MysqlHost,
		conf.App.MysqlPort, conf.App.MysqlDatabase, conf.App.MysqlCharset,
	)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
		Logger: ginLogger(),
	})
	if err != nil {
		panic("model.Init Open err:" + err.Error())
	}

	sqlDB, err := Db.DB()
	if err != nil {
		panic("model.Init DB err:" + err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 判断 adm_user 表是否存在(要在 AutoMigrate 操作之前)
	hasAdmUser := Db.Migrator().HasTable(&AdmUser{})

	err = Db.AutoMigrate(
		&AdmUser{},
		&LogAdmUserLogin{},

		&Host{},
	)
	if err != nil {
		panic("model.Init AutoMigrate err:" + err.Error())
	}

	// todo: 如果启动时 adm_user 表不存在就创建初始化管理员账号(登录后记得改密码或者删除这个账号)
	if !hasAdmUser {
		Db.Create(&AdmUser{
			Account:  conf.App.InitAccount,
			Password: util.BCryptHash(conf.App.InitPassword),
			Type:     constant.AdminUserTypeAdmin,
			Remark:   conf.App.InitAccount,
		})
	}

	if conf.App.GinMode == gin.DebugMode {
		Db = Db.Debug()
	}
}

func Close() {
	sqlDB, _ := Db.DB()
	_ = sqlDB.Close()
}


func ginLogger() logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             1 * time.Second,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})
}
