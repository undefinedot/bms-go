package global

import (
	"bms-go/config"

	"github.com/spf13/viper"

	"gorm.io/gorm"
)

var (
	SYS_CONFIG = new(config.Server) // 项目总配置
	SYS_VP     *viper.Viper         // Viper
	SYS_DB     *gorm.DB             // Mysql
)
