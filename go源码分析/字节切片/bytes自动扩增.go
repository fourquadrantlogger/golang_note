package main

import (
	"io/ioutil"
	"os"
)

func main() {
	var bs []byte = make([]byte, 1024)
	ioutil.WriteFile("bs.iso", bs, os.ModePerm)

	var toread []byte = make([]byte, 0)

}
