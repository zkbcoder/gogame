package util

// 随机数
type Rand struct {
}

//  要求素数
var z1 uint32
var z2 uint32
var z3 uint32
var z4 uint32

func (this *Rand) Init() {
	z1 = 5
	z2 = 13
	z3 = 29
	z4 = 137
}

// 设置随机种子[不能随便填，必需是上一次的4个值保存数据库，重新赋值]
func (this *Rand) SetSeed(v1 uint32, v2 uint32, v3 uint32, v4 uint32) {
	z1 = v1
	z2 = v2
	z3 = v3
	z4 = v4
}

func (this *Rand) lfsr113() uint32 {
	var b uint32
	/* 产生随机的32位数字*/
	b = (((z1 << 6) ^ z1) >> 13)
	z1 = (((z1 & 4294967294) << 18) ^ b)

	b = (((z2 << 2) ^ z2) >> 27)
	z2 = (((z2 & 4294967288) << 2) ^ b)

	b = (((z3 << 13) ^ z3) >> 21)
	z3 = (((z3 & 4294967280) << 7) ^ b)

	b = (((z4 << 3) ^ z4) >> 12)
	z4 = (((z4 & 4294967168) << 13) ^ b)

	return (z1 ^ z2 ^ z3 ^ z4)
}

func (this *Rand) GetRand(uMax uint32) uint32 {
	if 0 == uMax {
		return 0
	}

	return this.lfsr113() % uMax
}

func (this *Rand) GetSeed() (v1, v2, v3, v4 uint32) {
	v1 = z1
	v2 = z2
	v3 = z3
	v4 = z4
	return
}

var MyRand Rand
