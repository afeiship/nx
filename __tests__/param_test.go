package nx

import (
	"strings"
	"testing"

	nx "github.com/afeiship/nx/lib"
)

// $ go test __tests__/param_test.go -v
func TestParamData(f *testing.T) {
	datamap := make(map[string]any)
	datamap["name"] = "afei"
	datamap["age"] = 25

	res := nx.Param(datamap)
	if strings.Contains(res, "age=25") || strings.Contains(res, "name=afei") {
		f.Log("TestParamData pass")
	}
}
