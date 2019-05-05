package util

import (
	"fmt"
	"testing"
	"github.com/bmizerany/assert"
)

func TestBigNum(t *testing.T) {
	fmt.Println(90/1000.0)
	//m := make(map[int]map[int]string)
	//fmt.Println(m)
	//m[1] = make(map[int]string)
	//ok, vv := m[12]
	//fmt.Println(ok, vv)
	//fmt.Println(m[1])
	//m[1][2] = "123"
	//fmt.Println(m)

	v := BigNum{}
	v.DecVal(0, 999)

	fmt.Println(v)
	v.AddVal(0, 999)
	fmt.Println(v)
	v.AddVal(0, 999)
	fmt.Println(v)
	v.AddVal(0, 999)
	v.AddVal(0, 999)
	v.AddVal(0, 999)
	fmt.Println(v)
	//v.AddVal(4, 6)
	fmt.Println(v)

	v.AddInt(123456789, 0)
	fmt.Println(v)
	v.DecInt(123456789)
	fmt.Println(v)
	//a := 99898
	//fmt.Println(a %1000 , a / 1000)
}

func TestNumTH(t *testing.T) {
	assert.Equal(t, NumTH(1000, 1), 1000)
	assert.Equal(t, NumTH(1000, 2), 1000000)
	assert.Equal(t, NumTH(1000, 3), 1000000000)
	assert.Equal(t, NumTH(2, 32), 4294967296)
}

func TestBigNum_Add(t *testing.T) {
	a := BigNum{}
	a.AddInt(25000, 0)
	b := Str2BigNum("6.7MM")
	a.Add(b)
	assert.Equal(t, a.String(), "6.700MM", "TestBigNum_Add fail")
}

func TestBigNum_Compare(t *testing.T) {
	a := Str2BigNum("6.9MM")
	b := Str2BigNum("6.7MM")
	assert.Equal(t,  a.Equal(b), false,"TestBigNum_Equal fail")
	assert.Equal(t,  a.Greater(b), true,"TestBigNum_Greater fail")
	assert.Equal(t,  a.Less(b), false,"TestBigNum_Less fail")

	c := Str2BigNum("6.9MM")
	d := Str2BigNum("6.7MMM")
	assert.Equal(t,  c.Equal(d), false,"TestBigNum_Equal fail")
	assert.Equal(t,  c.Greater(d), false,"TestBigNum_Greater fail")
	assert.Equal(t,  c.Less(d), true,"TestBigNum_Less fail")
	assert.Equal(t,  c.Equal(a), true,"TestBigNum_Equal fail")
}

func TestBigNum_Mul(t *testing.T) {
	a := Str2BigNum("6.9MM")
	b := Str2BigNum("20.7NN")
	new := a.Mul(3000)
	fmt.Println(new)
	fmt.Println(a)
	fmt.Println(b)
	assert.Equal(t,  new.Equal(b), true,"TestBigNum_Equal fail")
}

func TestBigNum_Dec(t *testing.T) {
	a := Str2BigNum("6.1MM")
	b := Str2BigNum("2.2MM")
	a.Dec(b)
	assert.Equal(t,  a.String(), "3.900MM","TestBigNum_Equal fail")
}

func TestBigNum_AddInt(t *testing.T) {
	v := BigNum{}
	v.AddInt(10000, 0)
	assert.Equal(t,  v.String(), "10K","TestBigNum_AddInt fail")
}
