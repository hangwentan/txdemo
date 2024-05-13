package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	groupList := [][]string{
		{"少华", "少平", "少军", "少安", "少康"},
		{"福军", "福堂", "福民", "福平", "福心"},
		{"小明", "小红", "小花", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
	}
	fmt.Println(getList(groupList))
}

func getList(groupList [][]string) [][]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 计算总人数和每组人数
	numGroups := len(groupList) //组条数
	var numPeoplePerGroup []int //每组人数

	for _, group := range groupList {
		numPeoplePerGroup = append(numPeoplePerGroup, len(group))
	}

	totalPeople := 0
	for _, num := range numPeoplePerGroup {
		totalPeople += num
	}

	// 声明结果数组
	result := make([][]string, totalPeople/2)

	// 声明标记数组
	var usedGroups = make([][]bool, 0, numGroups)

	for i := 0; i < numGroups; i++ {
		usedGroups = append(usedGroups, make([]bool, numPeoplePerGroup[i]))
	}
	counter := 0
	var groupIndex = 0 //记录最后一个的索引
	// 随机分组
	for i := 0; i < len(result); i++ {
		var team []string

		// 随机选取2人或3人组队 这里可以把所有的随机先随机出来，然后判断是否2、3 都大于两个
		numPeople := r.Intn(2) + 2

		for len(team) < numPeople && counter < totalPeople {
			// 随机选取一个小组
			//for groupIndex := 0; groupIndex < numGroups; groupIndex++ {
			groupIndex = r.Intn(numGroups)
			// 从该小组中随机选取一个人
			personIndex := r.Intn(len(groupList[groupIndex]))
			// 如果该小组已被使用过，则跳过
			if usedGroups[groupIndex][personIndex] {
				continue
			}

			person := groupList[groupIndex][personIndex]

			// 检查是否与已选队员来自同一小组
			if sliceContains(groupList[groupIndex], team) {
				continue
			}

			// 将该人加入队伍
			team = append(team, person)
			counter++
			// 标记该小组已使用
			usedGroups[groupIndex][personIndex] = true
			//}

		}
		// 将队伍加入结果数组
		result[i] = team
	}

	//删除最后为nil 的结果
	for len(result[len(result)-1]) == 0 {
		result = result[:len(result)-1]
	}

	//最后等于1,需要合并到两的队伍里面
	if len(result[len(result)-1]) == 1 {
		temp := result[len(result)-1]
		result = result[:len(result)-1]
		for i := range result {
			if len(result[i]) == 2 {
				if !sliceContains(groupList[groupIndex], result[i]) {
					result[i] = append(result[i], temp...)
					break
				}
			}
		}
	}
	return result
}

func sliceContains(slice []string, target []string) bool {
	for _, value := range target {

		for _, v := range slice {
			if v == value {
				return true
			}
		}

	}
	return false
}
