package network

import (
	"bytes"
	//	"encoding/binary"
	"fmt"

	"testing"
)

func TestPacket(t *testing.T) {
	outbuff := bytes.NewBuffer([]byte{})
	var val1 uint16
	var val2 uint32
	val1 = 2
	val2 = 90

	Append(outbuff, val1)
	Append(outbuff, val2)
	fmt.Println("out=", outbuff.Bytes())
	// binary.BigEndian.PutUint16(outbuff.Bytes(), 8)
	//Append(outbuff, 2)
	fmt.Println("out=", outbuff.Bytes())
	// binary.BigEndian.PutUint16(outbuff.Bytes(), 50000)
	// Append(outbuff, 99)
	fmt.Println("out=", outbuff.Bytes())

	//	binary.

	//	fmt.Println(MsgFactoryMgr.AllMsg)
	//MsgFactoryMgr.AllMsg = make(map[uint16]IMsg)

	// MsgFactoryMgr.AllMsg[1] = nil
}

func TestUnPacket(t *testing.T) {
	outbuff := bytes.NewBuffer([]byte{})
	var val1 uint16
	var val2 uint32
	val1 = 10
	val2 = 901

	Append(outbuff, val1)
	Append(outbuff, val2)
	AppendString(outbuff, "ffefd")
	val2 = 4200000000
	Append(outbuff, val2)

	val1 = 22
	val2 = 190
	fmt.Println(outbuff.Bytes())
	buf := outbuff.Bytes()
	buf = UnPacket(buf, &val1)
	buf = UnPacket(buf, &val2)
	//	UnPacket(buf[2:6], &val2)
	fmt.Println(val1, val2)
	var str string
	buf = UnPacketString(buf, &str)
	buf = UnPacket(buf, &val2)
	//	UnPacketString(buf[6:], &str)
	fmt.Println(str)
	fmt.Println(val2)

	//	bytesBuffer := bytes.NewBuffer(buf[2:5])
	//	bytesBuffer1 := bytes.NewBuffer(buf[2:6])
	//	fmt.Println(bytesBuffer.Bytes())
	//	fmt.Println(bytesBuffer1.Bytes())
	//	fmt.Println(cap(buf))
	//	fmt.Println(buf)
	//	fmt.Println(buf[2:6])

	//	buf = []
	//	fmt.Println("slice")
	//	fmt.Println(unsafe.Pointer(&buf))
	//	buf = buf[0:0]
	//	fmt.Println(unsafe.Pointer(&buf))
	//	fmt.Println(buf)
	//	fmt.Println(cap(buf))
	//	//	for i := 0; i < 100; i++ {
	//	//		buf = append(buf, 1)
	//	//		if len(buf) >= 10 {
	//	//			buf = buf[5:]
	//	//		}
	//	//		fmt.Println(buf)
	//	//	}
	//	fmt.Println(unsafe.Pointer(&buf))
}
