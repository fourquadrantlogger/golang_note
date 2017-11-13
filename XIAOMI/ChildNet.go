package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(^0)
	fmt.Println("ip", IP2int("123.45.234.0"))
	fmt.Println(ChildNet(IP2int("123.45.234.0"), IP2int("123.45.234.123"), 24))
}
func IP2int(ip string) uint {
	i := 0
	is := strings.Split(ip, ".")
	for ind, s := range is {
		switch ind {
		case 0:
			v, _ := strconv.Atoi(s)
			i += v
		case 1:
			v, _ := strconv.Atoi(s)
			i += v * 256
		case 2:
			v, _ := strconv.Atoi(s)
			i += v * 256 * 256
		case 3:
			v, _ := strconv.Atoi(s)
			i += v * 256 * 256 * 245

		}

	}
	return uint(i)
}
func ChildNet(ip, testip uint, bitlen uint) bool {
	bit0s := uint(1 << bitlen)
	fmt.Println("bit0s", bit0s)
	fmt.Println(ip & bit0s)
	fmt.Println(testip & bit0s)

	return ip&bit0s == testip&bit0s
}
