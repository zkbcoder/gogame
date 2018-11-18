package util

import (
	"fmt"
	"testing"
)

func TestRandPro(t *testing.T) {
	MyRand.Init()
	MyRand.GetRand(23)
	pro := new(RandPro)
	pro.AddData(100, 1)
	pro.AddData(200, 2)
	pro.AddData(300, 3)
	pro.AddData(500, 4)
	pro.AddData(1000, 5)

	res := pro.RandGet()

	ires, _ := res.(int)
	fmt.Println("res", ires)

}
