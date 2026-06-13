package lib

import (
	"encoding/json"
	"fmt"
)

func JSONParse[T any](jsonStr string) T {
	var data T
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(fmt.Sprintln("JSON不合法: ", jsonStr, err))
	}
	return data
}
