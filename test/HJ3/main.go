package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//  获取字符串最后一个单词的长度
	GetLastString()

	// 计算某字符出现次数
	GetStringNum()

	// 检测括号是否闭合
	fmt.Printf("%v\n", is_valid("{}()[]"))

	// 明明的随机数
	GetRand()

	// 字符串分隔 输入一个字符串，请按长度为8拆分每个输入字符串并进行输出；长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
	GetStringSplit()

	// 进制转换 写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。
	GetCodeCheck()

	// 获取一个整数 n 的所有质因数，并将它们打印出来
	GetPrintFactor()

	//  取近似值 写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。如果小数点后数值大于等于 0.5 ,向上取整；小于 0.5 ，则向下取整。
	GetFloorAcquireInt()

	// 合并表记录 合并主键相同的记录，并将其value使用sum统计
	GetMergeIndexLog()

	GetDesNumber()

	// 字符个数统计
	GetStrNumCount()

	//数字颠倒
	GetStringSwift()

	// 购物单
	GetShopCar()

	GetRemoveCoord()
}

func GetLastString() {

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		arr := strings.Split(input.Text(), " ")
		fmt.Println("获取字符串最后一个单词的长度:", len(arr[len(arr)-1]))
	}
}

func GetStringNum() {

	map1 := make(map[byte]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		str := input.Text()

		input.Scan()

		str1 := input.Text()
		str = strings.ToLower(str)
		str1 = strings.ToLower(str1)
		b1 := str1[0]

		for i := 0; i < len(str); i++ {
			map1[str[i]]++
		}

		fmt.Println(map1[b1])
	}
}

func is_valid(str string) bool {

	strlen := len(str)

	if strlen == 0 {
		return false
	}
	if strlen/2 == 1 {
		return false
	}

	stack := list.New()
	pairsMap := map[byte]byte{'}': '{', ')': '(', ']': '['}
	for i := 0; i < strlen; i++ {
		value := str[i]

		if pairsMap[value] == 0 {
			stack.PushFront(value)
		} else {

			if stack.Len() == 0 || stack.Front().Value != pairsMap[value] {
				return false
			}
			stack.Remove(stack.Front())
		}
	}
	return stack.Len() == 0
}

func GetRand() {
	n := 0

	fmt.Scan(&n)

	narr := make([]int, 501)

	N := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&N)
		if narr[N] > 0 {
			continue
		}
		narr[N] = 18
	}

	for i := range narr {
		if narr[i] > 0 {
			fmt.Println(i)
		}
	}
}

func GetStringSplit() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		panic("error")
	}
	str := []rune(scanner.Text())
	m, n := len(str)/8, len(str)%8
	for i := 0; i < m; i++ {
		fmt.Println(string(str[i*8 : i*8+8]))
	}
	if n != 0 {
		fmt.Print(string(str[len(str)-n:]))
		for j := 0; j < 8-n; j++ {
			fmt.Print(0)
		}
		fmt.Println()
	}
}

func GetCodeCheck() {

	input := bufio.NewScanner(os.Stdin)

	for {
		input.Scan()
		temp := input.Text()
		res, _ := strconv.ParseInt(temp, 0, 32)
		fmt.Printf("\n%v", res)
	}
}

func GetPrintFactor() {

	input := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(input, &n)
	for i := 2; i*i <= n; i++ {

		for n%i == 0 {
			n /= i
			fmt.Printf("%d ", i)
		}
	}
	if n != 1 {
		fmt.Printf("%d", n)
	}
}

func GetFloorAcquireInt() {

	var num float32
	fmt.Scan(&num)

	res := int(num * 10)
	if res%10 >= 5 {
		fmt.Println(res/10 + 1)
	} else {
		fmt.Println(res / 10)
	}
}

func GetMergeIndexLog() {
	var n int
	var key []int

	map1 := make(map[int]int, 0)

	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		var key int
		var value int
		fmt.Scan(&key, &value)
		map1[key] += value
	}

	for k := range map1 {
		key = append(key, k)
	}

	sort.Ints(key)
	for _, v := range key {
		fmt.Println(v, map1[v])
	}
}

