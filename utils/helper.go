package utils

import "fmt"

func CheckErr(err error) bool {
	if err != nil {
		fmt.Print(err)
		return true
	}
	return false
}
