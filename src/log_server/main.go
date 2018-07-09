// httpfiles.go
package main

import (
	"flag"
	"fmt"
	"frameworks/util"
	"net/http"
	"sync"
)

var addr = flag.String("addr", "localhost:5001", "http server address")

// 开启http服务
func StartHttp(addr string /*, ch chan def.MsgChan*/) {
	fmt.Println("StartHttp!!", addr)

	http.HandleFunc("/", actionHandle)

	// 服务器要监听的主机地址和端口
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("ListenAndServer error: ", err.Error())
	}
}

func actionHandle(w http.ResponseWriter, req *http.Request) {
	data := make([]byte, 1024*1024)
	n, _ := req.Body.Read(data)
	// 最后err会等于EOF
	cBody := data[:n]
	strBody := string(cBody)
	util.RunRecord("%s", strBody)

	// 解析action
	req.ParseForm()

	// 消息解析和处理并返回
	resStr := "ok"
	fmt.Fprintf(w, resStr) // 200表示成功
}

func main() {
	go StartHttp(*addr)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() // 主线程不退出
}
