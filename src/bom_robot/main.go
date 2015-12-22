package main

import (
	"encoding/xml"
	"fmt"
	"handle"
	"io/ioutil"
	"message"
	"network"
	"time"

	"github.com/astaxie/beego/logs"
)

type Config struct {
	Addr        string
	NumBySec    int
	Loop        int
	AccountBase string
	Password    string
	ServerName  string
}

func InitMsg() {
	network.MsgFactoryMgr.Init()
	handle.HandleServerInfoIst.Init()
	handle.HandleLoginGameBackIst.Init()
	handle.HandleResultIst.Init()
	handle.HandleCreateRoleResIst.Init()
	handle.HandleHitMessageIst.Init()
	handle.HandleLoginOkIst.Init()
}

func InitXml(config *Config) {
	content, err := ioutil.ReadFile("config.xml")
	if err != nil {
		fmt.Println("Config Err=", err)
	}

	err = xml.Unmarshal(content, config)
	if err != nil {
		fmt.Println("xml err=", err)
	}
}

func main() {
	var config Config
	InitXml(&config)
	fmt.Println("Config=", config)
	InitMsg()

	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"test.log"}`)

	for i := 0; i < config.Loop; i++ {
		for j := 0; j < config.NumBySec; j++ {
			var client network.ClientNet
			client.Child = client
			client.StartTime = time.Now()
			client.Log = log
			account := fmt.Sprintf("%s%d", config.AccountBase, (i+1)*10000+j) // "kb2014002"
			client.RoleName = account
			log.Info("Begin Time=%s Name=%s", client.StartTime.String(), account)

			client.ServerAddr = config.Addr //"192.168.208.175:9987"
			client.ClientRun(config.Addr)

			// 请求服务器列表
			password := config.Password     //"1234567890"
			serverName := config.ServerName //"哈哈哈"
			// 注册 1003
			msg := message.MsgRegister{network.MsgHead{4, 1003}, account, password}
			client.SendMsg(&msg)

			// 账号登录 1002
			msgLogin := message.MsgLoginServer{network.MsgHead{4, 1002}, account, password, serverName}
			client.SendMsg(&msgLogin)
		}
	}

	for {
		time.Sleep(1 * 1e9)
	}
}
