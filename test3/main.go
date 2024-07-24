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

func main() {
	//NumsSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7)
	GetUrl()
}
