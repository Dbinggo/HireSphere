package initalize

import (
	"flag"
	"fmt"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func InitConfig() {
	// 初始化时间为东八区的时间
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "config", configs.DefaultConfigPath, "配置文件绝对路径或相对路径")
	flag.Parse()

	logrus.Printf("===> config path is: %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Printf("config file has changed")
		if err := viper.Unmarshal(&configs.Conf); err != nil {
			logrus.Fatalf("failed at unmarshal config file after change, err: %v", err)
		}
		logrus.Infof("%+v", configs.Conf)
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("failed at ReadInConfig, err: %v", err)
		panic(fmt.Sprintf("failed at init config: %v", err))
	}
	// 解析到变量中
	if err := viper.Unmarshal(&configs.Conf); err != nil {
		logrus.Errorf("failed at Unmarshal config file, err: %v", err)
		panic(fmt.Sprintf("failed at init config: %v", err))
	}
	logrus.Infof("%+v", configs.Conf)
}
