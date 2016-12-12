package timetools

import (
	"frameworks/util"
	"time"
)

const (
	ETimeType_Alarm      = 1
	ETimeType_Clock      = 2
	ETimeType_Moment     = 3
	ETimeType_WeekMoment = 4
)

type CTimeTools struct {
	AlarmTime      map[int64]*TimeToolData // 一天内的 整点触发
	ClockTime      map[int64]*TimeToolData //定时触发
	MomentTime     map[int64]*TimeToolData //每天的 某个点触发
	WeekMomentTime map[int64]*TimeToolData //每周的 某个点触发
}

func (this *CTimeTools) Init() {
	this.AlarmTime = make(map[int64]*TimeToolData)
	this.ClockTime = make(map[int64]*TimeToolData)
	this.MomentTime = make(map[int64]*TimeToolData)
	this.WeekMomentTime = make(map[int64]*TimeToolData)
}

func (this *CTimeTools) OnUpdate() {
	this.OnAlarmUpdate()
	this.OnClockUpdate()
	this.OnMomentUpdate()
	this.OnWeekMomentUpdate()
}

func (this *CTimeTools) OnAlarmUpdate() {
	tCruTime := time.Now().Unix()
	for i, v := range this.AlarmTime {
		if 0 == v.NextTime {
			v.SetNextTime(tCruTime + i - tCruTime%i)
		}
		if tCruTime >= v.NextTime {
			this.DispatchOnTimer(i, v)
			v.SetNextTime(tCruTime + i - tCruTime%i)
		}
	}
}

func (this *CTimeTools) OnClockUpdate() {
	tCruTime := time.Now().Unix()
	for i, v := range this.ClockTime {
		if 0 == v.NextTime {
			v.SetNextTime(tCruTime + i)
		}
		if tCruTime >= v.NextTime {
			this.DispatchOnTimer(i, v)
			v.SetNextTime(tCruTime + i)
		}
	}
}

func (this *CTimeTools) OnMomentUpdate() {
	tCruTime := time.Now().Unix()
	for i, v := range this.MomentTime {
		if 0 == v.NextTime {
			v.SetNextTime(util.GetNextTimeBySec(i))
		}
		if tCruTime >= v.NextTime {
			this.DispatchOnTimer(i, v)
			v.SetNextTime(util.GetNextTimeBySec(i))
		}
	}
}

func (this *CTimeTools) OnWeekMomentUpdate() {
	tCruTime := time.Now().Unix()
	for i, v := range this.WeekMomentTime {
		if 0 == v.NextTime {
			v.SetNextTime(util.GetNextWeekTimeBySec(i))
		}
		if tCruTime >= v.NextTime {
			this.DispatchOnTimer(i, v)
			v.SetNextTime(util.GetNextWeekTimeBySec(i))
		}
	}
}

func (this *CTimeTools) RegisterTimer(nType int, nIntervalTime int64, callback ITimeCallBack) {
	switch nType {
	case ETimeType_Alarm:
		this.registerTimer(&this.AlarmTime, nIntervalTime, callback)
	case ETimeType_Clock:
		this.registerTimer(&this.ClockTime, nIntervalTime, callback)
	case ETimeType_Moment:
		this.registerTimer(&this.MomentTime, nIntervalTime, callback)
	case ETimeType_WeekMoment:
		this.registerTimer(&this.WeekMomentTime, nIntervalTime, callback)
	}
}

func (this *CTimeTools) registerTimer(timer *map[int64]*TimeToolData, nIntervalTime int64, callback ITimeCallBack) {
	_, ok := (*timer)[nIntervalTime]
	if !ok {
		(*timer)[nIntervalTime] = CreateTimeToolData()
	}
	(*timer)[nIntervalTime].CallbackList[callback] = true
}

func (this *CTimeTools) UnRegisterTimer(nType int, nIntervalTime int64, callback ITimeCallBack) {
	switch nType {
	case ETimeType_Alarm:
		this.unRegisterTimer(&this.AlarmTime, nIntervalTime, callback)
	case ETimeType_Clock:
		this.unRegisterTimer(&this.ClockTime, nIntervalTime, callback)
	case ETimeType_Moment:
		this.unRegisterTimer(&this.MomentTime, nIntervalTime, callback)
	case ETimeType_WeekMoment:
		this.unRegisterTimer(&this.WeekMomentTime, nIntervalTime, callback)
	}
}

func (this *CTimeTools) unRegisterTimer(timer *map[int64]*TimeToolData, nIntervalTime int64, callback ITimeCallBack) {
	_, ok := (*timer)[nIntervalTime]
	if ok {
		delete((*timer)[nIntervalTime].CallbackList, callback)
	}

}

func (this *CTimeTools) DispatchOnTimer(nIntervalTime int64, timerSet *TimeToolData) {
	for iCallBack, _ := range timerSet.CallbackList {
		iCallBack.OnTimer(nIntervalTime)
	}
}

var CTimeToolsIst CTimeTools
