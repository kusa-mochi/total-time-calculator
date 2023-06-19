package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"ttc/lib"
)

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic("failed to open file.")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic("failed to read line.")
		}
		fmt.Print(string(line))
		if !isPrefix {
			fmt.Println()
		}

		t1 := lib.NewPunchInTimeUsingTimeParams(8, 20)
		t2 := lib.NewPunchInTimeUsingTimeParams(16, 50)
		s := t2.Sub(&t1)
		s.Print()
	}
}
