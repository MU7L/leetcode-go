#!/bin/bash

# 获取问题编号
no=$1

# 创建目录
dir_name="solutions/${no}"
mkdir -p "$dir_name"

# 创建 main.go
cat > "$dir_name/solution.go" << EOF
package main

//go:generate ../../testgen ./solution.go

// @input
// @want

EOF

echo "✅ Created problem: $dir_name"