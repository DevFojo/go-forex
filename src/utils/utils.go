package utils

import "fmt"

func ProcessError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
