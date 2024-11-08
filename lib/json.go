package nx

import (
	"encoding/json"
)

func JSONParse(input string) (any, error) {
	var result any
	err := json.Unmarshal([]byte(input), &result)
	return result, err
}

func JSONStringify(input any) (string, error) {
	jsonString, err := json.Marshal(input)
	return string(jsonString), err
}
