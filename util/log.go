/********************************************************************
                Copyright (af) 2016 kb code

    创建日期： 2016年8月25日20时49分
    文件名称： log.go
    说   明： 公共日志

    当前版本： 1.00
    作   者： zkb
    概   述：
********************************************************************/
package util

import (
	"os"

	"github.com/astaxie/beego/logs"
)

var runLog *logs.BeeLogger
var errLog *logs.BeeLogger

func InitLog() {
	os.Mkdir("log", 0777) // 创建目录
	// log初始化
	runLog = logs.NewLogger(10000)
	runLog.SetLogger("file", `{"filename":"log/run.log"}`)
	runLog = runLog.Async()
	// errlog初始化
	errLog = logs.NewLogger(10000)
	errLog.SetLogger("file", `{"filename":"log/err.log"}`)
	errLog = errLog.Async()
}

// 运行时日志
func RunRecord(format string, v ...interface{}) {
	if nil == runLog {
		InitLog()
	}
	runLog.Info(format, v...)
}

// 错误时日志
func ErrRecord(format string, v ...interface{}) {
	if nil == errLog {
		InitLog()
	}
	errLog.Info(format, v...)
}
