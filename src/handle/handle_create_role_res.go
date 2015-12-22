package handle

import (
	"message"
	"network"
)

type HandleCreateRoleRes struct {
}

func (this HandleCreateRoleRes) NewMsg() network.IMsg {
	return new(message.MsgCreateRoleRes)
}

func (this *HandleCreateRoleRes) Init() {
	network.MsgFactoryMgr.Register(21006, this)
}

func (this *HandleCreateRoleRes) MsgCallBack(client *network.NetConn, iMsg network.IMsg) {
	if msg, ok := iMsg.(*message.MsgCreateRoleRes); ok {
		reqmsg := message.MsgLoginRoleReq{network.MsgHead{4, 1005}, msg.RoleId, 1}
		client.SendMsg(&reqmsg)
	}
}

var HandleCreateRoleResIst HandleCreateRoleRes
