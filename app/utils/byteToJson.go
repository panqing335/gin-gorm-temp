package util

import "encoding/json"

func ByteToJson(jsonData []byte) map[string]any {
	var v any
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]any)
	return data
}
