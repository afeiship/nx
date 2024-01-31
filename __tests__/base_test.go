package nx

import (
	"fmt"
	"github.com/afeiship/nx/lib"
	"reflect"
	"testing"
)

func TestTypeOf(f *testing.T) {
	// prepare data
	v1 := "string"
	v2 := 1
	v3 := 1.0
	v4 := true
	v5 := []string{"a", "b", "c"}
	v6 := map[string]string{"a": "b", "c": "d"}
	v7 := struct {
		A string
		B int
	}{"a", 1}

	// start to write a test
	res1 := nx.TypeOf(v1)
	res2 := nx.TypeOf(v2)
	res3 := nx.TypeOf(v3)
	res4 := nx.TypeOf(v4)
	res5 := nx.TypeOf(v5)
	res6 := nx.TypeOf(v6)
	res7 := nx.TypeOf(v7)

	// check result
	if res1 != reflect.String {
		f.Errorf("type of string is not string")
	}

	if res2 != reflect.Int {
		f.Errorf("type of int is not int")
	}

	if res3 != reflect.Float64 {
		f.Errorf("type of float64 is not float64")
	}

	if res4 != reflect.Bool {
		f.Errorf("type of bool is not bool")
	}

	if res5 != reflect.Slice {
		f.Errorf("type of slice is not slice")
	}

	if res6 != reflect.Map {
		f.Errorf("type of map is not map")
	}

	if res7 != reflect.Struct {
		f.Errorf("type of struct is not struct")
	}
}

func TestGetenv(t *testing.T) {
	homedir := nx.Getenv("HOME")
	gdir := nx.Getenv("GITHUB_HOME")

	if gdir != homedir+"/github" {
		t.Errorf("getenv error")
	}
}

func TestToString(t *testing.T) {
	v1 := 1
	v2 := true

	if nx.ToString(v1) != "1" {
		t.Errorf("to string error")
	}

	if nx.ToString(v2) != "true" {
		t.Errorf("to string error")
	}
}

func TestToInt(t *testing.T) {
	v1 := "1.0"
	v2 := "1"
	v3 := "123a"

	if nx.ToInt(v1) != 0 {
		t.Errorf("to int error")
	}

	if nx.ToInt(v2) != 1 {
		t.Errorf("to int error")
	}

	if nx.ToInt(v3) != 0 {
		t.Errorf("to int error")
	}
}

// only
func TestTmpl(t *testing.T) {
	t1 := "hello {{.Name}}"
	data := map[string]string{"Name": "world"}
	res := nx.Tmpl(t1, data)
	// test res equal
	if res != "hello world" {
		t.Errorf("tmpl error, got: %s", res)
	}
	fmt.Println("tmpl test done:", res)
}
