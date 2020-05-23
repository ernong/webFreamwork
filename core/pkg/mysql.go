package pkg

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DBWriteConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}

type DBReadConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}
type IMDBReadConfig struct {
	Schema   string
	Host     string
	Port     string
	Username string
	Password string
}

func LoadDBWriteConfig(viper *viper.Viper) *DBWriteConfig {
	cfg := &DBWriteConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	fmt.Printf("LoadDBWriteConfig:%#v\n", cfg)
	return cfg
}

func LoadDBReadConfig(viper *viper.Viper) *DBReadConfig {
	cfg := &DBReadConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	fmt.Printf("LoadDBReadConfig:%#v\n", cfg)
	return cfg
}

func LoadIMDBReadConfig(viper *viper.Viper) *IMDBReadConfig {
	cfg := &IMDBReadConfig{
		Schema:   viper.GetString("schema"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
	fmt.Printf("LoadIMDBReadConfig:%#v\n", cfg)
	return cfg
}

//ill:sick@tcp(localhost:3306)/ill?parseTime=true

func (cfg *DBWriteConfig) InitDBWrite() *gorm.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	dbWrite, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("init wirtedb err:[%v]-[%v]", err, dataSourceName)
		panic("InitDBWrite error")
	}
	dbWrite.LogMode(true)
	dbWrite.Debug()
	return dbWrite
}

func (cfg *DBReadConfig) InitDBRead() *gorm.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	dbRead, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("init readdb err:[%v]-[%v]", err, dataSourceName)
		panic("InitDBRead error")
	}
	dbRead.LogMode(true)
	dbRead.Debug()
	return dbRead
}

func (cfg *IMDBReadConfig) InitIMDBRead() *gorm.DB {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local`, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	dbRead, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Printf("init im readdb err:[%v]-[%v]", err, dataSourceName)
		panic("InitIMDBRead error")
	}
	dbRead.LogMode(true)
	dbRead.Debug()
	return dbRead
}
