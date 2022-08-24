package db

import (
	"strings"
)

func pgArrayToSlice(pgArr string) []string {
	chars := []string{"{", "}"}
	for _, v := range chars {
		pgArr = strings.ReplaceAll(pgArr, v, "")
	}
	slice := strings.Split(pgArr, ",")

	return slice
}
