package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 对字符串中的所有单词进行倒排。
//
// 说明：
//
// 1、构成单词的字符只有26个大写或小写英文字母；
//
// 2、非构成单词的字符均视为单词间隔符；
//
// 3、要求倒排后的单词间隔符以一个空格表示；如果原字符串中相邻单词间有多个间隔符时，倒排转换后也只允许出现一个空格间隔符；
//
// 4、每个单词最长20个字母；

// 示例
// $bo*y gi!r#l
// l r gi y bo
func main() {
	//总之先从终端拿值
	reader := bufio.NewReader(os.Stdin)
	strInput, _ := reader.ReadString('\n')

	//循环输入先把不是字母的都变空格全塞进切片
	var inputSlic []rune
	for _, s := range strInput {
		if s >= 'A' && s <= 'z' {
			inputSlic = append(inputSlic, s)
		} else {
			inputSlic = append(inputSlic, ' ')
		}
	}

	//把并连的空格处理掉,转字符串再按空格切割，多出来的会按单个处理
	removeSep := strings.Fields(string(inputSlic))

	//倒序打印就完了
	for i := len(removeSep); i > 0; i-- {
		fmt.Print(removeSep[i-1] + " ")
	}

}
