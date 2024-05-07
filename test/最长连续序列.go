package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				fmt.Printf("currentNum>>%v\n", currentNum+1)
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

func main() {
	fmt.Println("index>>", longestConsecutive([]int{3, 8, 7, 11, 1, 2, 9, 4}))
}
