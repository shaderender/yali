package main

import (
	"ilox/lox"
	"os"
)

func main() {
	lx := lox.New()
	lx.Run(os.Args[1:])
}
