// ClientNet 客户端网络类
package network

import (
	//	"fmt"
	//	"net"
	"time"

	"github.com/astaxie/beego/logs"
)

type ClientNet struct {
	NetConn
	ServerAddr string // 服务端IP:Port
	StartTime  time.Time
	RoleName   string
	Log        *logs.BeeLogger
}
