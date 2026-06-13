package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("🔧 当前执行路径: ", dir)
	if len(os.Args) < 2 {
		fmt.Println("⚠️ 使用方法: testgen <file.go>")
		os.Exit(1)
	}

	solutionPath := os.Args[1]
	fmt.Println("📃 处理文件: ", path.Join(dir, solutionPath))
	cases := generateTestCases(solutionPath)
	fmt.Printf("✅ 生成测试用例 %d 条", len(cases))
	mainPath := path.Join(dir, solutionPath, "../main.go")
	generateTestFile(cases, mainPath)
	fmt.Println("✅ 生成文件: ", mainPath)
}
