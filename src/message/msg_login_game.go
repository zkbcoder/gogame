package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgLoginGame struct {
	Head      network.MsgHead
	AccountId uint32
	Timesec   uint32
	Channelid int32
	ServerId  int32
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgLoginGame) Packet() *bytes.Buffer {
	outbuff := bytes.NewBuffer([]byte{})
	network.Append(outbuff, msg.Head.MsgLen)
	network.Append(outbuff, msg.Head.MsgId)
	network.Append(outbuff, msg.AccountId)
	network.Append(outbuff, msg.Timesec)
	network.Append(outbuff, msg.Channelid)
	network.Append(outbuff, msg.ServerId)
	return outbuff
}

func (msg *MsgLoginGame) UnPacket(buf []byte) {

}
