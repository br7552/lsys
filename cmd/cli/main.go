package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"

	"github.com/br7552/asciiturtle"
	"github.com/br7552/lsys/lsystem"
)

const canvasSize = 50

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

	l := lsystem.New(axiom, rules...)

	for i := 0; i < depth; i++ {
		l.Grow()
	}

	canvas := asciiturtle.NewCanvas(canvasSize, canvasSize)
	pen, _ := asciiturtle.NewPen(canvas, '#', canvasSize/2, canvasSize/2)

	for _, v := range l.String() {
		switch v {
		case 'F':
			pen.Forward(3)
		case 'G':
			pen.PenUp()
			pen.Forward(3)
			pen.PenDown()
		case '+':
			pen.Right(60)
		case '-':
			pen.Left(60)
		}
	}

	fmt.Println(canvas)
}
