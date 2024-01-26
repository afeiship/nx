package nx

import (
	"github.com/afeiship/nx/lib"
	"reflect"
	"testing"
)

func TestJSONParse(t *testing.T) {
	// prepare data
	jsonString := `{"a":1}`

	// start to write a test
	res, err := nx.JSONParse(jsonString)

	// // check result
	if err != nil {
		t.Error(err)
	}

	resType := nx.TypeOf(res)
	if resType != reflect.Map {
		t.Error("type of result is not map")
	}
}

func TestJSONStringify(t *testing.T) {
	// prepare data
	jsonObject := map[string]interface{}{
		"a": 1,
	}

	// start to write a test
	res, err := nx.JSONStringify(jsonObject)

	// // check result
	if err != nil {
		t.Error(err)
	}

	if res != `{"a":1}` {
		t.Error("stringify error")
	}
}
