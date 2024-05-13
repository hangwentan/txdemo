package main

import (
	"bufio"
	"fmt"
	"os"
)

//对输入的字符串进行加解密，并输出。
//加密方法为：
//当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；
//当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；
//其他字符不做变化。
//解密方法为加密的逆过程。

func main() {

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		s := input.Text()
		li := make([]byte, 0)

		for i := 0; i < len(s); i++ {

			if s[i] == 'Z' {
				li = append(li, 'a')
			} else if s[i] == 'z' {
				li = append(li, 'A')
			} else if s[i] == '9' {
				li = append(li, '0')
			} else if s[i] >= 'a' && s[i] < 'z' {
				li = append(li, s[i]-32+1)
			} else if s[i] >= 'A' && s[i] < 'Z' {
				li = append(li, s[i]+32+1)
			} else if s[i] >= '0' && s[i] < '9' {
				li = append(li, s[i]+1)
			} else {
				li = append(li, s[i])
			}

		}

		fmt.Println(string(li))
		input.Scan()

		s = input.Text()
		li = make([]byte, 0)
		for i := 0; i < len(s); i++ {

			if s[i] == 'a' {
				li = append(li, 'Z')
			} else if s[i] == 'A' {
				li = append(li, 'z')
			} else if s[i] == '0' {
				li = append(li, '9')
			} else if s[i] > 'a' && s[i] <= 'z' {
				li = append(li, s[i]-32-1)
			} else if s[i] > 'A' && s[i] <= 'Z' {
				li = append(li, s[i]+32-1)
			} else if s[i] > '0' && s[i] <= '9' {
				li = append(li, s[i]-1)
			} else {
				li = append(li, s[i])
			}

		}

		fmt.Println(string(li))
	}
}