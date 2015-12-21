package handle

import (
	"message"
	"network"
)

type HandleHitMessage struct {
}

func (this HandleHitMessage) NewMsg() network.IMsg {
	return new(message.MsgHitMessage)
}

func (this *HandleHitMessage) Init() {
	network.MsgFactoryMgr.Register(21402, this)
}

func (this *HandleHitMessage) MsgCallBack(client *network.ClientNet, iMsg network.IMsg) {

}

var HandleHitMessageIst HandleHitMessage
