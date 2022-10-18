package utils

import (
	"reflect"
	"strings"
)

func ProcessLine[T any](line string) *T {
	tokens := strings.Split(line, "\t")
	t := reflect.TypeOf(*new(T))
	v := reflect.New(t)
	vElem := v.Elem()
	for i := 0; i < t.NumField(); i++ {
		vElem.Field(i).SetString(tokens[i])
	}

	return v.Interface().(*T)
}
