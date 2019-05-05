// 同步型客户端
// 多个客户端接收消息往一个channel写数据，这个channel由一个协程读并且执行，保证同步
package network

import (
	"net"
	"time"
)

type SynClient struct {
	SynNet    // 继承SynNet
	Addr      string
	IsConnect bool
	listener  INet
}

func (this *SynClient) Init(msgFactory IFactory, ch chan MsgData) {
	this.factory = msgFactory
	this.ch = ch
	this.monitor = this
}

func (this *SynClient) SetListener(listener INet) {
	this.listener = listener
}

func (this *SynClient) Run(addr string) {
	this.IsConnect = false
	if this.factory == nil {
		Log().Error("error: NetClient not Init!")
		return
	}

	this.Addr = addr
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		go this.OnReconnect()
		Log().Error("run ResolveTCPAddr error:[%s][%s]", addr, err.Error())
		return
	}

	connNet, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		go this.OnReconnect()
		Log().Error("run DialTCP error:[%s][%s]", addr, err.Error())
		return
	}
	this.conn = connNet
	this.IsConnect = true
	if this.listener != nil {
		this.listener.OnConnected(this)
	}
	go this.read()
	Log().Error("[%s]conn ok", addr)
}

func (this *SynClient) OnReconnect() {
	if this.IsConnect && this.listener != nil {
		this.listener.OnDisConnected(this)
	}
	this.IsConnect = false
	time.Sleep(time.Second * 5) // 隔5s重试
	tcpAddr, err := net.ResolveTCPAddr("tcp4", this.Addr)
	if err != nil {
		go this.OnReconnect()
		Log().Error("reconnect ResolveTCPAddr error:[%s][%s]", this.Addr, err.Error())
		return
	}

	connNet, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		go this.OnReconnect()
		Log().Error("reconnect DialTCP error:[%s][%s]", this.Addr, err.Error())
		return
	}

	this.conn = connNet
	this.IsConnect = true
	if this.listener != nil {
		this.listener.OnConnected(this)
	}
	go this.read()
	Log().Error("[%s]reconn ok", this.Addr)
}

func (this *SynClient) GetIsConnect() bool {
	return this.IsConnect
}
