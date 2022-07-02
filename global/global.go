package global

import (
	"bms-go/config"
	"time"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	"gorm.io/gorm"
)

var (
	SYS_CONFIG = new(config.Server) // 项目总配置
	SYS_VP     *viper.Viper         // Viper
	SYS_DB     *gorm.DB             // Mysql
	SYS_ZAP    *zap.Logger          // Zap日志
)

// BaseModel 数据表常用的4个字段
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
