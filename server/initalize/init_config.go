package initalize

import (
	"flag"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

// TODO 可以添加环境变量 环境变量没有加在上面 单纯觉得有点用不到
func InitConfig() {
	// 初始化时间为东八区的时间
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "c", global.Path+global.CONFIG_FILE_PATH_DEFAULT, "配置文件绝对路径或相对路径")
	flag.Parse()
	zlog.Infof("the config path is %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		zlog.Warnf("config file changed ")
		if err := viper.Unmarshal(&configs.Conf); err != nil {
			zlog.Errorf("failed to unmarshal with err :  %v", err)
		}
		zlog.Debugf("new config is %+v", configs.Conf)
		global.Config = configs.Conf
		Eve()
		InitLog(configs.Conf)
		InitDataBase(*configs.Conf)
		InitRedis(*configs.Conf)
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		zlog.Panicf("failed to get config.yaml err: %v", err)

	}
	// 解析到变量中
	if err := viper.Unmarshal(&configs.Conf); err != nil {
		zlog.Panicf("fail to unmarshal config.yaml with err: %v", err)
	}
	zlog.Debugf("config.yaml is : %+v", configs.Conf)
	global.Config = configs.Conf
}
