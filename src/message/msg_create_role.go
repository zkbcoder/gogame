package message

import (
	"bytes"
	"network"
)

// 请求服务列表
type MsgCreateRole struct {
	Head          network.MsgHead
	Name          string
	Sex           int32
	ServerId      int32
	Hero          int32
	ChannelId     int32
	Appid         int32
	Channel       string
	AppfameUserId int32
	Udid          string
	ModelType     int32
	Edition       string
	Ip            string
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgCreateRole) Packet() *bytes.Buffer {
	outbuff := bytes.NewBuffer([]byte{})
	network.Append(outbuff, msg.Head.MsgLen)
	network.Append(outbuff, msg.Head.MsgId)
	network.AppendString(outbuff, msg.Name)
	network.Append(outbuff, msg.Sex)
	network.Append(outbuff, msg.ServerId)
	network.Append(outbuff, msg.Hero)
	network.Append(outbuff, msg.ChannelId)
	network.Append(outbuff, msg.Appid)
	network.AppendString(outbuff, msg.Channel)
	network.Append(outbuff, msg.AppfameUserId)
	network.AppendString(outbuff, msg.Udid)
	network.Append(outbuff, msg.ModelType)
	network.AppendString(outbuff, msg.Edition)
	network.AppendString(outbuff, msg.Ip)
	return outbuff
}

func (msg *MsgCreateRole) UnPacket(buf []byte) {

}
