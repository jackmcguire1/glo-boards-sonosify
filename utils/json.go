package utils

import "encoding/json"

func ToJSON(i interface{}) string {
	data, _ := json.Marshal(i)
	return string(data)
}

func ToRawMessage(i interface{}) json.RawMessage {
	data, _ := json.Marshal(i)
	return data
}

func ToJSONArray(i ...interface{}) string {
	return ToJSON(i)
}
