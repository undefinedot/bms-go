package core

import (
	"bms-go/global"
	"flag"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	// 命令行参数 todo: 可优化
	var config string
	if len(path) > 1 {
		flag.StringVar(&config, "c", "", "choose a config file")
		flag.Parse()
		// 必须提供有效的选项
		if config != "" {
			fmt.Printf("使用配置文件为：%s\n", config)
		}
	} else {
		config = "./config.yaml" // 默认配置
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("加载配置失败：%s\n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置已更新:", in.Name)
		if err = v.Unmarshal(global.SYS_CONFIG); err != nil {
			fmt.Println("Unmarshal config failed:", err)
		}
	})

	if err = v.Unmarshal(global.SYS_CONFIG); err != nil {
		fmt.Println("Unmarshal config failed:", err)
	}

	log.Println("----- Viper init succeed -----")
	return v
}
