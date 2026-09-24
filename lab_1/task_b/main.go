package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(string(a[len(a)-1]))
}
