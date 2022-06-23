package initialize

import (
	"bms-go/global"
	"fmt"

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
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("初始化数据库失败：", err)
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	return db
}
