package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http get方式取数据
func HttpGetCfg(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return url, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}

//----------------------------------------------------------------
// http心跳
// 使用方法：详见TestHttpHeartBeat

// 心跳请求
var beatUrl string

// addr:进程ip端口  url:心跳要请求的链接
// 例如: url="http://192.168.208.34:9970/?action=2&s_process_id=ch_server"
func StartHttpHeartBeat(addr string, url string) {
	fmt.Println("Start Heart Beat!")

	beatUrl = url
	// 第一个参数为客户的发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/", heartBeat)

	// 服务器要监听的主机地址和端口
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("ListenAndServer error: ", err.Error())
	}
}

type HeartJson struct {
	Action    string `json:"action"`
	ProcessId string `json:"s_process_id"`
}

func heartBeat(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Heart Beat!")
	HttpGetCfg(beatUrl) // 心跳

	//"http://192.168.208.34:9970/?action=2&s_process_id=%s", "ch_server"
}
