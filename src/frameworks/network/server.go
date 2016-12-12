// ServerNet 服务的类
package network

import "github.com/astaxie/beego/logs"

//import (
//	"fmt"
//	"net"
//)

type ServerNet struct {
	NetConn
	Port string // 端口格式 :9988
	Log  *logs.BeeLogger
}
