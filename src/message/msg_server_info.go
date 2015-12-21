package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgServerInfo struct {
	Head      network.MsgHead
	AccountId uint32
	Timesec   uint32
	Channelid int32
	Port      int32
	Ip        string
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgServerInfo) Packet() *bytes.Buffer {
	return nil
}

func (msg *MsgServerInfo) UnPacket(buf []byte) {
	network.UnPacket(buf, &msg.Head.MsgLen)
	network.UnPacket(buf[2:], &msg.Head.MsgId)
	network.UnPacket(buf[4:], &msg.AccountId)
	network.UnPacket(buf[8:], &msg.Timesec)
	network.UnPacket(buf[12:], &msg.Channelid)
	network.UnPacket(buf[16:], &msg.Port)
	network.UnPacketString(buf[20:], &msg.Ip)
	// fmt.Println("serverinfo=", msg, &msg)
}
