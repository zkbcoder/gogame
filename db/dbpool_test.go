package db

import (
	//	"bytes"
	"fmt"
	//	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
)

//func TestRemove(t *testing.T) {
//	poolmgr := NewPool("192.168.99.3:3306", "public", "public", "ecaccount")
//	fmt.Println("befor size", poolmgr.Size())
//	conn := poolmgr.Grab()
//	_, _, err := (*conn).Query("select * from account limit 1")
//	if err != nil {
//		fmt.Println(err)
//	}
//	poolmgr.Release(conn)

//	fmt.Println("after size", poolmgr.Size())
//}
func TestStmt(t *testing.T) {
	poolmgr := NewPool("192.168.99.3:3306", "public", "public", "ecaccount", 200)
	conn := poolmgr.Grab()
	defer poolmgr.Release(conn)
	(*conn).SetMaxPktSize(4 * 1024 * 1024)
	ins, err := (*conn).Prepare("insert test_binary_file values (?,?)")
	sel, err := (*conn).Prepare("select data from test_binary_file where id = ?")
	if err != nil {
		fmt.Println("stmt", err)
	}
	var (
		//		rre native.RowsResErr
		id int64
	)
	sel.Bind(&id)

	//	res, err := sel.Run()
	//	if err != nil {
	//		fmt.Println("stmt", err)
	//	}
	//	row, err := res.GetRow()
	//	if err != nil {
	//		fmt.Println("stmt", err)
	//	}
	id = 1
	//bind 按同样参数类型顺序 传参数
	ins.Bind(&id, []byte(nil))

	file, err := os.Open("900#63331#1459244503#11")
	if err != nil {
		fmt.Println("stmt", err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		fmt.Println("stmt", err)
	}
	// 第一个参数代表sql语句的第几个参数需要发送 第二个参数可以是文件 也可以是［］bytes 第三个参数是长度
	ins.SendLongData(1, file, int(info.Size()))

	//	buf := make([]byte, 1024)
	//	ins.SendLongData(1, buf, 1024)

	res, err := ins.Run()
	fmt.Println("stmtres", res)
	if err != nil {
		fmt.Println("stmterr", err)
	}
	//	data, err := ioutil.ReadFile("900#63331#1459244503#11")
	//	if err != nil {
	//		fmt.Println("stmt", err)
	//	}
	//	fmt.Println("row ", row.Bin(0))
	//	fmt.Println("data ", data)
	//	if row == nil || row[0] == nil ||
	//		bytes.Compare(data, row.Bin(0)) != 0 {
	//		t.Fatal("Bad result")
	//	}
	//	ioutil.WriteFile("res", row.Bin(0), 0644)

}

func zzzTestPool(t *testing.T) {
	poolmgr := NewPool("192.168.99.3:3306", "public", "public", "ecaccount", 2)

	ch := make(chan []mysql.Row)

	runCount := 20

	for i := 0; i < runCount; i++ {
		go func() {
			conn := poolmgr.Grab()
			row, _, err := (*conn).Query("select * from account limit 1")
			if err != nil {
				fmt.Println("sql err ===== ", err)
			}
			ch <- row
			poolmgr.Release(conn)
		}()
	}

	count := 0
	for {
		row := <-ch
		fmt.Println(row)
		fmt.Println("poolmgr size = ", poolmgr.Size())
		count++
		if count == runCount {
			break
		}
	}

	//	for {
	//	select {
	//case
	<-time.After(time.Second * 1) //:
	//		goto exit
	//	}
	//}
	//exit:
	fmt.Println("second time---------------------------------")
	for i := 0; i < runCount; i++ {
		go func() {
			conn := poolmgr.Grab()
			fmt.Println("conn addr===", conn)
			row, _, err := (*conn).Query("select * from account limit 1")
			if err != nil {
				fmt.Println(err)
			}
			ch <- row
			poolmgr.Release(conn)

		}()
	}

	count = 0
	for {
		row := <-ch
		fmt.Println(row)
		fmt.Println("poolmgr size = ", poolmgr.Size())
		count++
		if count == runCount {
			break
		}
	}
}
