package message

import (
	"bytes"
	"fmt"
	"network"
)

// 请求服务列表
type MsgCreateRoleRes struct {
	Head   network.MsgHead
	Name   string
	RoleId int32
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgCreateRoleRes) Packet() *bytes.Buffer {
	return nil
}

func (msg *MsgCreateRoleRes) UnPacket(buf []byte) {
	network.UnPacket(buf, &msg.Head.MsgLen)
	network.UnPacket(buf[2:], &msg.Head.MsgId)
	nPos := 4
	network.UnPacketString(buf[nPos:], &msg.Name)
	nPos += 2
	nPos += len(msg.Name)
	network.UnPacket(buf[nPos:], &msg.RoleId)
	fmt.Println("CreateRoleRes=", msg)
}
