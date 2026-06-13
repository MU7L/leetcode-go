package main

// 测试标记
type TestgenFlag string

const (
	// 测试用例输入
	FL_INPUT TestgenFlag = "@input"
	// 测试用例输出
	FL_WANT TestgenFlag = "@want"
)
