package main

//Catcher是MCA国的情报员，他工作时发现敌国会用一些对称的密码进行通信，比如像这些ABBA，ABA，A，123321，但是他们有时会在开始或结束时加入一些无关的字符以防止别国破解。比如进行下列变化 ABBA->12ABBA,ABA->ABAKK,123321->51233214　。因为截获的串太长了，而且存在多种可能的情况（abaaab可看作是aba,或baaab的加密形式），Cathcer的工作量实在是太大了，他只能向电脑高手求助，你能帮Catcher找出最长的有效密码串吗？
//
//输入描述：
//输入一个字符串（字符串的长度不超过2500）
//
//输出描述：
//返回有效密码串的最大长度

import "fmt"

func main() {
	//读取输入的字符串
	var str string
	fmt.Scan(&str)

	//调用函数求取字符串中有效密码的最大长度
	res := LongestPsw(str)

	//输出结果
	fmt.Println(res)
}

// LongestPsw 求取字符串中有效密码的最大长度
func LongestPsw(s string) int {
	//动态规划
	maxL := 0  //maxL记录字符串中有效密码的最大长度
	tempL := 1 //tempL记录以当前字符为最右端的最长有效密码
	i := 1
	for i < len(s) {
		if i-tempL-1 >= 0 && s[i] == s[i-tempL-1] {
			tempL += 2
			i++
		} else if s[i] != s[i-1] {
			tempL = 1
			i++
		} else {
			tempL = 2
			i++
			for i < len(s) && s[i] == s[i-1] {
				tempL++
				i++
			}
		}

		if tempL > maxL { //及时判断更新最长长度
			maxL = tempL
		}

	}

	return maxL
}
