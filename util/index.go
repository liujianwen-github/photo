package util

import "fmt"

func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}