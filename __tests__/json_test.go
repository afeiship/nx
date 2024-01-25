package nx

import (
	"github.com/afeiship/nx"
	"testing"
	"fmt"
)

// define Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJSON(t *testing.T) {
	// test json parse
	var person = Person{}
	var data = []byte(`{"name":"afei","age":18}`)
	var res = nx.JsonParse(data, &person)
	if res != nil {
		t.Errorf("json parse failed")
	}

	fmt.Println(person)

	// test json stringify
	// var res2 = nx.JsonStringify(person)
	// if res2 != `{"name":"afei","age":18}` {
	// 	t.Errorf("json stringify failed")
	// }
}
