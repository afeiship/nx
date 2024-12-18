package nx

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/template"
)

func TypeOf(target any) reflect.Kind {
	return reflect.TypeOf(target).Kind()
}

func Get(target, key string) any {
	return reflect.ValueOf(target).FieldByName(key).Interface()
}

func Set(target, key string, value any) {
	reflect.ValueOf(target).FieldByName(key).Set(reflect.ValueOf(value))
}

func If(cond bool, a, b any) any {
	if cond {
		return a
	}
	return b
}

func ToString(target any) string {
	return fmt.Sprintf("%v", target)
}

func ToInt(target string) int64 {
	res, _ := strconv.ParseInt(target, 10, 32)
	return res
}

func Tmpl(tmplString string, data any) string {
	tmpl, err := template.New("_").Parse(tmplString)
	buffer := bytes.NewBufferString("")

	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	err = tmpl.Execute(buffer, data)
	if err != nil {
		fmt.Println("Error applying template:", err)
	}
	return buffer.String()
}
