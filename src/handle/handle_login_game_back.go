package handle

import (
	"fmt"
	"message"
	"network"
)

type HandleLoginGameBack struct {
	AccountId uint32
	Timesec   uint32
	Channelid int32
}

func (this HandleLoginGameBack) NewMsg() network.IMsg {
	return new(message.MsgLoginGameBack)
}

func (this *HandleLoginGameBack) Init() {
	network.MsgFactoryMgr.Register(21004, this)
}

func (this *HandleLoginGameBack) MsgCallBack(client *network.ClientNet, imsg network.IMsg) {
	if msg, ok := imsg.(*message.MsgLoginGameBack); ok {
		if msg.Count == 0 { // 没有角色 1006
			fmt.Println("Login Back", msg, client.Conn)
			msg := message.MsgCreateRole{network.MsgHead{4, 1006}, client.RoleName, 1, 1, 112, 1, 1, "1", 1, "1", 1, "1", "1"}
			client.SendMsg(&msg)
			fmt.Println("Create Role", msg, client.Conn)
		} else { // 登录账号  1005
			reqmsg := message.MsgLoginRoleReq{network.MsgHead{4, 1005}, msg.Role[0].Id, 1}
			client.SendMsg(&reqmsg)
			fmt.Println("Login Role", reqmsg)
		}
	}
}

var HandleLoginGameBackIst HandleLoginGameBack
