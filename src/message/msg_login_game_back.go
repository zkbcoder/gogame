package message

import (
	"bytes"
	"network"
)

type RoleInfo struct {
	Id         int32  //用户ID
	Lv         int32  //保留
	WeaponType int32  //保留
	Weaponimg  int32  //保留
	Scapulaimg int32  //保留
	Armorimg   int32  //保留
	Pantsimg   int32  //保留
	Hair       int32  //保留
	Porf       int32  //保留
	Name       string //保留
	index      int32  //保留
}

// 请求服务列表
type MsgLoginGameBack struct {
	Head  network.MsgHead
	Count int32
	Role  []RoleInfo
}

/////////////////////////////////////////////////////////////////////////////
// IMsg
func (msg *MsgLoginGameBack) Packet() *bytes.Buffer {
	return nil
}

func (msg *MsgLoginGameBack) UnPacket(buf []byte) {
	network.UnPacket(buf, &msg.Head.MsgLen)
	network.UnPacket(buf[2:], &msg.Head.MsgId)
	nPos := 4
	network.UnPacket(buf[nPos:], &msg.Count)
	nPos += 4
	var i int32
	for ; i < msg.Count; i++ {
		var role RoleInfo
		network.UnPacket(buf[nPos:], &role.Id)
		nPos += 4
		network.UnPacket(buf[nPos:], &role.Lv)
		nPos += 4
		network.UnPacket(buf[nPos:], &role.WeaponType)
		nPos += 4
		network.UnPacket(buf[nPos:], &role.Weaponimg)
		nPos += 4
		msg.Role = append(msg.Role, role)
	}
}
