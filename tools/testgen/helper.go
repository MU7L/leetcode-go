package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

// 生成测试用例
func generateTestCases(sourceFile string) []*TestCase {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, sourceFile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	builder := NewTestTaskBuilder()
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if funcDecl.Doc == nil {
			continue
		}

		handleFlag := func(fl TestgenFlag, param TestCaseParam) {
			switch fl {
			case FL_INPUT:
				err := builder.Input(param, funcDecl.Name.Name)
				if err != nil {
					panic(err)
				}
			case FL_WANT:
				index, c := builder.Want(param)
				fmt.Printf("🌟 case%d: %v\n", index, c)
			}
		}
		parseTestCaseParam(funcDecl.Doc.Text(), handleFlag)
	}
	cases, err := builder.Build()
	if err != nil {
		panic(err)
	}
	return cases
}

// 读取一个注释块 解析为测试用例参数
func parseTestCaseParam(comment string, onFlag func(fl TestgenFlag, param TestCaseParam)) {
	handleLine := func(fl TestgenFlag, line string) {
		tokens := strings.Fields(line)
		if len(tokens) < 3 {
			panic(fmt.Sprintf("[testgen#parseTestCaseParam] 缺少参数: %v", tokens))
		}
		valueJson := tokens[2]
		if !json.Valid([]byte(valueJson)) {
			panic(fmt.Sprintf("[testgen#parseTestCaseParam] 非法JSON: %s", valueJson))
		}
		onFlag(fl, TestCaseParam{
			ParamType:  tokens[1],
			ParamValue: valueJson,
		})
	}

	for _, line := range strings.Split(comment, "\n") {
		if strings.HasPrefix(line, string(FL_INPUT)) {
			handleLine(FL_INPUT, line)
		} else if strings.HasPrefix(line, string(FL_WANT)) {
			handleLine(FL_WANT, line)
		}
	}
}

//go:embed main.tmpl
var mainTmpl string

// 生成测试main文件
func generateTestFile(cases []*TestCase, dest string) {
	t := template.Must(template.New("mainTmpl").Parse(mainTmpl))

	file, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, cases)
	if err != nil {
		panic(err)
	}
}
