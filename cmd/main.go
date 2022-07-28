package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	go_qwe()
}

func go_qwe() {
	fmt.Println("qwe " + quote.Go())
}
