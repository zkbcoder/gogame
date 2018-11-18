package db

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/ziutek/mymysql/autorc"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
)

var querylog *logs.BeeLogger
var errlog *logs.BeeLogger

func QueryLog() *logs.BeeLogger {
	if nil == querylog {
		os.MkdirAll("logs/sql/", 0777)
		newLog := logs.NewLogger(10000)
		newLog.SetLogger("file", `{"filename":"logs/sql/sqlquery.log","maxdays":365}`)
		newLog.Async()
		//	newLog.EnableFuncCallDepth(true)
		querylog = newLog
	}
	return querylog
}

func ErrLog() *logs.BeeLogger {
	if nil == errlog {
		os.MkdirAll("logs/sqlerror/", 0777)
		newLog := logs.NewLogger(10000)
		newLog.SetLogger("file", `{"filename":"logs/sqlerror/sqlerror.log","maxdays":365}`)
		newLog.Async()
		//newLog.EnableFuncCallDepth(true)
		errlog = newLog
	}
	return errlog
}

type MyDB struct {
	DB *autorc.Conn
}

func (this *MyDB) Query(sql string, params ...interface{}) (rows []mysql.Row, res mysql.Result, err error) {
	rows, res, err = this.DB.Query(sql, params...)
	if err != nil {
		errsql := fmt.Sprintf(sql, params...)
		ErrLog().Error("[%v][%s]", err, errsql)
	} else {
		QueryLog().Error(sql, params...)
	}
	return
}
