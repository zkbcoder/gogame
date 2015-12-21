package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgLoginOk struct {
	Head network.MsgHead
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgLoginOk) Packet() *bytes.Buffer {
	return nil
}

func (msg *MsgLoginOk) UnPacket(buf []byte) {

}
