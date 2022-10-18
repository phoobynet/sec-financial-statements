package utils

import (
	"log"
	"strconv"
)

func mustBool(v string) bool {
	b, err := strconv.ParseBool(v)

	if err != nil {
		log.Fatalln(err)
	}

	return b
}
