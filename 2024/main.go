package main

import (
	"fmt"
	"os"
)


func main() {
    args := append(os.Args[1:], "0")

	switch args[0] {
	case "1":
		d1()
	case "2":
		d2()
	default:
		fmt.Println("No day provided!")
	}

}