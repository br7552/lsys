package main

import (
	"fmt"

	"github.com/br7552/lsystem/lsys"
)

func main() {
	l := lsys.NewLsystem("F", "F", "F-F++F-F")

	for i := 0; i < 3; i++ {
		l.Grow()
		fmt.Println(l)
	}
}
