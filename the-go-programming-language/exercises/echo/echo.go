package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// items := strings.Join(os.Args, " ")
	for idx, item := range os.Args {
		s += fmt.Sprintf("%s%d:%s\n", sep, idx, item)
		sep = " "
	}
	fmt.Println(s)
}
