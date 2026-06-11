package json

import (
	"encoding/json"
	. "leetcode-go/utils/logger"
)

var logger = NewLogger("JSON")

func Parse[T any](jsonStr string) *T {
	var result T
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		logger.Fatal("解析失败: %s", jsonStr)
	}
	return &result
}
