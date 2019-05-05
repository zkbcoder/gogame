package util

import (
	"testing"
	"fmt"
)

type ItemBase struct {
	Type int `json:"type"`
	Id int `json:"id"`
	Num string `json:"num"`
}

type Equip struct {
	Quality int `json:"quality"`
	Item ItemBase `json:"item"`
	Transform_item ItemBase `json:"transform_item"`
}

type EquipCfg struct {
	Equips []*Equip // 装备
}

type Bonus struct {
	Type int `json:"type"`
	Param1 string `json:"param_1"`
	Param2 string `json:"param_2"`
}

type EquipData struct {
	Id int `json:"id"`
	EquipId int `json:"equip_id"`
	Quality int `json:"quality"`
	AttrBonus []Bonus `json:"attr_bonus"`
}

func TestLoadJson(t *testing.T) {
	//var v EquipCfg
	m := make(map[string]*EquipData)
	LoadJson("test.json", &m)
	for k,v := range m {
		fmt.Println(k, v)
	}

	v, ok := m["1"]
	if ok {
		fmt.Println(v)
	}
}