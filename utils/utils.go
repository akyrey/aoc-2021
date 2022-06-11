package utils

import (
	"fmt"
	"os"
)

func Contains[T int | string](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFileToReadFrom(day int, test bool) (*os.File, error) {
	dayStr := getDayString(day)
	if test {
		return os.Open(fmt.Sprintf("day%s/test%s.txt", dayStr, dayStr))
	}
	return os.Open(fmt.Sprintf("day%s/input%s.txt", dayStr, dayStr))
}

func getDayString(day int) string {
	if day < 10 {
		return fmt.Sprintf("0%d", day)
	}

	return fmt.Sprintf("%d", day)
}
