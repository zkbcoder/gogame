package network

import "net"

//	"github.com/astaxie/beego/logs"

// 同步服务器
type SynServer struct {
}

func (this *SynServer) Run(port string, msgFactory IFactory, chMsg chan MsgData) {
	netListen, err := net.Listen("tcp", port)
	if nil != err {
		Log().Error("Listen:[%s]", err.Error())
		return
	}
	defer netListen.Close()

	Log().Info("Server Run!") // 服务器启动记录
	// 服务器创建一个channel,所有消息往这个channel写
	//chMsg := CoreRun()
	for {
		connNet, err := netListen.Accept()
		if err != nil {
			Log().Error("Accept:[%s]", err.Error())
			continue
		}

		clietConn := SynNet{conn: connNet, factory: msgFactory, ch: chMsg}
		Log().Info("Run:[%s]", connNet.RemoteAddr().String())
		go clietConn.read()
	}
}
