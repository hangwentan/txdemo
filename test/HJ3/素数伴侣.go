package main

import (
	"fmt"
	"io"
)

//题目描述
//若两个正整数的和为素数，则这两个正整数称之为“素数伴侣”，如2和5、6和13，它们能应用于通信加密。现在密码学会请你设计一个程序，从已有的 N （ N 为偶数）个正整数中挑选出若干对组成“素数伴侣”，挑选方案多种多样，例如有4个正整数：2，5，6，13，如果将5和6分为一组中只能得到一组“素数伴侣”，而将2和5、6和13编组将得到两组“素数伴侣”，能组成“素数伴侣”最多的方案称为“最佳方案”，当然密码学会希望你寻找出“最佳方案”。
//输入:
//有一个正偶数 n ，表示待挑选的自然数的个数。后面给出 n 个具体的数字。
//输出:
//输出一个整数 K ，表示你求得的“最佳方案”组成“素数伴侣”的对数。

func isPrime(num int) bool {

	if num == 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func match(odd int, evens []int, visited map[int]int, suited map[int]int) bool {
	for _, even := range evens {
		if isPrime(odd+even) && visited[even] == 0 {
			visited[even] = 1
			if suited[even] == 0 || match(suited[even], evens, visited, suited) {
				suited[even] = odd
				return true
			}
		}
	}
	return false
}

func main() {
	for {
		var n int
		var num int
		var odds []int
		var evens []int

		c, err := fmt.Scanf("%d\n", &n)
		if c == 0 || err == io.EOF {
			break
		}

		//分奇数和偶数
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &num)
			if num%2 == 0 {
				evens = append(evens, num)
			} else {
				odds = append(odds, num)
			}
		}

		var suited map[int]int = make(map[int]int)
		var res int
		for i := 0; i < len(odds); i++ {

			var visited map[int]int = make(map[int]int)

			ok := match(odds[i], evens, visited, suited)
			if ok {
				res++
			}
		}

		fmt.Println(res)
	}
}
