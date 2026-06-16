package main

//go:generate ../../testgen ./solution.go

// @input string "a#b%*"
// @want string "ba"
func processStr(s string) string {
	// 使用byte比使用rune快
	res := make([]byte, 0)

	for _, r := range s {
		switch r {
		case '*':
			if len(res) == 0 {
				continue
			}
			res = res[:len(res)-1]
		case '#':
			res = append(res, res...)
		case '%':
			// 比使用 slices.Reverse 快
			for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
				res[i], res[j] = res[j], res[i]
			}
		default:
			res = append(res, byte(r))
		}
	}

	// 不需要 strings.Builder
	return string(res)
}
