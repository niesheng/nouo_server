package main

import (
	"fmt"
	"os"
	"reflect"
)

func Exit(err interface{}) {
	fmt.Println(err)
	os.Exit(0)
}

func fmtParams(s map[string]interface{}) func(k, v []byte) {
	return func(k, v []byte) {
		if p, _ := s[string(k)]; p == nil {
			s[string(k)] = string(v)
		} else {
			var vv []string
			if reflect.TypeOf(p).String() == "string" {
				vv = []string{p.(string), string(v)}
			} else {
				vv = append(p.([]string), string(v))

			}
			s[string(k)] = vv
		}
	}
}
