package main

import "fmt"

// 现在数据库有一张表，用来存储一个多叉树，id为主键，pid 表示父节点的 id，已知 "-1" 表示根节点，现在要求打印出从根节点到每个子节点的路径（可以是无序的）。
//
// | id      | pid    |
// |---------|--------|
// | "A"     | "-1"   |
// | "A-1"   | "A"    |
// | "A-2"   | "A"    |
// | "A-3"   | "A"    |
// | "A-2-1" | "A-2"  |
// | "A-2-2" | "A-2"  |
// | "A-2-3" | "A-2"  |
//
// Input: [
//   {
//       "id": "A",
//       "pid": "-1"
//   },
//   {
//       "id": "A-1",
//       "pid": "A"
//   },
//   {
//       "id": "A-2",
//       "pid": "A"
//   },
//   {
//       "id": "A-3",
//       "pid": "A"
//   },
//   {
//       "id": "A-2-1",
//       "pid": "A-2"
//   },
//   {
//       "id": "A-2-2",
//       "pid": "A-2"
//   },
//   {
//       "id": "A-2-3",
//       "pid": "A-2"
//   }
// ]
// Output: [
//   "/A",
//   "/A/A-1",
//   "/A/A-2",
//   "/A/A-3",
//   "/A/A-2/A-2-1",
//   "/A/A-2/A-2-2",
//   "/A/A-2/A-2-3",
// ]

type Node struct {
	ID  string
	PID string
}

func main() {
	nodes := []Node{
		{
			"A",
			"-1",
		},
		{
			"A-1",
			"A",
		},
		{
			"A-2",
			"A",
		},
		{
			"A-3",
			"A",
		},
		{
			"A-2-1",
			"A-2",
		},
		{
			"A-2-2",
			"A-2",
		},
		{
			"A-2-3",
			"A-2",
		},
	}

	str := make(map[string]string, 0)
	str2 := make([]string, 0)

	for _, node := range nodes {
		if _, ok := str[node.PID]; ok {
			str[node.ID] = str[node.PID] + "/" + node.ID
			str2 = append(str2, str[node.PID]+"/"+node.ID)
		} else {
			str[node.ID] = node.ID
			str2 = append(str2, str[node.PID]+"/"+node.ID)
		}
	}

	fmt.Println(nodes)
	for _, s2 := range str2 {
		fmt.Printf("%v\n", s2)
	}
	//fmt.Printf("%v", str)
}
