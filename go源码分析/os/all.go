package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid(), "Getpid")
	fmt.Println("Getegid", os.Getegid())
	fmt.Println("Geteuid", os.Geteuid())
	fmt.Println(os.Getgroups())

	fmt.Print("Getwd")
	fmt.Println(os.Getwd())
	fmt.Println("Executable", os.Executable())
	fmt.Println("Hostname", os.Hostname())
}
