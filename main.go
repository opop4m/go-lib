package main

import (
	"encoding/json"

	"github.com/opop4m/go-log/log"
)

func main() {
	conf := &log.LogConf{}
	conf.Dir = "bin/logs"
	conf.Debug = true
	body, _ := json.Marshal(&conf)
	var confMap map[string]interface{}
	json.Unmarshal(body, &confMap)

	log.InitLog(conf.Debug, "pid", conf.Dir, confMap)

	log.Info("test 11")
	log.Info("test %v", conf)
	// select {}
}
