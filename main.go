package main

import (
	"fmt"
	"strings"
)

func vuln(x string) string {
	xSlice := strings.Split(x, ",")
	return xSlice[len(xSlice)]
}

func static() {
	test := [5]string{"1", "2", "3", "4", "6"}
	var x int
	fmt.Scan(&x) //Untrusted Input

	println(test[x])

	if x < 5 {
		println(test[x])
	}
	if x <= 5 {
		println(test[x])
	}
	if x < 6 {
		println(test[x])
	}
	if x < 100 {
		println(test[x])
	}
	if x < 0 {
		println(test[x])
	}
}
func new() {
	var x = "A,B,CCC,DDDDDDDDDDD,,EE,"
	xx := strings.Split(x, ",")
	fmt.Println(xx[0], xx[1], xx[2])
	for i := 0; i <= len(xx); i++ {
		fmt.Println(xx[i])
	}
	var z = [1]int{1}
	var INT int = 1
	if INT <= len(z) {
		println(z[INT])
	}

	println("-----")
}
func main() {
	new()
	var target = "A,B,CCC,DDDDDDDDDDD,,EE,"
	targetSlice := strings.Split(target, ",")
	for i := 0; i <= len(targetSlice); i++ {
		println("[" + targetSlice[i] + "]")
	}
	result := vuln(target)
	println(result)
	for i := 0; i != len(targetSlice); i++ {
		println("[" + targetSlice[i] + "]")
	}
	static()
	for i := 0; i == len(targetSlice); i++ {
		println("[" + targetSlice[i] + "]")
	}
	for i := 0; i < len(targetSlice)+4; i++ {
		println("[" + targetSlice[i] + "]")
	}
	for i := 0; i <= len(targetSlice)+4; i++ {
		println("[" + targetSlice[i] + "]")
	}
	for i := 0; i <= len(targetSlice)-1; i++ {
		println("[" + targetSlice[i] + "]")
	}
	for i := 0; i < len(targetSlice)+1; i++ {
		println("[" + targetSlice[i] + "]")
	}
	for i := 0; i < len(targetSlice)-1; i++ {
		println("[" + targetSlice[i] + "]")
	}
}
