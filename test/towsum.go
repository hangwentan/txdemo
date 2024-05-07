package main

import "fmt"

func twoSum(number []int, target int) []int {

	hasTarget := map[int]int{}

	for i, x := range number {

		if p, ok := hasTarget[target-x]; ok {
			return []int{p, i}
		}
		hasTarget[x] = i

	}
	return nil
}

func main() {
	number := []int{2, 1, 3, 4, 6, 5, 8}
	fmt.Printf("number:%v", twoSum(number, 9))
}
