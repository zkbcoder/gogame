// 时间操作
package util

type TimeQueueEvent interface {
	TimeOver()
}

// 时间队列
type TimeQueue struct {
	mapQueue map[int64]interface{}
}

func (this *TimeQueue) Register(monitor interface{}) {

}
