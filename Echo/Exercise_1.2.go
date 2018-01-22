package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:]{
		s += strconv.Itoa(i) + ":" + arg + sep
		sep = " "
	}
	fmt.Println(s)
}
