package nx

import (
	"fmt"
	"reflect"
)

func Log(args ...interface{}) {
	for _, param := range args {
		fmt.Printf("%+v ", param)
	}
	fmt.Println("")
}

func Hello() {
	fmt.Printf("hello world\n")
}

func TypeOf(target interface{}) interface{} {
	return reflect.ValueOf(target).Kind()
}
