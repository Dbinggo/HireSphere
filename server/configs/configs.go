package configs

type Config struct {
	App   ApplicationConfig `mapstructure:"app"`
	Log   LoggerConfig      `mapstructure:"log"`
	DB    DBConfig          `mapstructure:"database"`
	Redis RedisConfig       `mapstructure:"redis"`
	ES    ESConfig          `mapstructure:"es"`
}

type ApplicationConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}
type LoggerConfig struct {
	Format   string `mapstructure:"format"`
	Debug    bool   `mapstructure:"debug"`
	Director string `mapstructure:"director"`
	Caller   bool   `mapstructure:"caller"`
	Level    string `mapstructure:"level"`
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
	host string `mapstructure:"host"`
	port int    `mapstructure:"port"`
}
type ESConfig struct {
	Enable   bool   `mapstructure:"enable"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
}
