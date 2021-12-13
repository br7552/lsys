package data

import (
	"github.com/br7552/lsys/internal/asciiturtle"
	"github.com/br7552/lsys/internal/validator"
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

func ValidateFractal(v *validator.Validator, fractal *Fractal) {
	v.Check(fractal.Axiom != "", "axiom", "must be provided")
	v.Check(len(fractal.Axiom) <= 500, "axiom",
		"must not be longer than 500 characters")

	v.Check(len(fractal.Rules) > 0, "rules", "must be provided")

	v.Check(fractal.Depth >= 0 && fractal.Depth <= 10, "depth",
		"must be between 0 and 10 (inclusive)")

	v.Check(fractal.Angle >= 0.0 && fractal.Angle <= 360.0, "angle",
		"must be between 0 and 360 (inclusive)")

	v.Check(fractal.StartAngle >= 0.0 && fractal.Angle <= 360.0, "angle",
		"must be between 0 and 360 (inclusive)")

	v.Check(fractal.Step >= 1 && fractal.Step <= 500, "step",
		"must be between 1 and 500 (inclusive)")

	v.Check(fractal.Width >= 1 && fractal.Step <= 500, "width",
		"must be between 1 and 500 (inclusive)")

	v.Check(fractal.Height >= 1 && fractal.Step <= 500, "height",
		"must be between 1 and 500 (inclusive)")
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
