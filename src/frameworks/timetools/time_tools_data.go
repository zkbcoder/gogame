package timetools

type ITimeCallBack interface {
	OnTimer(nIntervalTime int64)
}

type TimeToolData struct {
	CallbackList map[ITimeCallBack]bool //先用list go 没有 set 用map代替又很难看 重复那边有业务去做吧

	NextTime int64
}

func CreateTimeToolData() *TimeToolData {
	return &TimeToolData{
		CallbackList: make(map[ITimeCallBack]bool),
	}
}

func (this *TimeToolData) SetNextTime(nextTime int64) {
	this.NextTime = nextTime
}

func (this *TimeToolData) GetNextTime() int64 {
	return this.NextTime
}
