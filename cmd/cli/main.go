package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/br7552/lsys/internal/asciiturtle"
	"github.com/br7552/lsys/lsystem"
)

const (
	depthDefault        = 3
	angleDefault        = 45
	stepDefault         = 3
	canvasHeightDefault = 50
	canvasWidthDefault  = 50
)

var (
	buildTime string
	version   string
)

func main() {
	var (
		axiom        string
		rules        []string
		depth        int
		startAngle   float64
		angle        float64
		step         int
		canvasHeight int
		canvasWidth  int
	)

	flag.StringVar(&axiom, "axiom", "", "Lsystem axiom")
	flag.Func("rules", "Lsystem production rules", func(s string) error {
		rules = strings.FieldsFunc(s, func(c rune) bool {
			return unicode.IsSpace(c) || c == '='
		})
		return nil
	})
	flag.IntVar(&depth, "depth", depthDefault, "Number of iterations")
	flag.Float64Var(&startAngle, "start-angle", 0, "Initial angle")
	flag.Float64Var(&angle, "angle", angleDefault, "Rotation angle")
	flag.IntVar(&step, "step", stepDefault, "Step Size")
	printVersion := flag.Bool("version", false, "Print version")
	flag.IntVar(&canvasWidth, "width", canvasWidthDefault, "Plot width")
	flag.IntVar(&canvasHeight, "height", canvasHeightDefault, "Plot height")

	flag.Parse()

	if *printVersion {
		fmt.Printf("Version:\t%s\nBuild time:\t%s\n", version, buildTime)
		os.Exit(0)
	}

	if axiom == "" || rules == nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	l := lsystem.New(axiom, rules...)

	for i := 0; i < depth; i++ {
		l.Grow()
	}

	canvas := asciiturtle.NewCanvas(canvasWidth, canvasHeight)
	pen, _ := asciiturtle.NewPen(canvas, startAngle, canvas.Width()/2,
		canvas.Height()/2)

	for _, v := range l.String() {
		switch v {
		case 'F':
			pen.Forward(step)
		case 'G':
			pen.PenUp()
			pen.Forward(step)
			pen.PenDown()
		case '+':
			pen.Right(angle)
		case '-':
			pen.Left(angle)
		case '[':
			pushState(pen.X, pen.Y, pen.GetHeading())
		case ']':
			var heading float64
			pen.X, pen.Y, heading = popState()
			pen.SetHeading(heading)
		case '|':
			pen.Forward(2)
		}
	}

	fmt.Println(canvas)
}

type state struct {
	x       int
	y       int
	heading float64
}

var stack []state

func pushState(x int, y int, heading float64) {
	stack = append(stack, state{x, y, heading})
}

func popState() (int, int, float64) {
	if len(stack) == 0 {
		panic("stack underflow")
	}

	state := stack[len(stack)-1]

	stack = stack[:len(stack)-1]

	return state.x, state.y, state.heading
}
