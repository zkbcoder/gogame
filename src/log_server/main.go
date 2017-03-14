// httpfiles.go
package main

import (
	"fmt"
	"frameworks/util"
	"net/http"
	"sync"
)

// 我也改  
// 开启http服务
func StartHttp(addr string /*, ch chan def.MsgChan*/) {
	fmt.Println("StartHttp!!", addr)
	//	heartUrl = strHeart
	//	msgChan = ch
	// 第一个参数为客户的发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	//	http.HandleFunc("/", heartBeat)
	http.HandleFunc("/", actionHandle)

	// 服务器要监听的主机地址和端口
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("ListenAndServer error: ", err.Error())
	}
}

func actionHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Action Handle!")

	//	ret := define.HttpReturn{1, "UnKnown Error"} // 返
	// read body
	data := make([]byte, 1024*1024)
	n, _ := req.Body.Read(data)
	//	if n >= 1024 {
	//		fmt.Fprintf(w, ret.ToJson()) // 返回回去
	//		return
	//	}
	// 最后err会等于EOF
	cBody := data[:n]
	strBody := string(cBody)
	util.RunRecord("%s", strBody)
	fmt.Println(strBody)

	// 解析action
	req.ParseForm()
	action, found := req.Form["action"]
	fmt.Println(action, found)

	// 消息解析和处理并返回
	resStr := "ok"
	fmt.Fprintf(w, resStr) // 200表示成功
}

func main() {
	go StartHttp("127.0.0.1:5001")

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() // 主线程不退出
}
