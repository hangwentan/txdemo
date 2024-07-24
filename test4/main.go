package main

import "fmt"

func NumSum(nums []int, target int) {

	map1 := make(map[int]int, 0)

	for key, value := range nums {

		map1[value] = key

		if _, ok := map1[target-value]; ok {
			fmt.Printf("两数之和：%v %v\n", map1[target-value], key)
		}

	}
}

// 冒泡排序
func BobbleSort(nums []int) {

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] > nums[j] {
				n := nums[i]
				nums[i] = nums[j]
				nums[j] = n
			}

		}
	}
	return
}

// 快速排序

func QuickSort(nums []int) {

	if len(nums) < 1 {
		return
	}

	Porint := nums[0]
	left, right := 0, len(nums)-1

	for i := 1; i <= right; {

		if Porint > nums[i] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else if Porint < nums[i] {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}

	}
	QuickSort(nums[:left])
	QuickSort(nums[right+1:])
}

func BinarySort(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid - 1
		}
	}
	return -1
}

func 金字塔(num int) {

	for i := 1; i <= num; i++ {
		n1 := num - i
		n2 := i*2 - 1
		for i := 0; i < n1; i++ {
			fmt.Printf(" ")
		}
		for i := 0; i < n2; i++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	NumSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	BobbleSort(nums)
	fmt.Printf("冒泡：%v\n", nums)
	nums = []int{9, 8, 7, 4, 5, 6, 3, 2, 1}
	QuickSort(nums)
	fmt.Printf("快排：%v\n", nums)
	fmt.Printf("二分：%v\n", BinarySort(nums, 8))
	金字塔(5)
}
