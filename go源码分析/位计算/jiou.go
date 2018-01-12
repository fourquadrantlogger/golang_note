package main

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			println(i)
		}
		if i&0x1 == 0 {
			println(i)
		}
	}
}
