package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 加载json文件
func LoadJson(filename string, v interface{}) bool {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error!->", filename)
		return false
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("Error! decode json->", filename, err)
		return false
	}
	return true
}
