package pkg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"runtime"
)

type LoggerConfig struct {
	Writers        string
	LoggerLevel    string
	LoggerFile     string
	LogFormatText  bool
	RollingPolicy  string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
}

func LoadLoggerConfig(viper *viper.Viper) *LoggerConfig {
	cfg := &LoggerConfig{
		Writers:        viper.GetString("writers"),
		LoggerLevel:    viper.GetString("logger_level"),
		LoggerFile:     viper.GetString("logger_file"),
		LogFormatText:  viper.GetBool("log_format_text"),
		RollingPolicy:  viper.GetString("rollingPolicy"),
		LogRotateDate:  viper.GetInt("log_rotate_date"),
		LogRotateSize:  viper.GetInt("log_rotate_size"),
		LogBackupCount: viper.GetInt("log_backup_count"),
	}
	fmt.Printf("LoadLoggerConfig:%#v\n", cfg)
	return cfg
}

func (cfg *LoggerConfig) InitLogger() {
	//passLagerCfg := log.PassLagerCfg{
	//	Writers:        cfg.Writers,
	//	LoggerLevel:    cfg.LoggerLevel,
	//	LoggerFile:     cfg.LoggerFile,
	//	LogFormatText:  cfg.LogFormatText,
	//	RollingPolicy:  cfg.RollingPolicy,
	//	LogRotateDate:  cfg.LogRotateDate,
	//	LogRotateSize:  cfg.LogRotateSize,
	//	LogBackupCount: cfg.LogBackupCount,
	//}
	//if err := log.InitWithConfig(&passLagerCfg); err != nil {
	//	panic("InitLogger error")
	//}

	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000000-0700 MST",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)
	logFile := &lumberjack.Logger{
		Filename:   cfg.LoggerFile,
		MaxSize:    cfg.LogRotateSize, // megabytes
		MaxBackups: cfg.LogBackupCount,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	level, err := log.ParseLevel(cfg.LoggerLevel)
	if nil != err {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}
