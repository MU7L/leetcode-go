package utils

import (
	"encoding/json"
)

func JSONParse[T any](jsonStr string) *T {
	var result T
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		Logger("json").Fatal("解析失败: %s", jsonStr)
	}
	return &result
}
