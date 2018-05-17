package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	n := len(bs)
	if n%4 != 0 {
		os.Exit(2)
	}

	acc := 0
	for i := 0; i < n; i += 4 {
		if bs[i+0] == 0xFF && bs[i+1] == 0xFF &&
			bs[i+2] == 0xFF && bs[i+3] == 0xFF {
			acc += 1
		}
	}
	fmt.Printf("%d/%d\n", acc, n/4)
}