func GetDesNumber() {
	n, r := 0, 0
	map1 := make(map[int]int)
	fmt.Scanf("%d", &n)
	for ; n > 0; n /= 10 {
		k := n % 10
		if _, ok := map1[k]; ok {

		} else {
			map1[k] = 0
			r = r*10 + k
		}
	}
	fmt.Printf("%d", r)
}

func GetStrNumCount() {

	var str string
	fmt.Scan(&str)
	s := []byte(str)
	arr_hash := make(map[byte]int)
	count := 0
	for i := 0; i < len(s); i++ {

		if arr_hash[s[i]] == 0 {
			arr_hash[s[i]] = 1
			count++
		}
		continue
	}
	fmt.Println(count)
}

func GetStringSwift() {
	var str string
	fmt.Scan(&str)

	s := []byte(str)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(string(s))
}

func GetShopCar() {
	var money, count int
	fmt.Scan(&money, &count)

	goods := make([][]int, count)
	for i := 0; i < count; i++ {
		goods[i] = make([]int, 3)
		fmt.Scan(&goods[i][0], &goods[i][1], &goods[i][2])
	}

	costs := make([][]int, 0)
	prices := make([][]int, 0)
	for i := 0; i < count; i++ {
		if goods[i][2] != 0 {
			continue
		}
		cost := make([]int, 0)
		price := make([]int, 0)
		cost = append(cost, goods[i][0]*goods[i][1])
		price = append(price, goods[i][0])
		for j := 0; j < count; j++ {
			if goods[j][2]-1 == i {
				cost = append(cost, goods[i][0]*goods[i][1]+cost[0])
				price = append(price, goods[i][0]+cost[0])
			}
		}
		if len(cost) == 3 {
			cost = append(cost, cost[1]+cost[2]-cost[0])
			price = append(price, price[1]+price[2]-price[0])
		}
		costs = append(costs, cost)
		prices = append(prices, price)
	}
	size := len(costs)
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, money+1)
		dp[i][0] = 0
	}
	for i := 0; i < money+1; i += 10 {
		dp[0][i] = 0
	}
	for i := 1; i < size+1; i++ {
		for j := 0; j < money+1; j += 10 {
			max := dp[i-1][j]
			for k, v := range prices[i-1] {
				if j-v >= 0 {
					if max < dp[i-1][j-v]+costs[i-1][k] {
						max = dp[i-1][j-v] + costs[i-1][k]
					}
				}
			}
			dp[i][j] = max
		}
	}
	fmt.Println(dp[size][money])
}

// 坐标移动
// 描述
// 开发一个坐标计算工具， A表示向左移动，D表示向右移动，W表示向上移动，S表示向下移动。从（0,0）点开始移动，从输入字符串里面读取一些坐标，并将最终输入结果输出到输出文件里面。
// 输入：
// 合法坐标为A(或者D或者W或者S) + 数字（两位以内）坐标之间以;分隔。
// 非法坐标点需要进行丢弃。如AA10;  A1A;  $%$;  YAD; 等。
// 下面是一个简单的例子 如：
// A10;S20;W10;D30;X;A1A;B10A11;;A10;
// 处理过程：
// 起点（0,0）
// +   A10   =  （-10,0）
// +   S20   =  (-10,-20)
// +   W10  =  (-10,-10)
// +   D30  =  (20,-10)
// +   x    =  无效
// +   A1A   =  无效
// +   B10A11   =  无效
// +  一个空 不影响
// +   A10  =  (10,-10)
// 结果 （10， -10）
func GetRemoveCoord() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	var x, y int = 0, 0
	s1 := strings.Split(s, ";")
	for i := 0; i < len(s1); i++ {
		s2 := s1[i]
		if len(s2) >= 2 {
			first_num := s2[0]
			if first_num == 'A' || first_num == 'D' || first_num == 'S' || first_num == 'W' {
				s2 = s2[1:]
				num, err := strconv.Atoi(s2)
				if err == nil {
					switch first_num {
					case 'A':
						x -= num
					case 'S':
						y -= num
					case 'W':
						y += num
					case 'D':
						x += num
					}
				}
			}
		}

		continue
	}
	fmt.Printf("%v,%v", x, y)
}
