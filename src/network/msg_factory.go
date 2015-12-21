// 消息工程
package network

type FunCreateMsg func() IMsg

type IHandle interface {
	NewMsg() IMsg
	MsgCallBack(client *ClientNet, iMsg IMsg)
}

type MsgFactory struct {
	AllMsg map[uint16]IHandle
}

func (this MsgFactory) GetMsg(msgId uint16) IMsg {
	hd := this.AllMsg[msgId]
	if nil != hd {
		return hd.NewMsg()
	}
	return nil
}

func (this MsgFactory) CallBack(msgId uint16, client *ClientNet, msg IMsg) {
	hd := this.AllMsg[msgId]
	if nil != hd {
		hd.MsgCallBack(client, msg)
	}
}

func (this *MsgFactory) Register(msgId uint16, fun IHandle) {
	this.AllMsg[msgId] = fun
}

var MsgFactoryMgr MsgFactory

func (this *MsgFactory) Init() {
	this.AllMsg = make(map[uint16]IHandle)
}
