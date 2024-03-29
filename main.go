package main

import (
	"encoding/json"

	"github.com/opop4m/go-lib/log"
	"github.com/opop4m/go-lib/mgo"
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

	uri := "mongodb://root:opop4M@127.0.0.1:37017/test?authSource=admin"
	dbName := "test"
	mgo.InitMongoDB(uri, dbName)
	c := mgo.GetMongoDB().C("user")
	var results []map[string]interface{}
	c.Find(nil).All(&results)
	log.Info("res: %v", results)

	// mysqltool.New(
	// 	slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
	// 	mysqltool.Config{
	// 		SlowThreshold: time.Second / 2, // Slow SQL threshold
	// 		LogLevel:      l.Error,         // Log level
	// 		Colorful:      true,            // Disable color
	// 	},
	// )
	// mysqltool.InitMysql("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	// select {}
}
