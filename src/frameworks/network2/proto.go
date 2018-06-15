// protobuf 消息包
package network2

import (
	"bytes"
)

type ProtoMsg struct {
	Head MsgHead // 消息头
}

func (this *ProtoMsg) Pack() []byte {
	outbuff := bytes.NewBuffer([]byte{})
	this.Head.MsgLen = 7
	Append(outbuff, this.Head.ReqId)
	Append(outbuff, this.Head.MsgType)
	Append(outbuff, this.Head.MsgLen)
	Append(outbuff, this.Head.MsgId)

	return outbuff.Bytes()
}

func (this *ProtoMsg) UnPack(buf []byte) {
	buf = UnPacket(buf, &this.Head.ReqId)
	buf = UnPacket(buf, &this.Head.MsgType)
	buf = UnPacket(buf, &this.Head.MsgLen)
	buf = UnPacket(buf, &this.Head.MsgId)
}
