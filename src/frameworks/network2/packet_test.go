package network2

import (
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	msg := ProtoMsg{
		Head: MsgHead{
			ReqId:   1,
			MsgType: 2,
			MsgLen:  10,
			MsgId:   1001,
		},
	}

	fmt.Println(msg)
	fmt.Println(msg.Pack())

	newMsg := ProtoMsg{}
	newMsg.UnPack(msg.Pack())
	fmt.Println(newMsg)

	//	Head.ReqId)
	//	Head.MsgType)
	//	Head.MsgLen)
	//	Head.MsgId)

}
