package utils

import (
	"errors"
	"fmt"
	"time"
)

const TimeLayout = "2006-01-02"

func ProcessError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func ExtractDate(path string) (time.Time, error) {
	if len(path) > 1 {
		return time.Parse(TimeLayout, path[1:])
	}
	return time.Time{}, errors.New("invalid string")
}
