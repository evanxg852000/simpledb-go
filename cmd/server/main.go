package main

import (
	"fmt"

	"github.com/evanxg852000/simpledb/internal/file"
)

func main() {
	fmt.Println("Welcome SimpleDB Server!")
	_ = file.NewBlockId("users.tbl", 32)
}
