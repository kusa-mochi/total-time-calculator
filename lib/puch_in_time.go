package lib

import "fmt"

type PunchInTime struct {
	Hour   int
	Minute int
}

func NewPunchInTime() PunchInTime {
	return PunchInTime{
		Hour:   0,
		Minute: 0,
	}
}

func NewPunchInTimeUsingTimeParams(h int, m int) PunchInTime {
	return PunchInTime{
		Hour:   h,
		Minute: m,
	}
}

// t + tm
func (t *PunchInTime) Add(tm *PunchInTime) PunchInTime {
	var output PunchInTime = NewPunchInTime()
	output.Hour = t.Hour + tm.Hour
	output.Minute = t.Minute + tm.Minute

	for output.Minute > 60 {
		output.Hour++
		output.Minute -= 60
	}

	return output
}

// t - tm
func (t *PunchInTime) Sub(tm *PunchInTime) PunchInTime {
	var output PunchInTime = NewPunchInTime()
	output.Hour = t.Hour - tm.Hour
	output.Minute = t.Minute - tm.Minute

	for output.Minute < 0 {
		output.Hour--
		output.Minute += 60
	}

	return output
}

func (t *PunchInTime) Print() {
	h := t.Hour
	m := t.Minute
	fmt.Printf("%v:%v\n", h, m)
}
