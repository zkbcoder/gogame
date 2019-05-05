package timetools

import (
	"fmt"
	"github.com/zkbcoder/gogame/util"
	"testing"
	"time"
)

type timetoolstest struct {
	EType int
}

func (this *timetoolstest) OnTimer(nIntervalTime int64) {
	fmt.Println("OnTimer()---", this.EType, "now= ", time.Now())
}

func TestTimeTools(t *testing.T) {
	alarm := timetoolstest{
		EType: ETimeType_Alarm,
	}
	clock := timetoolstest{
		EType: ETimeType_Clock,
	}
	moment := timetoolstest{
		EType: ETimeType_Moment,
	}
	weekMoment := timetoolstest{
		EType: ETimeType_WeekMoment,
	}
	CTimeToolsIst.Init()
	CTimeToolsIst.RegisterTimer(ETimeType_Alarm, 60, &alarm)
	CTimeToolsIst.RegisterTimer(ETimeType_Clock, 5, &clock)
	CTimeToolsIst.RegisterTimer(ETimeType_Moment, util.GetTodayPassTime()+10, &moment)
	CTimeToolsIst.RegisterTimer(ETimeType_WeekMoment, util.GetTimeUnix()-util.GetWeekBeginTime()+20, &weekMoment)

	fmt.Println("weektime", util.GetTimeUnix()-util.GetWeekBeginTime()+20)

	for {

		CTimeToolsIst.OnUpdate()
		time.Sleep(time.Second)
	}
}
