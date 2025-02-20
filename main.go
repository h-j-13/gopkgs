package main

import "fmt"

// isBalance 检查 k 是否是p的一个平衡数字
func isBalance(k int, p []int) bool {
	// 遍历l~r的区间
	for l := 0; l < len(p)-k+1; l++ {
		r := l + k
		// 检查这个区间是否是1~k的
		m := make([]int, k)
		for i := l; i < r; i++ {
			if p[i] > k {
				continue // 这个组合不行
			}
			m[p[i]-1] = 1 // 通过map记录
		}

		fmt.Print(p[l:r], ":")
		fmt.Println(m)
		// 检查是否是1~k
		match := true
		for i := 0; i < k; i++ {
			if m[i] == 0 {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}
	return false
}

func main() {
	p := []int{4, 1, 3, 2}
	fmt.Println(isBalance(1, p))
	fmt.Println(isBalance(2, p))
	fmt.Println(isBalance(3, p))
	fmt.Println(isBalance(4, p))
}
