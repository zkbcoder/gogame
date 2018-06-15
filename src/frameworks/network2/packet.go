// 打包
package network2

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/astaxie/beego/logs"
)

// 使用大端编码
var (
	ByteOrderType = binary.BigEndian
)

type MsgHead struct {
	ReqId   uint16 // 消息包序列号
	MsgType uint8  // 消息类型 2 protobuf
	MsgLen  uint16
	MsgId   uint16
}

type IMsg interface {
	Packet() *bytes.Buffer
	UnPacket(buf []byte)
}

//type IHandle interface {
//	MsgCallBack(msg IMsg, conn *SynNet)
//}

//type INet interface {
//	OnConnected(conn *SynClient)
//	OnDisConnected(conn *SynClient)
//}

// 创建msg
//type IFactory interface {
//	CreateMsg(msgId uint16) (IMsg, IHandle)
//}

type NetConn struct {
	Conn net.Conn
}

// log
var log *logs.BeeLogger

func Log() *logs.BeeLogger {
	if nil == log {
		newLog := logs.NewLogger(10000)
		newLog.SetLogger("file", `{"filename":"run.log","maxdays":365}`)
		newLog.Async()
		newLog.EnableFuncCallDepth(true)
		log = newLog
	}
	return log
}

// 打包基本类型，除了string
func Append(buf *bytes.Buffer, v interface{}) {
	// 自动算大小
	err := binary.Write(buf, ByteOrderType, v)
	if nil != err {
		Log().Error("Append=[%s]", err.Error())
	}
	//	binary.BigEndian.PutUint16(buf.Bytes(), uint16(buf.Len()))
}

// 打包基本类型，除了string
func AppendBuff(buf *bytes.Buffer, buffdata []byte) {
	// 自动算大小
	var dataLen uint16
	dataLen = uint16(len(buffdata))
	Append(buf, dataLen)
	Append(buf, buffdata)
}

// 打包string  [字符串长度用uint16]
func AppendString(buf *bytes.Buffer, v string) {
	err := binary.Write(buf, ByteOrderType, uint16(len(v)))
	if nil != err {
		Log().Error("AppendStringLen=[%s]", err.Error())
	}
	err = binary.Write(buf, ByteOrderType, []byte(v))
	if nil != err {
		Log().Error("AppendString=[%s]", err.Error())
	}
	//	binary.BigEndian.PutUint16(buf.Bytes(), uint16(buf.Len()))
}

// 解包
func UnPacket(buf []byte, v interface{}) []byte {
	bytesBuffer := bytes.NewBuffer(buf)
	err := binary.Read(bytesBuffer, ByteOrderType, v)
	if nil != err {
		Log().Error("UnPacket=[%s] [%v]", err.Error(), bytesBuffer)
	}

	buflen := intDataSize(v)
	if int(buflen) > len(buf) {
		Log().Error("uPack size error")
		return buf
	}
	buf = buf[buflen:]
	return buf
}

// 解包
func UnPacketBuff(buf []byte, outbuf *[]byte) []byte {
	var dataLen uint16
	buf = UnPacket(buf, &dataLen)
	(*outbuf) = buf[:dataLen]
	buf = buf[dataLen:]
	return buf
}

//// 字符串解包
func UnPacketString(buf []byte, str *string) []byte {
	// 先读取长度
	var strLen uint16
	buf = UnPacket(buf, &strLen)

	// 再加载内容
	if int(strLen) > len(buf) {
		Log().Error("uPack size error")
		return buf
	}
	*str = string(buf[:strLen])
	buf = buf[strLen:]
	return buf
}

func intDataSize(data interface{}) uint16 {
	switch data := data.(type) {
	case int8, *int8, *uint8:
		return 1
	case []int8:
		return (uint16)(len(data))
	case []uint8:
		return (uint16)(len(data))
	case int16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * (uint16)(len(data))
	case []uint16:
		return 2 * (uint16)(len(data))
	case int32, *int32, *uint32:
		return 4
	case []int32:
		return 4 * (uint16)(len(data))
	case []uint32:
		return 4 * (uint16)(len(data))
	case int64, *int64, *uint64:
		return 8
	case []int64:
		return 8 * (uint16)(len(data))
	case []uint64:
		return 8 * (uint16)(len(data))
	}
	return 0
}
