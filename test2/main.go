package main

import "fmt"

// 已知数组 A, B, 如果 A 中元素在 B 数组存在，打印出这个元素的下标，B 数组是不重复的.
// Input: [5, 3, 1, 5, 4] [5, 3]
// Output: [0, 1, 3]

func main() {
	//a := []int{5, 3, 1, 5, 4}
	//b := []int{5, 3}
	//be := make(map[int]int, 0)
	//for index, v := range b {
	//	be[v] = index
	//}
	//
	//nums := make([]int, 0)
	//
	//for index, v := range a {
	//
	//	if _, ok := be[v]; ok {
	//		nums = append(nums, index)
	//	}
	//
	//	//for _, v2 := range b {
	//	//	if v2 == v {
	//	//		nums = append(nums, index)
	//	//	}
	//	//
	//	//}
	//}
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//
	//fmt.Printf("%v", nums)

	fmt.Printf("方法2 中心扩展法 结果： %+v\n", logCenterString("ababababcdcdcdcdcdc"))
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
}

// 获取最长回文子串

func logCenterString(str string) string {

	if len(str) < 1 {
		return str
	}
	var start, end int

	for i := 0; i < len(str); i++ {

		len1 := center(str, i, i)
		len2 := center(str, i, i+1)
		maxLen := len1
		if len1 < len2 {
			maxLen = len2
		}

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return str[start : end+1]
}

func center(str string, left int, right int) int {

	for left >= 0 && len(str) > right && str[left] == str[right] {
		left--
		right++
	}

	return right - left - 1
}

// 二分查找

func BinarySearch(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid - 1
		}
	}
	return -1
}

// 冒泡排序

// 快速排序

func QuickSort(nums []int) []int {

	porint := nums[0]

	left, right := 0, len(nums)-1

	for i := 1; i < right; {

		if porint > nums[i] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else if porint > nums[i] {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}
	return nums
}
