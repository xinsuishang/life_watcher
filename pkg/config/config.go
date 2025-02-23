package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type LogConfig struct {
	Path string `mapstructure:"path"`
}

type PongTimeConfig struct {
	Version string `mapstructure:"version"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Security SecurityConfig `mapstructure:"security"`
	Log      LogConfig      `mapstructure:"log"`
	PongTime PongTimeConfig `mapstructure:"pong_time"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	TLS      string `mapstructure:"tls"`
}

type SecurityConfig struct {
	JWTSecret string `mapstructure:"jwt_secret"`
	JWTExpire int    `mapstructure:"jwt_expire"`
	CryptoKey string `mapstructure:"crypto_key"`
}

var (
	once sync.Once
	conf *Config
)

func GetConfig() *Config {
	once.Do(initConfig)
	return conf
}

func initConfig() {
	// AIGC START
	// 设置环境变量前缀
	viper.SetEnvPrefix("APP_CONFIG")

	// 将.替换为_用于环境变量映射
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 启用环境变量自动映射
	viper.AutomaticEnv()

	// 配置文件作为后备
	if viper.GetBool("local") {
		viper.SetConfigName("config")
	} else {
		// 非本地环境，使用config_template.yml 做解析占位
		viper.SetConfigName("config_template")
	}
	viper.AddConfigPath("conf")
	viper.SetConfigType("yml")

	// 尝试读取配置文件，但不强制要求存在
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("读取配置文件失败: %w", err))
		} else {
			fmt.Println("配置文件不存在，使用环境变量")
		}
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("解析配置文件失败: %w", err))
	}
}

func (dc *DatabaseConfig) DSN() string {
	return MySqlDSNFormatUtil(dc.User, dc.Password, dc.Host, dc.Port, dc.DBName, dc.TLS, "")
}
