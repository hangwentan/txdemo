package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	var inputArr []string
	for {
		input, _ := inputReader.ReadString('\n')
		if len(input) == 0 {
			break
		}
		inputArr = append(inputArr, input)
	}
	//将字符串转化为数组
	strR := strings.Fields(inputArr[1])
	var strRArr []int
	for _, v := range strR {
		n, _ := strconv.Atoi(string(v))
		strRArr = append(strRArr, n)
	}
	//去掉首个元素
	strRSlice := strRArr[1:]
	//排序strRArr
	quickSort(strRSlice, 0, len(strRSlice)-1)
	strRSlice = removeDuplication_sort(strRSlice)

	strI := strings.Fields(inputArr[0])
	var strIArr []string
	for _, v := range strI {
		strIArr = append(strIArr, v)
	}
	//去掉首个元素
	strISlice := strIArr[1:]
	//fmt.Println(strings.Contains(inputArr[0], "5"))
	var res []int
	for _, v := range strRSlice {
		n := strconv.Itoa(v)
		if strings.Contains(inputArr[0][strings.Index(inputArr[0], " "):], n) {
			var j int //该数匹配到了几个数字
			var temp []int
			res = append(res, v)
			for i, value := range strISlice {
				if strings.Contains(value, n) {
					nb, _ := strconv.Atoi(string(value))
					temp = append(temp, i)
					temp = append(temp, nb)
					j++
				}
			}
			res = append(res, j)
			res = append(res, temp...)
		}
	}
	lens := len(res)
	//头插
	res = append(res, 0)
	index := 0
	copy(res[index+1:], res[index:])
	res[index] = lens
	for _, v := range res {
		fmt.Printf("%v ", v)
	}

}

func quickSort(values []int, left int, right int) {
	key := values[left] //取出第一项
	p := left
	i, j := left, right

	for i <= j {
		//由后开始向前搜索(j--)，找到第一个小于key的值values[j]
		for j >= p && values[j] >= key {
			j--
		}

		//第一个小于key的值 赋给 values[p]
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= key && i <= p {
			i++
		}

		if i < p {
			values[p] = values[i]
			p = i
		}

		values[p] = key
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}
}

func removeDuplication_sort(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}

	j := 0
	for i := 1; i < length; i++ {
		if arr[i] != arr[j] {
			j++
			if j < i {
				swap(arr, i, j)
			}
		}
	}

	return arr[:j+1]
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
