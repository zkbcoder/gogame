package util

import (
	"fmt"
)

// 相邻单位之间间隔10^3
// 显示数值分为：数值部分与单位部分，例如10k，100m，30bbb
// 如果加成后数值部分大于等于1000，则数值部分除以1000，单位部分变为下一个级别的单位
// 数值部分仅会显示3位有效数值
// 不做负数 a < b 那么 a - b = 0

var ABBR = []string{
	"K",
	"M",
	"B",
	"T",
	"AA",
	"BB",
	"CC",
	"DD",
	"EE",
	"FF",
	"GG",
	"HH",
	"II",
	"JJ",
	"KK",
	"LL",
	"MM",
	"NN",
	"OO",
	"PP",
	"QQ",
	"RR",
	"SS",
	"TT",
	"UU",
	"VV",
	"WW",
	"XX",
	"YY",
	"ZZ",
	"AAA",
	"BBB",
	"CCC",
	"DDD",
	"EEE",
	"FFF",
	"GGG",
	"HHH",
	"III",
	"JJJ",
	"KKK",
	"LLL",
	"MMM",
	"NNN",
	"OOO",
	"PPP",
	"QQQ",
	"RRR",
	"SSS",
	"TTT",
	"UUU",
	"VVV",
	"WWW",
	"XXX",
	"YYY",
	"ZZZ",
	"E",
}

type BigNum struct {
	Bits []int16
}

// 求num的th次方
func NumTH(num int, th uint) int {
	ret := 1
	var i uint
	for i = 0; i < th; i++ {
		ret = ret * num
	}
	return ret
}

// 加一个值<1000的,指定位[如果超过高位，低位补0]
func (this *BigNum) AddVal(post int, val int16) {
	if val >= 1000 || post < 0 {
		fmt.Println("E: Addval->", post, val)
		return
	}

	// 高位超过了，补上空位
	if post >= len(this.Bits) {
		zeroPos := post - len(this.Bits)
		for i := 0; i < zeroPos; i++ {
			this.Bits = append(this.Bits, 0) // 直接赋值高位
		}
		this.Bits = append(this.Bits, val) // 直接赋值高位
		return
	}

	this.Bits[post] += val
	if this.Bits[post] >= 1000 {
		this.AddVal(post+1, 1) // 进位
	}
	this.Bits[post] %= 1000
}

// 减一个值<1000的,指定位
func (this *BigNum) DecVal(post int, val int16) bool {
	// 减法不能超过最高位
	if post >= len(this.Bits) || val >= 1000 {
		fmt.Println("E: DecVal->", post, val)
		return false
	}

	// 当前值不够向高位借1
	if this.Bits[post] < val {
		if !this.DecVal(post+1, 1) {
			return false // 没有高位了
		}
		this.Bits[post] = this.Bits[post] + 1000 - val
	} else {
		this.Bits[post] = this.Bits[post] - val
	}

	// 如果高位变成0了，去掉
	higtPos := len(this.Bits) - 1
	if this.Bits[post] == 0 && post == higtPos {
		this.Bits = this.Bits[:higtPos]
	}
	return true
}

// 加上一个int值
func (this *BigNum) AddInt(left int, index int) {
	if left < 0 {
		return
	}

	l := 0 // 从个位开始
	for {
		// 从低位开始截取
		th := NumTH(10, uint(l*3))
		if left >= th {
			bit := left / (1 * th) % 1000
			this.AddVal(index+l, int16(bit))
			//fmt.Println(index+l, int16(bit))
		} else {
			break
		}
		l++
	}
}

func (this *BigNum) DecInt(left int) {
	if left < 0 {
		return
	}

	l := 0 // 从个位开始
	for {
		// 从低位开始截取
		th := NumTH(10, uint(l*3))
		if left > th {
			bit := left / (1 * th) % 1000
			this.DecVal(l, int16(bit))
		} else {
			break
		}
		l++
	}
}

func (this *BigNum) Add(left *BigNum) {
	l := left.Len()
	for i:=0; i<l; i++ {
		this.AddVal(i, left.Bits[i])
	}
}

func (this *BigNum) Dec(left *BigNum) {
	if this.Compare(left) < 0 {
		// 负数不计算
		return
	}

	l := left.Len()
	for i:=0; i<l; i++ {
		this.DecVal(i, left.Bits[i])
	}
}

// 乘法
func (this BigNum) Mul(left int) *BigNum {
	newNum := BigNum{}

	l := this.Len()
	for i:=0; i<l; i++ {
		tmp := int(this.Bits[i]) * left
		if tmp > 1000 {
			newNum.AddInt(tmp, i)
		} else {
			newNum.AddVal(i, int16(tmp))
		}
	}
	return &newNum
}

func GetPost(str string) int {
	max := len(ABBR)
	post := 0
	for i := 0; i < max; i++ {
		if str == ABBR[i] {
			return i + 1
		}
	}
	return post
}

func (this BigNum) Len() int {
	return len(this.Bits)
}

// 比较两个数，返回<0, 0, >0
func (this BigNum) Compare(left *BigNum) int {
	if this.Len() != left.Len() {
		return this.Len() - left.Len()
	}

	l := this.Len()

	for i:=l-1; i>=0; i-- {
		if this.Bits[i] != left.Bits[i] {
			return int(this.Bits[i] - left.Bits[i])
		}
	}

	return 0
}

func (this BigNum) Equal(left *BigNum) bool {
	return this.Compare(left) == 0
}

func (this BigNum) Greater (left *BigNum) bool {
	return this.Compare(left) > 0
}

func (this BigNum) Less (left *BigNum) bool {
	return this.Compare(left) < 0
}

func (this BigNum) String() string {
	l := len(this.Bits)
	if l == 0 {
		return "0"
	}

	if l == 1 {
		return fmt.Sprintf("%d", this.Bits[0])
	}

	if this.Bits[l-2] == 0 {
		return fmt.Sprintf("%d%s", this.Bits[l-1],  ABBR[l-2])
	}
	return fmt.Sprintf("%d.%03d%s", this.Bits[l-1], this.Bits[l-2], ABBR[l-2])
}
