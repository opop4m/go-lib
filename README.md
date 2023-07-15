# go-log

## Getting Started

```
import (
    "github.com/opop4m/go-lib/mgo"
)

    uri := "mongodb://xxxx:password@127.0.0.1:27017/dbName?authSource=admin"
	dbName := "dbName"
	mgo.InitMongoDB(uri, dbName)
	c := mgo.GetMongoDB().C("user")
	var results []map[string]interface{}
	c.Find(nil).All(&results)
	log.Info("res: %v", results)
```

https://pkg.go.dev/github.com/globalsign/mgo