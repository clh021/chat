package main

import (
	"fmt"

	"github.com/clh021/chat"
)

var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)
	chat.Run()
}
