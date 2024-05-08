package main

import "fmt"

func bubbleSort1(nums []int) []int {
	if len(nums) < 2 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	intData := []int{10, 2, 3, 6, 9, 1}
	fmt.Printf("data:%v\n", bubbleSort1(intData))
}
