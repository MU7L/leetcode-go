package main

import "strings"

type TestCaseParam struct {
	ParamType  string // 测试参数类型
	ParamValue string // 测试参数值
}

type TestCase struct {
	FuncName string           // 被测函数名
	Inputs   []*TestCaseParam // 被测入参
	// Outputs  *TestCaseParam
}

func newTestCase(fn string) *TestCase {
	return &TestCase{
		FuncName: fn,
		Inputs:   make([]*TestCaseParam, 0),
	}
}

func (c TestCase) String() string {
	var b strings.Builder
	b.WriteString(c.FuncName)
	b.WriteString("(")
	for i, input := range c.Inputs {
		b.WriteString(input.ParamValue)
		if i < len(c.Inputs)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")
	return b.String()
}
