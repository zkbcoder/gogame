// GS
package main

import (
	//	"handle"
	"network"
)

func InitMsg() {
	network.MsgFactoryMgr.Init()

}

func main() {
	InitMsg()
	var net network.ServerNet
	// net.Port = ":9988"
	net.ServerRun(":9988")
}
