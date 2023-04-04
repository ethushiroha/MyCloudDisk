package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitViper() *viper.Viper {
	rootPath, err := os.Executable()
	if err != nil {
		panic("初始化失败：可执行程序路径获取失败")
	}
	rootPath = filepath.Dir(rootPath)
	path := filepath.Join(rootPath, "config/BurpConfig_dev.yaml")
	fmt.Println(path)

	v := viper.New()
	v.SetConfigFile(path)
	fmt.Println(path)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		panic("初始化失败：读取配置文件失败")
	}
	v.Set("root_path", rootPath)
	return v
}
