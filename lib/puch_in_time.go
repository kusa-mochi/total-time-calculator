package lib

import (
	"fmt"
	"strconv"
	"strings"
)

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

func NewPunchInTimeUsingTimeFormat(s string) PunchInTime {
	timeValues := strings.Split(s, ":")
	h, _ := strconv.Atoi(timeValues[0])
	m, _ := strconv.Atoi(timeValues[1])
	return NewPunchInTimeUsingTimeParams(h, m)
}

func IsOverLunchTime(startTime *PunchInTime, endTime *PunchInTime) bool {
	isBeforeNeen := startTime.Hour < 12
	isAfter13 := endTime.Hour > 12
	isLunchTimeIn12 := endTime.Hour == 12 && endTime.Minute >= 45
	return (isBeforeNeen && (isAfter13 || isLunchTimeIn12))
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

func (t *PunchInTime) IsInLunchTime() bool {
	return (t.Hour == 12 && 0 <= t.Minute && t.Minute <= 45)
}
