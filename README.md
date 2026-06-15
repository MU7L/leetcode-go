# leetcode代码仓库 golang版

1. 做题记录收集
2. 解析注释生成测试用例和测试代码
3. 本地运行/调试 leetcode 代码

## Step.0 使用前准备

```shell
go build ./tools/testgen # 编译代码生成器
```

## Step.1 生成模板代码

以 [3838.带权单词映射](https://leetcode.cn/problems/weighted-word-mapping/description/)为例

```shell
sh ./new.sh 3838
# 或者可以先给脚本增加执行权限
chmod +r ./new.sh
./new.sh 3838
```

脚本将在 `./solutions/{no}/solution.go` 创建模板文件

## Step.2 解题略

## Step.3 生成测试代码

用注释添加测试用例

- `@input type value` 表示函数的一个入参，需要包含类型、json字符串两部分。一个函数可能有多个入参，因此需要多个 `@input`
- `@want type value` 表示该测试用例的期望输出。一个用例只能有一个输出
- 每一个用例由多行 `@input` 和一行 `@want` 组成。如果需要多个测试用例，可以增加一个空行隔开

例如

```go
// @input []string ["abcd","def","xyz"]
// @input []int [5,3,12,14,1,2,3,2,10,6,6,9,7,8,7,10,8,9,6,9,9,8,3,7,7,2]
// @want string "rij"
//
// @input []string ["a","b","c"]
// @input []int [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
// @want string "yyy"
func mapWordWeights(words []string, weights []int) string {
	// 在这里完成剩余代码...
}
```

接下来运行 `./run.sh 3838`

1. 脚本首先执行 `go generate` ，调用 Step.0 中构建的 testgen 来生成测试入口文件
2. 测试入口文件将位于 `./solutions/{no}/main.go`，本质上是将测试用例传入到函数，并打印返回值
3. 脚本最后执行 `go run`，终端打印最终的执行结果

## TODO

- 当前版本只有测试结果输出，结果对比待开发
- 未支持复杂输入输出（树、链表、图）
