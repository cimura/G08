package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if piscine.Comcheck(os.Args) {
		fmt.Println("Alert!!!")
	}
}
