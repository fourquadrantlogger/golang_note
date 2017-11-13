package main

func main() {
	print(MIHOME(3, 4, []int{5, 6, 1}, []int{8, 2, 3, 5}))
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
func MIHOME(n, m int, mihome []int, mifans []int) (sum int) {
	mihome_fans := make(map[int][]int)
	for _, fan := range mifans {
		min_home := Abs(mihome[0] - fan)
		for _, home := range mihome[1:] {
			distance := Abs(home - fan)
			if min_home > distance {
				min_home = distance
			}
		}
		mihome_fans[min_home] = mihome_fans + distance
	}
	for k, v := range mihome_fans {
		sum += v
	}
	return
}
