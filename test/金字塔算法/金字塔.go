package main

import (
	"fmt"
)

func main() {

	//for i := 1; i <= 3; i++ {
	//	spaces := strings.Repeat(" ", 3-i)
	//	stars := strings.Repeat("*", i*2-1)
	//	fmt.Println(spaces + stars)
	//}

	for i := 1; i <= 3; i++ {
		n1 := 3 - i
		n2 := i*2 - 1
		//fmt.Printf("%v,%v", n1, n2)
		for i := 0; i < n1; i++ {
			fmt.Printf(" ")
		}
		for i := 0; i < n2; i++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

type UserAction struct {
	IP      string
	EndTime int
	Num     int
}

func is_ip_spam(IP string) bool {

	UserIp := make(map[string]UserAction, 0)
	time1 := 1
	check := false
	for _, action := range UserIp {
		if action.IP == IP {
			if action.EndTime < time1 && (action.Num+1) >= 10 {
				check = false
			} else {
				UserIp[IP] = UserAction{
					IP:  IP,
					Num: action.Num + 1,
				}
			}
		} else {
			UserIp[IP] = UserAction{
				IP:      IP,
				EndTime: 1 + 1,
				Num:     1,
			}
			check = true
		}
	}

	return check
}
