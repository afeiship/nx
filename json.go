package nx

import (
	"encoding/json"
)

func JsonStringify(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonParse(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
