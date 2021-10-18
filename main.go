package main

import (
	"encoding/json"
	"os"
	"time"

	slog "log"

	"github.com/opop4m/go-lib/log"
	"github.com/opop4m/go-lib/mgo"
	"github.com/opop4m/go-lib/myLog"
	l "gorm.io/gorm/logger"
)

func main() {
	conf := &log.LogConf{}
	conf.Dir = "bin/logs"
	conf.Debug = true
	body, _ := json.Marshal(&conf)
	var confMap map[string]interface{}
	json.Unmarshal(body, &confMap)

	log.InitLog(conf.Debug, "pid", conf.Dir, confMap)

	log.Info("test 111")
	log.Info("test %v", conf)

	uri := "mongodb://root:Ibdj782__@127.0.0.1:27017/remind?authSource=admin"
	dbName := "remind"
	mgo.InitMongoDB(uri, dbName)
	c := mgo.GetMongoDB().C("user")
	var results []map[string]interface{}
	c.Find(nil).All(&results)
	log.Info("res: %v", results)

	myLog.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		myLog.Config{
			SlowThreshold: time.Second / 2, // Slow SQL threshold
			LogLevel:      l.Error,         // Log level
			Colorful:      true,            // Disable color
		},
	)
	// select {}
}
