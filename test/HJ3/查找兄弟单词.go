package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputs := strings.Split(input.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])       // 有多少组数据
	index, _ := strconv.Atoi(inputs[n+2]) // 输出结果的第几项
	tag := inputs[n+1]                    // 以tag字符串为基准
	inputs = inputs[1 : n+1]              // 需要处理的字符串

	map1 := make(map[string]int) // 基准字符每个字母出现的次数
	res := make([]string, 0)     // 符合条件字符串
	for _, v := range tag {
		map1[string(v)]++
	}

	for _, v := range inputs { // 遍历每一项
		map2 := make(map[string]int)
		// 只有长度相同且不全等的字符串才进入循环
		if len(v) == len(tag) && v != tag {
			ta := 0
			for _, v := range v { // 当前字符生成map
				map2[string(v)]++
			}
			for _, v := range tag { // 两个map作比较
				if map1[string(v)] != map2[string(v)] {
					ta = 1
				}
			}
			if ta == 0 {
				res = append(res, v)
			}
		}
	}
	sort.Strings(res)
	fmt.Println(len(res))
	if len(res) < index {
		return
	}
	fmt.Print(res[index-1])
}
