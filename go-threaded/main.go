package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const numWorkers = 8

func checkPart(bs []byte, valueCh chan int) {
	n := len(bs)
	v := 0
	for i := 0; i < n; i += 4 {
		if bs[i+0] == 0xFF && bs[i+1] == 0xFF &&
			bs[i+2] == 0xFF && bs[i+3] == 0xFF {
			v += 1
		}
	}
	valueCh <- v
}

func main() {
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	n := len(bs)
	if n%4 != 0 {
		os.Exit(2)
	}
	np := n / 4
	cs := (np + numWorkers - 1) / numWorkers
	vch := make(chan int)
	for i := 0; i < numWorkers; i++ {
		l := i * (cs * 4)
		h := (i + 1) * (cs * 4)
		if h > n {
			h = n
		}
		go checkPart(bs[l:h], vch)
	}

	acc := 0
	for i := 0; i < numWorkers; i++ {
		acc += <-vch
	}
	fmt.Printf("%d/%d\n", acc, n/4)
}
