package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// InitConfig 初始化配置文件
func InitConfig(configFile string) {
	viper.SetConfigType("json")
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	viper.WatchConfig()
}

func GetToolSchedule(name string) string {
	key := "tools." + name + ".schedule"
	return viper.GetString(key)
}

func IsToolActive(name string) bool {
	key := "tools." + name + ".active"
	return viper.GetBool(key)
}

func GetLogFile() string {
	return viper.GetString("log_file")
}

func GetBrowserPath() string {
	return viper.GetString("browser_path")
}

func GetCRMUrl() string {
	return viper.GetString("crm_url")
}

func GetOAUrl() string {
	return viper.GetString("oa_url")
}
