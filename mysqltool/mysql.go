package mysqltool

import (
	slog "log"
	"os"
	"sync"
	"time"

	"github.com/opop4m/go-lib/log"

	"github.com/panjf2000/ants/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

var mysqlDB *gorm.DB
var onceMysql sync.Once

type QueueFuc func()

func InitMysql(dsn string) {
	onceMysql.Do(func() {
		newLogger := New(
			slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
			Config{
				SlowThreshold: time.Second / 2, // Slow SQL threshold
				LogLevel:      l.Error,         // Log level
				Colorful:      true,            // Disable color
			},
		)
		// dsn := GlobalConfig.Mysql.Dsn()
		MysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		db, _ := MysqlDB.DB()
		db.SetMaxIdleConns(512)
		db.SetMaxOpenConns(512)
		db.SetConnMaxLifetime(30 * time.Second)
		mysqlDB = MysqlDB
		if err != nil {
			log.Error("failed to connect database:", err)
			os.Exit(1)
		} else {
			log.Info("connect mysql database success")
		}

		pool, _ = ants.NewPool(200)
		pool.Running()
	})
}

var pool *ants.Pool

func GetMysql() *gorm.DB {
	if mysqlDB != nil {
		return mysqlDB
	}
	panic("mysql db is not init.")
}

func ExecQueueForMysql(f QueueFuc) {
	if err := pool.Submit(func() {
		f()
	}); err != nil {
		log.Error("mysql queue error: %s", err)
	}
}
