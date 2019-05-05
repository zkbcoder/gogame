// 同步socket
// 往一个channel里面写数据，由一个线程统一处理
package network

import (
	"net"
)
import "time"

type MsgData struct {
	net    *SynNet // 链接
	msg    IMsg    // 数据
	handle IHandle
}

type IReconnet interface {
	OnReconnect()
}

type SynNet struct {
	conn    net.Conn // 链接
	factory IFactory // 消息处理工厂
	ch      chan MsgData
	monitor IReconnet
}

func (this *SynNet) SendMsg(msg IMsg) {
	if nil != this.conn {
		this.conn.Write(msg.Packet().Bytes())
	}
}

func (this *SynNet) Close() {
	this.conn.Close()
}

////把数据写到一个协程[主协程]
func CoreRun() chan MsgData {
	ch := make(chan MsgData)

	go func() {
		for {
			data := <-ch // 从队列里面读取消息
			msg := data.msg
			handle := data.handle
			if nil != msg && nil != handle {
				//				fmt.Println("msg=", msg)
				handle.MsgCallBack(msg, data.net)
			} else {
				Log().Error("coreRun read err!")
			}

		}
	}()
	return ch
}

////把数据写到一个协程[主协程]
func CoreRunTime(callback func()) chan MsgData {
	ch := make(chan MsgData)

	go func() {
		for {
			select {
			case data := <-ch: // 从队列里面读取消息
				msg := data.msg
				handle := data.handle
				if nil != msg && nil != handle {
					//				fmt.Println("msg=", msg)
					handle.MsgCallBack(msg, data.net)
				} else {
					Log().Error("coreRun read err!")
				}
			case <-time.After(time.Second):
				callback()
			}
		}
	}()
	return ch
}

func (this *SynNet) read() {
	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := this.conn.Read(buffer)
		if err != nil {
			this.Reconnet() // 重连
			//			fmt.Println(this.conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		tmpBuffer = append(tmpBuffer, buffer[:n]...)
		var bRet bool
		for {
			tmpBuffer, bRet = this.dealMsg(tmpBuffer)
			if bRet {
				break
			}
		}
	}
}

func (this *SynNet) dealMsg(buf []byte) ([]byte, bool) {
	if len(buf) < 2 {
		return buf, true
	}
	// 先读取包大小
	var uLen uint16
	//	fmt.Println(buf, "--", uLen)
	UnPacket(buf, &uLen)
	//	fmt.Println(buf, "--", uLen, len(buf))
	if int(uLen) > len(buf) {
		return buf, true
	}
	// 开始解包
	var msgId uint16
	UnPacket(buf[2:], &msgId)

	msg, handle := this.factory.CreateMsg(msgId)
	if nil != msg && nil != handle {
		msg.UnPacket(buf)
		//		fmt.Println(msg)
		// 数据写入channel
		data := MsgData{net: this, msg: msg, handle: handle}
		this.ch <- data
	} else {
		Log().Error("No Deal Msg=%d", msgId)
	}

	buf = buf[uLen:]
	return buf, false
}

func (this *SynNet) Reconnet() {
	if this.monitor != nil {
		this.monitor.OnReconnect()
	}
}
