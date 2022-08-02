package nx

import (
	"fmt"
	"os"
	"reflect"
)

func Hi() {
	fmt.Println("hello world")
}

// Return the type of the target
func TypeOf(target any) reflect.Kind {
	return reflect.TypeOf(target).Kind()
}

// Getenv returns the value of the environment variable named by the key.
func Getenv(name string) string {
	return os.Getenv(name)
}
