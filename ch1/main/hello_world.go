package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		fmt.Println("Hello World", os.Args[1])
	}
}
