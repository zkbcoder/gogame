package message

import (
	"bytes"
	"fmt"
	"network"
)

// 请求服务列表
type MsgHitMessage struct {
	Head     network.MsgHead
	Action   int32
	Sender   string
	Receiver string
	Content  string
	UserInfo string
	Title    string
	Level    string
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgHitMessage) Packet() *bytes.Buffer {
	return nil
}

func (msg *MsgHitMessage) UnPacket(buf []byte) {
	network.UnPacket(buf, &msg.Head.MsgLen)
	network.UnPacket(buf[2:], &msg.Head.MsgId)
	network.UnPacket(buf[4:], &msg.Action)
	nPos := 8
	network.UnPacketString(buf[nPos:], &msg.Sender)
	nPos = nPos + 2 + len(msg.Sender)
	network.UnPacketString(buf[nPos:], &msg.Receiver)
	nPos = nPos + 2 + len(msg.Receiver)
	network.UnPacketString(buf[nPos:], &msg.Content)
	nPos = nPos + 2 + len(msg.Content)
	network.UnPacketString(buf[nPos:], &msg.UserInfo)
	nPos = nPos + 2 + len(msg.UserInfo)
	network.UnPacketString(buf[nPos:], &msg.Title)
	nPos = nPos + 2 + len(msg.Title)
	network.UnPacketString(buf[nPos:], &msg.Level)

	fmt.Println("msg=", msg)
}

//func NewMsg1() network.IMsg {
//	return new(MsgHitMessage)
//}

//func Init1() {
//	network.MsgFactoryMgr.Register(21402, NewMsg1)
//}
