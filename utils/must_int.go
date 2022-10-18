package utils

import (
	"log"
	"strconv"
)

func mustInt(v string) int {
	i, err := strconv.Atoi(v)

	if err != nil {
		log.Fatalln(err)
	}

	return i
}
