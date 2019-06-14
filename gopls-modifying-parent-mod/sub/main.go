package main

import (
	"fmt"

	"github.com/kylelemons/godebug/diff"
)

func main() {
	fmt.Println(diff.DiffChunks("a", "b"))
}
