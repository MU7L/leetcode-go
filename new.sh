#!/bin/bash

# 获取问题编号和名称
read -p "1️⃣ No  : " number
read -p "2️⃣ Name: " name

# 创建目录
dir_name="solutions/${number}-${name}"
mkdir -p "$dir_name"

# 创建 main.go
cat > "$dir_name/main.go" << EOF
// https://leetcode.com/problems/${name}/

package main

import (
	"fmt"
	"leetcode-go/utils"
)


func main() {
	_input := ""
    input := utils.JSONParse[](input)
}
EOF

echo "✅ Created problem: $dir_name"