package handle

import (
	"fmt"
	"message"
	"network"
	"time"
)

type HandleLoginOk struct {
}

func (this HandleLoginOk) NewMsg() network.IMsg {
	return new(message.MsgLoginOk)
}

func (this *HandleLoginOk) Init() {
	network.MsgFactoryMgr.Register(21200, this)
}

func (this *HandleLoginOk) MsgCallBack(client *network.ClientNet, msg network.IMsg) {
	end := time.Now()
	result := end.Sub(client.StartTime).Nanoseconds() / 1000000
	fmt.Println(client.RoleName, " LoginTime =", result)
	client.Log.Info("info=%s%s%d", client.RoleName, " LoginTime =", result)
}

var HandleLoginOkIst HandleLoginOk
