// https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-i/

package main

import (
	"fmt"
	"leetcode-go/utils/json"
)

func main() {
	_input := "[[1,2],[1,3],[3,4],[3,5]]"
	input := json.Parse[[][]int](_input)
	ans := assignEdgeWeights(*input)
	fmt.Println(ans)
}

func assignEdgeWeights(edges [][]int) int {
	// 1. 建图
	g := genG(edges)
	// 2. 找深度
	depth := dfs(g, 1, make(map[int]bool))
	// 3. n个数排列组合 只要有奇数个1则满足条件
	return pow2(depth - 2)
}

func genG(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return g
}

func dfs(g map[int][]int, root int, visited map[int]bool) int {
	currentMaxDepth := 0
	visited[root] = true
	for _, child := range g[root] {
		if _, ok := visited[child]; ok {
			continue
		}
		newMaxDepth := dfs(g, child, visited)
		currentMaxDepth = max(currentMaxDepth, newMaxDepth)
	}
	return currentMaxDepth + 1
}

const MOD = 1_000_000_007

var pow2Cache = make(map[int]int)

func pow2(n int) int {
	if value, ok := pow2Cache[n]; ok {
		return value
	}

	if n == 0 {
		return 1
	}

	tmp := pow2(n / 2)
	var res int
	if n%2 == 1 {
		res = tmp * tmp * 2 % MOD
	} else {
		res = tmp * tmp % MOD
	}
	pow2Cache[n] = res
	return res
}
