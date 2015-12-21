package message

import (
	"bytes"
	"fmt"
	"network"
)

// 请求服务列表
type MsgRegister struct {
	Head     network.MsgHead
	Account  string
	Password string
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgRegister) Packet() *bytes.Buffer {
	outbuff := bytes.NewBuffer([]byte{})
	network.Append(outbuff, msg.Head.MsgLen)
	network.Append(outbuff, msg.Head.MsgId)
	network.AppendString(outbuff, msg.Account)
	network.AppendString(outbuff, msg.Password)
	fmt.Println(outbuff.Bytes())
	return outbuff
}

func (msg *MsgRegister) UnPacket(buf []byte) {

}
