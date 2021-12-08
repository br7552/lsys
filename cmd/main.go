package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"

	"github.com/br7552/lsystem/lsys"
)

func main() {
	var (
		depth int
		axiom string
		rules []string
	)

	flag.IntVar(&depth, "depth", 3, "Number of iterations")
	flag.StringVar(&axiom, "axiom", "", "Lsystem axiom")
	flag.Func("rules", "Lsystem production rules", func(s string) error {
		rules = strings.FieldsFunc(s, func(c rune) bool {
			return unicode.IsSpace(c) || c == '='
		})
		return nil
	})

	flag.Parse()

	l := lsys.NewLsystem(axiom, rules...)

	for i := 0; i < depth; i++ {
		l.Grow()
		fmt.Println(l)
	}
}
