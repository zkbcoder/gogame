// 消息工程
package network

// 去掉
//type FunCreateMsg func() IMsg

//type IHandle interface {
//	NewMsg() IMsg
//	MsgCallBack(client *NetConn)
//}

//type MsgFactory struct {
//	AllMsg map[uint16]IMsg
//}

//func (this MsgFactory) GetMsg(msgId uint16) IMsg {
//	//hd := this.AllMsg[msgId]
//	//if nil != hd {
//	//	return hd.NewMsg()
//	//}
//	return this.AllMsg[msgId]
//}

//func (this MsgFactory) CallBack(msgId uint16, client *NetConn, msg IMsg) {
//	hd := this.AllMsg[msgId]
//	if nil != hd {
//		hd.MsgCallBack(client, msg)
//	}
//}

//func (this *MsgFactory) Register(msgId uint16, fun IMsg) {
//	this.AllMsg[msgId] = fun
//}

//var MsgFactoryMgr MsgFactory

//func (this *MsgFactory) Init() {
//	this.AllMsg = make(map[uint16]IMsg)
//}
