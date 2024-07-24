package main

import "fmt"

func main() {
	var n1, n2, n3, n4 int
	fmt.Scanf("%d.%d.%d.%d", &n1, &n2, &n3, &n4)

	fmt.Println(n1<<24 + n2<<16 + n3<<8 + n4)

	var n5 int
	fmt.Scanln(&n5)

	fmt.Println(fmt.Sprintf("%d", n5>>24) + "." + fmt.Sprintf("%d", (n5>>16)&0xff) + "." + fmt.Sprintf("%d", (n5>>8)&0xff) + "." + fmt.Sprintf("%d", n5&0xff))
}
