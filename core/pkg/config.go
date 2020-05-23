package pkg

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	redis "gopkg.in/redis.v5"
)

var (
	UserName = "username"
	Avatar   = "avatar"
	UID      = "uid"
	PID      = "pid"
	UserType = "userType"
)
var (
	Cfg        *Config
	DBWrite    *gorm.DB
	DBRead     *gorm.DB
	IMDBRead   *gorm.DB
	RedisCli   *redis.Client
	HttpClient *http.Client
)

type Config struct {
	Logger     *LoggerConfig
	DBWrite    *DBWriteConfig
	DBRead     *DBReadConfig
	IMDBRead   *IMDBReadConfig
	Redis      *RedisConfig
	HttpClient *HttpClientConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
	rand.Seed(time.Now().UnixNano())
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("nitConfig error:%s", err)
		panic("initConfig error")
	}
}

func loadConfig() *Config {
	cfg := &Config{
		Logger:     LoadLoggerConfig(viper.Sub("logger")),
		DBWrite:    LoadDBWriteConfig(viper.Sub("db_write")),
		DBRead:     LoadDBReadConfig(viper.Sub("db_read")),
		IMDBRead:   LoadIMDBReadConfig(viper.Sub("im_db_read")),
		Redis:      LoadRedisConfig(viper.Sub("redis")),
		HttpClient: LoadHttpClientConfig(),
	}
	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
	DBWrite = cfg.DBWrite.InitDBWrite()
	DBRead = cfg.DBRead.InitDBRead()
	IMDBRead = cfg.IMDBRead.InitIMDBRead()
	RedisCli = cfg.Redis.InitRedis()
	HttpClient = cfg.HttpClient.InitHttpClient()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s", e.Name)
	})
}
