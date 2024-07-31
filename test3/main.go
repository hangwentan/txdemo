package main

import (
	"fmt"
	"sync"
	"time"
)

func NumsSum(nums []int, target int) {

	map1 := make(map[int]int, 0)

	for index, num := range nums {
		map1[num] = index
		if _, ok := map1[target-num]; ok {
			fmt.Printf("%v,%v\n", index, map1[target-num])
		}

	}

}

func GetUrl() {

	url := []string{"url1", "url2"}

	wait := sync.WaitGroup{}

	for i := 0; i < len(url); i++ {
		wait.Add(1)
		go func() {
			fmt.Printf("%v\n", url[i])
			wait.Done()
			time.Sleep(5)
		}()

	}

	wait.Wait()

	//fmt.Println("ok")

}

func GetUrl2() {
	Url := []string{"url1", "url2"}

	wait := sync.WaitGroup{}

	for i := 0; i < len(Url); i++ {
		wait.Add(1)
		go func() {
			fmt.Println(Url[i])
			wait.Done()
			time.Sleep(10)
		}()

	}

	wait.Wait()

}

func QuickSort(nums []int) {
	if len(nums) < 1 {
		return
	}
	pivot := nums[0]

	left, right := 0, len(nums)-1

	for i := 1; i <= right; {
		if pivot > nums[i] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else if pivot < nums[i] {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}
	QuickSort(nums[:left])
	QuickSort(nums[right+1:])
}

func binarySearch(nums []int, target int) {

}

func main() {
	//NumsSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7)
	//GetUrl()
	GetUrl2()
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(nums)
	fmt.Println(nums)
}
