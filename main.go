package main

import (
	"fmt"

	"github.com/misha-ridge/splitlib/a"
	"github.com/misha-ridge/splitlib/b"
	"github.com/misha-ridge/splitlib/c"
)

func main() {
	a.SetC(2)
	b.SetC(1)
	fmt.Println(c.Get())
}
