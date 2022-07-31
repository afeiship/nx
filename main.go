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
	fmt.Println("hello world")
}

func TypeOf(target interface{}) reflect.Kind {
	return reflect.ValueOf(target).Kind()
}
