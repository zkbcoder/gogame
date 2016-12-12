////// socket
package network

//import (
//	"fmt"
//	"net"
//)

//type NetConn struct {
//	conn    net.Conn // 链接
//	factory IFactory
//}

//type TransData struct {
//	netConn *NetConn
//	msg     IMsg
//}

//func (this *NetConn) SendMsg(msg IMsg) {
//	this.conn.Write(msg.Packet().Bytes())
//}

//func (this *NetConn) Close() {
//	this.conn.Close()
//}

//func (this *NetConn) ClientRun(addr string) {
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
//	if err != nil {
//		fmt.Println("Fatal error:", err.Error())
//		return
//	}

//	conn, err := net.DialTCP("tcp", nil, tcpAddr)
//	if err != nil {
//		fmt.Println("Fatal error: ", err.Error())
//		return
//	}

//	ch := this.coreRun()
//	this.conn = conn
//	go this.reader(ch)
//	fmt.Println("conn ok", conn, this.conn)
//}

////func (this *NetConn) Reader() {
////	//声明一个临时缓冲区，用来存储被截断的数据
////	tmpBuffer := make([]byte, 0)

////	buffer := make([]byte, 1024)
////	for {
////		n, err := this.conn.Read(buffer)
////		if err != nil {
////			fmt.Println(this.conn.RemoteAddr().String(), " connection error: ", err)
////			return
////		}

////		tmpBuffer = append(tmpBuffer, buffer[:n]...)
////		var bRet bool
////		for {
////			tmpBuffer, bRet = this.dealMsg(tmpBuffer)
////			if bRet {
////				break
////			}
////		}
////	}
////}

////把数据写到一个协程
//func (this *NetConn) coreRun() chan TransData {
//	ch := make(chan TransData)

//	go func() {
//		for {

//			transData := <-ch // 从队列里面读取消息
//			msg := transData.msg
//			if nil != msg {
//				fmt.Println(msg)
//				msg.MsgCallBack(transData.netConn)
//			} else {

//			}
//		}
//	}()
//	return ch
//}

//func (this *NetConn) ServerRun(port string) {
//	netListen, err := net.Listen("tcp", port)
//	if nil != err {
//		fmt.Println(err)
//		return
//	}
//	defer netListen.Close()

//	//	log := logs.NewLogger(10000)
//	//	log.SetLogger("file", `{"filename":"test.log"}`)
//	//	log = log.Async()

//	fmt.Println("111", this)
//	ch := this.coreRun()
//	for {
//		var clietConn NetConn
//		conn, err := netListen.Accept()
//		fmt.Println("Accept Err ", err, conn)
//		if err != nil {
//			fmt.Println("Accept Err ", err)
//			continue
//		}

//		clietConn.conn = conn
//		fmt.Println(conn.RemoteAddr().String(), "tcp connet success!")
//		go clietConn.reader(ch)
//	}
//	fmt.Println("222")
//}

//func (this *NetConn) reader(ch chan TransData) {
//	//声明一个临时缓冲区，用来存储被截断的数据
//	tmpBuffer := make([]byte, 0)

//	buffer := make([]byte, 1024)
//	for {
//		n, err := this.conn.Read(buffer)
//		if err != nil {
//			fmt.Println(this.conn.RemoteAddr().String(), " connection error: ", err)
//			return
//		}

//		tmpBuffer = append(tmpBuffer, buffer[:n]...)
//		//		fmt.Println("buf=", tmpBuffer)
//		var bRet bool
//		for {
//			tmpBuffer, bRet = this.dealMsg(tmpBuffer, ch)
//			if bRet {
//				break
//			}
//		}
//	}
//}

//func (this *NetConn) dealMsg(buf []byte, ch chan TransData) ([]byte, bool) {
//	if len(buf) < 2 {
//		return buf, true
//	}
//	// 先读取包大小
//	var uLen uint16
//	//	fmt.Println(buf, "--", uLen)
//	UnPacket(buf, &uLen)
//	//	fmt.Println(buf, "--", uLen, len(buf))
//	if int(uLen) > len(buf) {
//		return buf, true
//	}
//	// 开始解包
//	var msgId uint16
//	UnPacket(buf[2:], &msgId)

//	// TODO...
//	msg := this.factory.CreateMsg(msgId)
//	if nil != msg {

//		msg.UnPacket(buf)
//		// 放到处理ch里面
//		transData := TransData{
//			netConn: this,
//			msg:     msg,
//		}
//		ch <- transData
//	} else {
//		fmt.Println("No Deal Msg=%d", msgId)
//	}

//	buf = buf[uLen:]
//	return buf, false
//}
