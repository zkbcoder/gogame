package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgRegisterResult struct {
	Head   network.MsgHead
	Result int32
	Code   int32
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgRegisterResult) Packet() *bytes.Buffer {
	return nil
}

func (msg *MsgRegisterResult) UnPacket(buf []byte) {
	network.UnPacket(buf, &msg.Head.MsgLen)
	network.UnPacket(buf[2:], &msg.Head.MsgId)
	network.UnPacket(buf[4:], &msg.Result)
	network.UnPacket(buf[8:], &msg.Code)
	//fmt.Println("msg=", msg)
}
