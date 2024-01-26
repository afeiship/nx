package nx

import (
	"encoding/json"
)

func JSONParse(input string) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal([]byte(input), &result)
	return result, err
}

func JSONStringify(input interface{}) (string, error) {
	jsonString, err := json.Marshal(input)
	return string(jsonString), err
}
