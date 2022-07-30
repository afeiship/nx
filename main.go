package nx

import "fmt"

var VERSION = "0.0.1"

func Log(args ...interface{}) {
  for _, param := range args {
    fmt.Printf("%+v ", param)
  }
  fmt.Println("")
}

func Hi() {
  fmt.Printf("Hello go cli")
}
