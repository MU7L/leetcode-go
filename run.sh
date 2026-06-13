#!/bin/bash

# 获取问题编号
no=$1

# 创建目录
dir_name="solutions/${no}"
mkdir -p "$dir_name"

# 创建 main.go
go generate "./solutions/${no}/solution.go"

# 执行 main.go
go run "./solutions/${no}"
