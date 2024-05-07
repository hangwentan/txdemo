package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		fmt.Printf("s>>%v\n", s)
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}

	for _, str := range strs {
		cnt := [26]int{}
		for _, s := range str {
			cnt[s-'a']++
			fmt.Printf("%v:=%v\n", s, s-'a')
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Printf("str1>>%v\n", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("str2>>%v\n", groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("a:=%v\n", 'a')
}
