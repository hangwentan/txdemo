package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ErrorRecord struct {
	Name  string
	Rows  string
	Count int
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	dict := make(map[string]int)
	arr := make([]ErrorRecord, 0)
	for in.Scan() {
		text := strings.Split(in.Text(), " ")
		path := strings.Split(text[0], "\\")
		name := path[len(path)-1]
		if len(name) > 16 {
			name = name[len(name)-16:]
		}
		if v, ok := dict[name+text[1]]; ok {
			arr[v].Count += 1
		} else {
			dict[name+text[1]] = len(arr)
			arr = append(arr, ErrorRecord{
				Name:  name,
				Rows:  text[1],
				Count: 1,
			})
		}
	}
	if len(arr) > 8 {
		arr = arr[len(arr)-8:]
	}
	for i := range arr {
		fmt.Println(arr[i].Name + " " + arr[i].Rows + " " + strconv.Itoa(arr[i].Count))
	}
}
