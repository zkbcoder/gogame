// network 定义文件
package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//	"io"
)

// 使用大端编码
var (
	ByteOrderType = binary.BigEndian
)

type MsgHead struct {
	MsgLen uint16
	MsgId  uint16
}

type IMsg interface {
	Packet() *bytes.Buffer
	UnPacket(buf []byte) // 有待改进
}

// 打包基本类型，除了string
func Append(buf *bytes.Buffer, v interface{}) {
	// 自动算大小
	err := binary.Write(buf, ByteOrderType, v)
	if nil != err {
		fmt.Println("Append=", err)
	}
	binary.BigEndian.PutUint16(buf.Bytes(), uint16(buf.Len()))
}

// 打包string  [字符串长度用uint16]
func AppendString(buf *bytes.Buffer, v string) {
	err := binary.Write(buf, ByteOrderType, uint16(len(v)))
	if nil != err {
		fmt.Println("AppendStringLen=", err)
	}
	err = binary.Write(buf, ByteOrderType, []byte(v))
	if nil != err {
		fmt.Println("AppendString=", err)
	}
	binary.BigEndian.PutUint16(buf.Bytes(), uint16(buf.Len()))
}

// 解包
func UnPacket(buf []byte, v interface{}) {
	bytesBuffer := bytes.NewBuffer(buf)
	err := binary.Read(bytesBuffer, ByteOrderType, v)
	if nil != err {
		fmt.Println("UnPacket=", err)
	}
}

//// 字符串解包
func UnPacketString(buf []byte, str *string) {
	// 先读取长度
	var strLen uint16
	UnPacket(buf, &strLen)
	// 再加载内容
	*str = string(buf[2 : strLen+2])
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
