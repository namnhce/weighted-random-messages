package main

import (
	"fmt"
	"os"

	"github.com/namnhce/weighted-random-messege/random"
	"github.com/namnhce/weighted-random-messege/random/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing file path")
		return
	}

	filePath := os.Args[1]
	input, err := util.ReadData(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(random.F(input))
}
