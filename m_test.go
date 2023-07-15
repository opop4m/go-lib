package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/opop4m/go-lib/log"
	"github.com/opop4m/go-lib/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

func BenchmarkMgo(b *testing.B) {
	InitMgo()
	for i := 0; i < b.N; i++ {
		query()
	}
}

func TestQeury(t *testing.T) {
	InitMgo()
	query()
	query()
	query()
	query()
}

func TestInitMgo(t *testing.T) {
	InitMgo()
	c := mgo.GetMongoDB().C("user")
	// var results []map[string]interface{}
	// c.Find(nil).All(&results)
	// log.Info("res: %v", results)

	for i := 0; i < 100000; i++ {
		u := User{
			Name: fmt.Sprintf("name_%d", i),
			Age:  fmt.Sprintf("age_%d", i),
		}
		if err := c.Insert(&u); err != nil {
			log.Error("err: %v", err)
		}
	}
}

type User struct {
	Name string
	Age  string
}

func query() {
	i := RandInt32(0, 51181)
	c := mgo.GetMongoDB().C("user")
	f := bson.M{"name": fmt.Sprintf("name_%d", i)}
	var querys []User
	if err := c.Find(f).All(&querys); err != nil {
		log.Error(err.Error())
	}
	log.Info("res: %v", querys)
}

// func query2() {
// 	i := RandInt32(0, 51181)
// 	c := mgo.GetMongoDB().C("user")
// 	f := bson.M{"name": fmt.Sprintf("name_%d", i)}
// 	var querys []User
// 	if err := c.Find(f).All2(&querys); err != nil {
// 		log.Error(err.Error())
// 	}
// 	log.Info("res: %v", querys)
// }

var RandPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

func RandInt32(min int32, max int32) int32 {
	if min == max {
		return min
	}
	r := RandPool.Get().(*rand.Rand)
	defer RandPool.Put(r)
	return min + r.Int31n(max-min)
}

func InitMgo() {
	uri := "mongodb://root:opop4M@127.0.0.1:37017/test?authSource=admin"
	dbName := "test"
	mgo.InitMongoDB(uri, dbName)
}
