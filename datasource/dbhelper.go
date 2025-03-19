package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"learnLottery/conf"
	"log"
	"sync"
)

var dbLock sync.Mutex
var masterInstance *xorm.Engine

func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if masterInstance != nil {
		masterInstance = NewDbMaster()
	}
	return masterInstance
}

func NewDbMaster() *xorm.Engine {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)

	instance, err := xorm.NewEngine(conf.DriverName, sourceName)
	if err != nil {
		log.Fatal("dbhelper.InstanceDbMaster NewEngine error ", err)
		return nil
	}
	instance.ShowSQL(true)
	masterInstance = instance
	return masterInstance
}
