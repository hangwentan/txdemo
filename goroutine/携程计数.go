package main

import (
	"fmt"
	"sync"
)

const (
	PageSize = 500
)

// 用channel实现

func chanAddNum(nums []int) {

	var mxt sync.WaitGroup
	resultCh := make(chan int, len(nums)/PageSize)

	for i := 0; i < len(nums); i += PageSize {

		mxt.Add(1)
		number := nums[i : i+PageSize]

		go func(number []int) {
			mxt.Done()

			sum := 0

			for _, i3 := range number {
				sum += i3
			}

			resultCh <- sum

		}(number)
	}

	mxt.Wait()
	close(resultCh)

	var total int

	for sum := range resultCh {
		total += sum
	}

	fmt.Println("chan addNum：", total)
}

func main() {

	nums := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		nums[i] = i + 1
	}

	chanAddNum(nums)
}
