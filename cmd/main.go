package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"ttc/lib"
)

func GetTotalTimeInLine(line string) lib.PunchInTime {
	output := lib.NewPunchInTime()

	lineData := strings.Split(string(line), ",")
	date := lineData[0]
	numData := len(lineData)
	isWorking := false
	fromTime := lib.NewPunchInTime()
	toTime := lib.NewPunchInTime()

	for i := 1; i < numData; i++ {
		if isWorking {
			toTime = lib.NewPunchInTimeUsingTimeFormat(lineData[i])
			if lib.IsOverLunchTime(&fromTime, &toTime) == true {
				duration := toTime.Sub(&fromTime)
				lunchDuration := lib.NewPunchInTimeUsingTimeFormat("0:45")
				duration = duration.Sub(&lunchDuration)
				output = output.Add(&duration)
			} else if fromTime.IsInLunchTime() == false && toTime.IsInLunchTime() == true {
				toTime = lib.NewPunchInTimeUsingTimeFormat("12:00")
				duration := toTime.Sub(&fromTime)
				output = output.Add(&duration)
			} else if fromTime.IsInLunchTime() == true && toTime.IsInLunchTime() == false {
				fromTime = lib.NewPunchInTimeUsingTimeFormat("12:45")
				duration := toTime.Sub(&fromTime)
				output = output.Add(&duration)
			} else {
				duration := toTime.Sub(&fromTime)
				output = output.Add(&duration)
			}
			isWorking = false
		} else {
			fromTime = lib.NewPunchInTimeUsingTimeFormat(lineData[i])
			isWorking = true
		}
	}

	fmt.Printf("%v  %02d:%02d\n", date, output.Hour, output.Minute)

	return output
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic("failed to open file.")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	totalTime := lib.NewPunchInTime()
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("failed to read line.")
		}
		// fmt.Print(string(line))
		if !isPrefix {
			// fmt.Println()
		}

		totalTimeInLine := GetTotalTimeInLine(string(line))
		totalTime = totalTime.Add(&totalTimeInLine)
	}

	fmt.Printf("total  %02d:%02d\n", totalTime.Hour, totalTime.Minute)
}
