package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

//编写一个程序，将输入字符串中的字符按如下规则排序。
//规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。
//如，输入： Type 输出： epTy
//规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。
//如，输入： BabA 输出： aABb
//规则 3 ：非英文字母的其它字符保持原来的位置。
//如，输入： By?e 输出： Be?y

func main() {
	type indexLetter struct {
		index int
		c     rune
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		chars := []rune(input.Text())
		otherChars := make([]bool, len(chars))
		letters := []indexLetter{}
		for i, c := range chars {
			if c < 'A' || (c > 'Z' && c < 'a') || c > 'z' {
				otherChars[i] = true
			} else {
				il := indexLetter{i, c}
				letters = append(letters, il)
			}
		}
		plus := 'a' - 'A'
		sort.Slice(letters, func(i, j int) bool {
			if math.Abs(float64(letters[i].c-letters[j].c)) == float64(plus) || letters[i].c == letters[j].c {
				return letters[i].index < letters[j].index
			}
			ti, tj := letters[i].c, letters[j].c
			if ti >= 'a' {
				ti -= plus
			}
			if tj >= 'a' {
				tj -= plus
			}
			return ti < tj
		})
		for i, c := range chars {
			if otherChars[i] == true {
				fmt.Print(string(c))
			} else {
				fmt.Print(string(letters[0].c))
				letters = letters[1:len(letters)]
			}
		}
	}

}
