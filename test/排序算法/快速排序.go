package main

import "fmt"

// QuickSort 函数实现快速排序
func QuickSort(arr []int) {

	if len(arr) < 1 {
		return
	}

	// 选择基准元素（这里简单选择第一个元素）
	pivot := arr[0]

	// 将小于基准的元素放在左边，大于基准的元素放在右边
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {

		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		} else if arr[i] > pivot {
			arr[right], arr[i] = arr[i], arr[right]
			right--
		} else {
			i++
		}
	}
	// 递归排序左右两部分
	QuickSort(arr[:left])
	QuickSort(arr[right+1:])
}

func QuickSortCopy(arr []int) {
	if len(arr) < 1 {
		return
	}

	pivot := arr[0]

	left, right := 0, len(arr)-1

	for i := 1; i <= right; {
		if pivot > arr[i] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
			i++
		} else if pivot < arr[i] {
			arr[right], arr[i] = arr[i], arr[right]
			right--
		} else {
			i++
		}
	}
	QuickSort(arr[:left])
	QuickSort(arr[right+1:])
}

func main() {
	// 待排序的数组
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4}

	// 快速排序
	QuickSort(arr)
	// 打印排序后的数组
	fmt.Println("Sorted array:", arr)
	QuickSortCopy(arr)
	// 打印排序后的数组
	fmt.Println("Sorted array:", arr)
}
