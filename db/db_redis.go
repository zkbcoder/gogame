package db

import (
	"fmt"
	"github.com/zkbcoder/gogame/util"
	"time"

	"github.com/garyburd/redigo/redis"
)

type DBRedis struct {
	IP   string     // redisIP
	DBID int        // 连接哪个库
	conn redis.Conn // redis连接
}

// 连接redis
func (this *DBRedis) Dial() {
	c, err := redis.Dial("tcp", this.IP)
	if err != nil {
		util.ErrRecord("redis dial err [%s]", err)
		this.conn.Close()
	} else {
		fmt.Println("connect")
		this.conn = c
		c.Do("SELECT", this.DBID)
	}
}

func (this *DBRedis) Close() {
	this.conn.Close()
}

func (this DBRedis) NewDail() (redis.Conn, error) {
	c, err := redis.Dial("tcp", this.IP)
	if err != nil {
		util.ErrRecord("redis dial err [%s]", err)
	} else {
		c.Do("SELECT", this.DBID)
	}
	return c, err
}

// Delete 不直接提供链接接口 通过分装的Do调用,加上了重连机制
//func (this *DBRedis) GetConn() redis.Conn {
//	return this.conn
//}

// 断线重连
func (this *DBRedis) Reconnect() {
	time.Sleep(time.Second * 1)
	this.Dial()
	if this.conn == nil { // 如果一直连接不上就继续重试
		this.Reconnect()
	}
}

// 检查重连
func (this *DBRedis) CheckConnect() {
	if this.conn == nil {
		util.ErrRecord("Not init redis Reconnect")
		this.Reconnect()
	}

	if this.conn.Err() != nil {
		util.ErrRecord("redis Reconnect %s", this.conn.Err())
		this.conn = nil
		this.Reconnect()
	}
}

// 不做直接访问conn，并且加上重连
func (this *DBRedis) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	this.CheckConnect() //检查链接是否正常
	return this.conn.Do(commandName, args...)
}
