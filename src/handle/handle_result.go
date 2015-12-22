package handle

import (
	"message"
	"network"
)

type HandleResult struct {
}

func (this HandleResult) NewMsg() network.IMsg {
	return new(message.MsgRegisterResult)
}

func (this *HandleResult) Init() {
	network.MsgFactoryMgr.Register(21099, this)
}

func (this *HandleResult) MsgCallBack(client *network.NetConn, msg network.IMsg) {

}

var HandleResultIst HandleResult
