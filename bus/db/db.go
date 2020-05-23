package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"oceanEngineService/bus/entity/errmsg"
	"oceanEngineService/core/pkg"
	"path"
	"runtime"
)

var (
	dbr *gorm.DB
	dbw *gorm.DB

	TestDSN = "vlm_user:@tcp(localhost:3306)/vlm_mine?parseTime=true&loc=Local"
)

func InitDBFromCfg() {
	dbw = pkg.DBWrite
	dbr = pkg.DBRead
}

// BeginTx start new transaction
func BeginTx() *gorm.DB {
	return dbw.Begin()
}

// LockTable lock table in transaction
func LockTable(txdb *gorm.DB, tableName string) error {
	if nil == txdb {
		return errmsg.ErrInternalServer
	}

	var tmp = make([]int, 0, 1)
	return txdb.Raw(fmt.Sprintf("select distinct 1 from %v for update", tableName)).Scan(&tmp).Error
}

// InitLog ...
func InitLog() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000000-0700 MST",
		// FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) { // 返回的是函数名和文件位置
			//repopath := fmt.Sprintf("%s/src/github.com/bob", os.Getenv("GOPATH"))
			//filename := strings.Replace(f.File, repopath, "", -1)
			//filenames := strings.Split(f.File, string(os.PathSeparator))
			//return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filenames[len(filenames)-1], f.Line)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
		},
	})
}

// InitDB by DSN name, e.g. : dsn = ill:sick@tcp(localhost:3306)/ill?parseTime=true
// parseTime=true will affect time.Time colume get/put, we need it
func InitDB(dsn string) {
	InitLog()
	var err error
	dbr, err = gorm.Open("mysql", dsn)
	if nil != err {
		log.Fatal("Init DB failed", err)
		return
	}
	dbw, err = gorm.Open("mysql", dsn)
	if nil != err {
		log.Fatal("Init DB failed", err)
		return
	}
	dbr.Debug()
	dbr.LogMode(true)
}
