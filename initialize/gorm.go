package initialize

import (
	"bms-go/global"
	"bms-go/model"
	"fmt"
	"os"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	cfg := global.SYS_CONFIG.Mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.Dns(),
		DefaultStringSize:         191, // string字段的长度
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用物理外键
	})
	if err != nil {
		fmt.Println("初始化数据库失败：", err)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	zap.L().Info("----- gorm init mysql succeed -----")
	return db
}

// RegisterTables 初始化数据表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Authority{},
		model.BaseMenu{},
	)
	if err != nil {
		zap.L().Error("initialize.RegisterTables 注册数据表失败", zap.Error(err))
		os.Exit(1)
	}
	zap.L().Info("注册数据表成功！")
}
