package db_helper

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/3 15:13
//

import (
	"fmt"
	"gin-server/app/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 12:12
//

var dbEngine *xorm.Engine

func SetUpDB() {
	if dbEngine != nil {
		return
	}
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Instance.DB.Username,
		config.Instance.DB.Password,
		config.Instance.DB.Host,
		config.Instance.DB.Port,
		config.Instance.DB.Database,
		config.Instance.DB.Charset,
	)
	if engine, err := xorm.NewEngine(config.Instance.DB.Engine, sourceName); err != nil {
		log.Fatalf("db_helper.InitDb(%s),%v\n", sourceName, err)
	} else {
		dbEngine = engine
	}

	if config.Instance.DB.MaxIdleConns > 0 {
		dbEngine.SetMaxIdleConns(config.Instance.DB.MaxIdleConns)
	}

	if config.Instance.DB.MaxOpenConns > 0 {
		dbEngine.SetMaxOpenConns(config.Instance.DB.MaxOpenConns)
	}
	if config.Instance.DB.ConnMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(config.Instance.DB.ConnMaxLifetime))
	}
}

func GetDb() *xorm.Engine {
	if dbEngine == nil {
		SetUpDB()
	}
	return dbEngine
}
