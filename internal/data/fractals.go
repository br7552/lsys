package data

import (
	"github.com/br7552/lsys/internal/asciiturtle"
	"github.com/br7552/lsys/lsystem"
)

type Fractal struct {
	Axiom      string
	Rules      map[string]string
	Depth      int
	Angle      float64
	StartAngle float64
	Step       int
	Width      int
	Height     int
	Data       string
}

type state struct {
	x       int
	y       int
	heading float64
}

func Generate(f *Fractal) error {
	var rules []string

	for k, v := range f.Rules {
		rules = append(rules, k)
		rules = append(rules, v)
	}

	l := lsystem.New(f.Axiom, rules...)

	for i := 0; i < f.Depth; i++ {
		l.Grow()
	}

	canvas := asciiturtle.NewCanvas(f.Width, f.Height)
	pen, _ := asciiturtle.NewPen(canvas, f.StartAngle, canvas.Width()/2,
		canvas.Height()/2)

	var stack []state

	for _, v := range l.String() {
		switch v {
		case 'F':
			pen.Forward(f.Step)
		case 'G':
			pen.PenUp()
			pen.Forward(f.Step)
			pen.PenDown()
		case '+':
			pen.Right(f.Angle)
		case '-':
			pen.Left(f.Angle)
		case '[':
			stack = append(stack, state{pen.X, pen.Y, pen.GetHeading()})
		case ']':
			if len(stack) == 0 {
				panic("stack underflow")
			}

			state := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			pen.X = state.x
			pen.Y = state.y
			pen.SetHeading(state.heading)
		case '|':
			pen.Forward(2)
		}
	}

	f.Data = canvas.String()
	return nil
}
