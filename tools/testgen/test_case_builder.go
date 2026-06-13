package main

import "fmt"

type TestTaskBuilder struct {
	cases     []*TestCase
	completed bool
}

func NewTestTaskBuilder() *TestTaskBuilder {
	return &TestTaskBuilder{
		cases:     make([]*TestCase, 0),
		completed: true,
	}
}

func (this *TestTaskBuilder) Input(param TestCaseParam, fn string) error {
	if this.completed {
		c := newTestCase(fn)
		c.Inputs = append(c.Inputs, &param)
		this.cases = append(this.cases, c)
	} else {
		_, lastCase := this.tail()
		if lastCase == nil {
			return fmt.Errorf("[TestTaskBuilder.Input] 检查测试用例顺序，当前测试用例为空")
		}
		if lastCase.FuncName != fn {
			return fmt.Errorf("[TestTaskBuilder.Input] 检查测试用例顺序，最近两次Input不属于同一个被测函数: %s, %s)", lastCase.FuncName, fn)
		}
		lastCase.Inputs = append(lastCase.Inputs, &param)
	}
	this.completed = false
	return nil
}

func (this *TestTaskBuilder) Want(param TestCaseParam) (int, *TestCase) {
	// TODO: output
	this.completed = true
	return this.tail()
}

func (this *TestTaskBuilder) Build() ([]*TestCase, error) {
	if !this.completed {
		return nil, fmt.Errorf("[TestTaskBuilder.Build] 检查测试用例顺序，末尾没有@want")
	}
	return this.cases, nil
}

func (this *TestTaskBuilder) tail() (int, *TestCase) {
	size := len(this.cases)
	if size == 0 {
		return 0, nil
	}
	return size - 1, this.cases[size-1]
}
