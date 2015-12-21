package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgLoginRoleReq struct {
	Head     network.MsgHead
	RoleId   int32
	ServerId int32
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgLoginRoleReq) Packet() *bytes.Buffer {
	outbuff := bytes.NewBuffer([]byte{})
	network.Append(outbuff, msg.Head.MsgLen)
	network.Append(outbuff, msg.Head.MsgId)
	network.Append(outbuff, msg.RoleId)
	network.Append(outbuff, msg.ServerId)
	// fmt.Println("req login", msg)
	return outbuff
}

func (msg *MsgLoginRoleReq) UnPacket(buf []byte) {

}
