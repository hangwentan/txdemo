package main

import "fmt"

func main() {
	var n, scanN int
	var heights []int

	for {
		scanN, _ = fmt.Scan(&n)
		if scanN == 0 {
			break
		}
		heights = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&heights[i])
		}

		fmt.Println(chorus(heights))
	}
}

// 动态规划
func chorus(heights []int) int {
	// 两个数组分别表示每个人左边与右边最多站的人数
	var leftMost, rightMost = make([]int, len(heights)), make([]int, len(heights))

	// 以每个人为中心求解每个人左边最多站的人数
	for center := 1; center < len(heights); center++ {
		// 根据 center 左边已经知晓的每个人的左边最多人数获得 center 左边最多人数
		for i := 0; i < center; i++ {
			if heights[center] > heights[i] && leftMost[center] < leftMost[i]+1 {
				leftMost[center] = leftMost[i] + 1
			}
		}
	}

	// 以每个人为中心求解每个人右边最多站的人数
	for center := len(heights) - 2; center >= 0; center-- {
		// 根据 center 右边已经知晓的每个人的右边最多人数获得 center 右边最多人数
		for i := len(heights) - 1; i > center; i-- {
			if heights[center] > heights[i] && rightMost[center] < rightMost[i]+1 {
				rightMost[center] = rightMost[i] + 1
			}
		}
	}

	// 获取合唱队的最多人数
	var max = 1
	for i := 0; i < len(heights); i++ {
		if max < leftMost[i]+rightMost[i]+1 {
			max = leftMost[i] + rightMost[i] + 1
		}
	}

	return len(heights) - max
}
