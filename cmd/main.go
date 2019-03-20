package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/namnhce/weighted-random-messages/random"
	"github.com/namnhce/weighted-random-messages/random/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing file path")
		return
	}

	filePath := os.Args[1]
	ext := filepath.Ext(filePath)

	if ext != ".txt" {
		panic("invalid file extension, .txt is required")
	}
	input, err := util.ReadData(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(random.F(input))
}
