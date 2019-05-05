// 同步计数
package util

type SynCounter struct {
	Ch    chan int
	count int
}

func (this *SynCounter) Add() {
	this.Ch <- 1
}

func (this *SynCounter) Dec() {
	this.Ch <- -1
}

func (this *SynCounter) Update() {
	for {
		res := <-this.Ch
		this.count += res
		ErrRecord("Count = [%d]", this.count)
		if this.count <= 0 {
			break
		}
	}
}
