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
	fmt.Println(getTeamsGroup(groupList))
}

// generateTeams 生成成员队伍列表

func getTeamsGroup(groupList [][]string) [][]string {

	// 设置随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// TODO 计算数据组数量
	GroupsNum := len(groupList)

	// TODO 统计每组人数
	var GroupsUserNums []int
	for _, value := range groupList {
		GroupsUserNums = append(GroupsUserNums, len(value))
	}

	// TODO 统计用户总数
	userNums := 0
	for num := range GroupsUserNums {
		userNums += num
	}

	// TODO 初始化结果数据
	teams := make([][]string, userNums/2)

	// TODO 初始化组标记数组
	usedGroups := make([][]bool, 0, GroupsNum)
	for i := 0; i < GroupsNum; i++ {
		usedGroups = append(usedGroups, make([]bool, GroupsUserNums[i]))
	}

	count := 0
	var groupIndex = 0
	for i := 0; i < len(teams); i++ {

		var team []string
		num := r.Intn(2) + 2

		for len(team) < num && userNums > count {

			// 随机选取一个小组
			groupIndex = r.Intn(GroupsNum)

			// 从该小组中随机取出一人
			userIndex := r.Intn(len(groupList[groupIndex]))

			// 若该小组已使用跳过
			if usedGroups[groupIndex][userIndex] {
				continue
			}

			// 获取该用户
			user := groupList[groupIndex][userIndex]

			// 检测是否已选队员来自同一组
			if check(groupList[groupIndex], team) {
				continue
			}

			// 将用户加入队伍中
			team = append(team, user)

			// 标记该小组已使用

			usedGroups[groupIndex][userIndex] = false
			count++
		}
		teams[i] = team
	}

	// 删除最后为nil的结果
	for len(teams[len(teams)-1]) == 0 {
		teams = teams[:len(teams)-1]
	}

	//最后等于1,需要合并到两的队伍里面
	if len(teams[len(teams)-1]) == 1 {
		temp := teams[len(teams)-1]
		teams = teams[:len(teams)-1]
		for i := range teams {
			if len(teams[i]) == 2 {

				if !check(groupList[groupIndex], teams[i]) {
					teams[i] = append(teams[i], temp...)
					break
				}
			}
		}
	}
	return teams
}

func check(slice []string, target []string) bool {
	for _, value := range target {

		for _, v := range slice {
			if v == value {
				return true
			}
		}

	}
	return false
}
