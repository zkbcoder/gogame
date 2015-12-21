package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgServerList struct {
	Head network.MsgHead
	Name string
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgServerList) Packet() *bytes.Buffer {
	outbuff := bytes.NewBuffer([]byte{})
	// msg.Head.MsgLen = msg.Length()
	network.Append(outbuff, msg.Head.MsgLen)
	network.Append(outbuff, msg.Head.MsgId)
	network.AppendString(outbuff, msg.Name)
	return outbuff
}

func (msg *MsgServerList) UnPacket(buf []byte) {

}
