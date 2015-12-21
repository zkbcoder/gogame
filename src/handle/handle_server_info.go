package handle

import (
	"fmt"
	"message"
	"network"
	"time"
)

type HandleServerInfo struct {
	AccountId uint32
	Timesec   uint32
	Channelid int32
}

func (this HandleServerInfo) NewMsg() network.IMsg {
	return new(message.MsgServerInfo)
}

func (this *HandleServerInfo) Init() {
	network.MsgFactoryMgr.Register(21002, this)
}

func (this *HandleServerInfo) Reconect(client *network.ClientNet) {
	time.Sleep(1 * 10)
	// 重新链接
	client.Run()
	// 角色登录 1004
	//	serverName := "哈哈哈"
	// 注册 1003
	msg := message.MsgLoginGame{network.MsgHead{4, 1004}, this.AccountId, this.Timesec, this.Channelid, 1}
	// fmt.Println("1004=", msg)
	client.SendMsg(&msg)
}

func (this *HandleServerInfo) MsgCallBack(client *network.ClientNet, iMsg network.IMsg) {
	client.Close()

	// reconnect
	if a, ok := iMsg.(*message.MsgServerInfo); ok {
		a.Ip = a.Ip + fmt.Sprintf(":%d", a.Port)
		client.ServerAddr = a.Ip
		this.Channelid = a.Channelid
		this.AccountId = a.AccountId
		this.Timesec = a.Timesec
	}
	go this.Reconect(client)
}

var HandleServerInfoIst HandleServerInfo
