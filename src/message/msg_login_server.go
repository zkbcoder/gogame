package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgLoginServer struct {
	Head       network.MsgHead
	Account    string
	Password   string
	ServerName string
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgLoginServer) Packet() *bytes.Buffer {
	outbuff := bytes.NewBuffer([]byte{})
	network.Append(outbuff, msg.Head.MsgLen)
	network.Append(outbuff, msg.Head.MsgId)
	network.AppendString(outbuff, msg.Account)
	network.AppendString(outbuff, msg.Password)
	network.AppendString(outbuff, msg.ServerName)
	return outbuff
}

func (msg *MsgLoginServer) UnPacket(buf []byte) {

}
