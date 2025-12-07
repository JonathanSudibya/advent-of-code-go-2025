package lib

import (
	"errors"
	"fmt"
	"time"
)

func CalculateAndPrintResultExecution(name string, input []byte, execute func([]byte) (uint64, bool)) {
	timeStart := time.Now()
	result, ok := execute(input)
	durationOfExecution := time.Since(timeStart)
	if !ok {
		fmt.Printf("result[%s] :\n", name)
	} else {
		fmt.Printf("result[%s] (%v) : %v\n", name, durationOfExecution, result)
	}
}

var ErrInvalidDigit = errors.New("invalid digit")

func BytesToInt(b []byte) (int, error) {
	n := 0
	for _, c := range b {
		if c < '0' || c > '9' {
			return 0, ErrInvalidDigit
		}
		n = n*10 + int(c-'0')
	}
	return n, nil
}
