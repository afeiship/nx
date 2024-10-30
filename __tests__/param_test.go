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
		f.Log(res)
	}
}

func TestParamWithBaseURL(f *testing.T) {
	datamap := make(map[string]any)
	datamap["name"] = "afei"
	datamap["age"] = 25
	baseURL := "https://www.example.com"
	res := nx.Param(datamap, baseURL)
	if strings.Contains(res, "age=25") || strings.Contains(res, "name=afei") || strings.Contains(res, "https://www.example.com") {
		f.Log("TestParamWithBaseURL pass")
		f.Log(res)
	}
}
