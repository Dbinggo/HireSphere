package configs

const DefaultConfigPath = "./config.yaml"

var Conf = new(Config)

type Config struct {
	App ApplicationConfig `mapstructure:"app"`
	Log LoggerConfig      `mapstructure:"log"`
	DB  DBConfig          `mapstructure:"database"`
	Rdb RedisConfig       `mapstructure:"redis"`
}

type ApplicationConfig struct {
	Port int `mapstructure:"port"`
}
type LoggerConfig struct {
	Level    int8   `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Director string `mapstructure:"director"`
}

type DBConfig struct {
	Driver      string `mapstructure:"driver"`
	AutoMigrate bool   `mapstructure:"migrate"`
	Dsn         string `mapstructure:"dsn"`
}
type RedisConfig struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type KafkaConfig struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}
