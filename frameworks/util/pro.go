package util

type RandPro struct {
	randArr []int
	datas   []interface{}
	max_num int
}

func (this *RandPro) AddData(nWeith int, data interface{}) {
	this.max_num += nWeith
	this.randArr = append(this.randArr, this.max_num)
	this.datas = append(this.datas, data)
}

func (this *RandPro) RandGet() interface{} {
	randIndex := MyRand.GetRand(uint32(this.max_num))
	pos := My_lower_bound(this.randArr, int(randIndex))
	return this.datas[pos]
}
