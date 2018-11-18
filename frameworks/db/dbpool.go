package db

import (
	"fmt"
	"frameworks/util"
	"sync"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

type ConnInfo struct {
	Conn     *mysql.Conn
	LastTime int64
	IsUsed   bool
}

type DBPool struct {
	Pool   map[*mysql.Conn]*ConnInfo
	Raddr  string
	DB     string
	User   string
	Passwd string
	Db     string
	Mutex  *sync.Mutex

	KeepTime int64
}

func NewPool(raddr, user, passwd, db string, keepTime int64) *DBPool {

	return &DBPool{
		Raddr:    raddr,
		DB:       db,
		User:     user,
		Passwd:   passwd,
		Mutex:    new(sync.Mutex),
		Pool:     make(map[*mysql.Conn]*ConnInfo),
		KeepTime: keepTime,
	}
}

func (this *DBPool) Empty() bool {
	return len(this.Pool) == 0
}

func (this *DBPool) Size() int {
	return len(this.Pool)
}

func (this *DBPool) Grab() *mysql.Conn {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	this.remove_old_connections()
	conn := this.find_mru()
	if nil == conn {
		db := mysql.New("tcp", "", this.Raddr, this.User, this.Passwd, this.DB)
		conn = &db
		err := (*conn).Connect()
		if err != nil {
			fmt.Println(err)
		} else {
			conn = &db
			connInfo := &ConnInfo{
				Conn:     conn,
				IsUsed:   true,
				LastTime: util.GetTimeUnix(),
			}
			this.Pool[conn] = connInfo
		}
	}

	return conn
}

func (this *DBPool) Release(conn *mysql.Conn) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	this.Pool[conn].IsUsed = false
	this.Pool[conn].LastTime = util.GetTimeUnix()
}

func (this *DBPool) Clear() {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	for i, _ := range this.Pool {
		(*i).Close()
	}
	this.Pool = make(map[*mysql.Conn]*ConnInfo)
}

func (this *DBPool) find_mru() *mysql.Conn {
	var connInfo *ConnInfo
	connInfo = nil
	for _, v := range this.Pool {
		if !v.IsUsed {
			if connInfo == nil || connInfo.LastTime > v.LastTime {
				connInfo = v
			}
		}
	}

	if nil == connInfo {
		return nil
	}
	connInfo.IsUsed = true
	return connInfo.Conn
}

func (this *DBPool) remove_old_connections() {
	t := util.GetTimeUnix()
	delConn := make([]*mysql.Conn, 0)

	for i, v := range this.Pool {
		if !v.IsUsed && t-v.LastTime > this.KeepTime {
			delConn = append(delConn, i)
		}
	}

	for _, conn := range delConn {
		if nil != conn {
			(*conn).Close()
			delete(this.Pool, conn)
		}
	}
}
