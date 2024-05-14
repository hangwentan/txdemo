package main

import (
	"fmt"
	"strings"
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

	for _, value := range url {
		wait.Add(1)
		go func() {
			fmt.Printf("%v\n", value)
			wait.Done()
			time.Sleep(5)
		}()

	}

	wait.Wait()

	//fmt.Println("ok")

}

func RequestChan() {
	ch := make(chan int)

	go func() {
		fmt.Println("Received data form chan:", <-ch)
	}()

	ch <- 42
	close(ch)
	fmt.Println("Send data form chan")
}

func Nums金字塔() {

	for i := 1; i <= 3; i++ {
		N := 3 - i
		X := i*2 - 1
		first := strings.Repeat(" ", N)
		last := strings.Repeat("*", X)
		fmt.Println(first + last)

	}

}

func QuickSort(nums []int) {
	if len(nums) < 1 {
		return
	}

	mid := nums[0]

	left, right := 0, len(nums)-1

	for i := 0; i <= right; {
		if nums[i] < mid {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else if nums[i] > mid {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}

	QuickSort(nums[:left])
	QuickSort(nums[right+1:])
}

func BinarySearch(nums []int, target int) (idx int) {

	if len(nums) < 2 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if target == nums[mid] {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			low = mid - 1
		}

	}
	return -1
}

func BubbleSort(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func main() {
	NumsSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7)
	GetUrl()
	RequestChan()
	Nums金字塔()
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(arr)
	fmt.Println("QuickSort", arr)
	fmt.Println("BinarySearch", BinarySearch(arr, 9))
	arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(arr)
	fmt.Println("BubbleSort", arr)
}
