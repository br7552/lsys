package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/br7552/lsys/internal/data"
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
	var fractal data.Fractal

	flag.StringVar(&fractal.Axiom, "axiom", "", "Lsystem axiom")
	flag.Func("rules", "Lsystem production rules", func(s string) error {
		rules := make(map[string]string)

		for _, v := range strings.Fields(s) {
			parts := strings.Split(v, "=")
			if len(parts) != 2 {
				return errors.New("rules must have the form LHS=RHS")
			}
			rules[parts[0]] = parts[1]
		}

		fractal.Rules = rules
		return nil
	})
	flag.IntVar(&fractal.Depth, "depth", depthDefault, "Number of iterations")
	flag.Float64Var(&fractal.StartAngle, "start-angle", 0, "Initial angle")
	flag.Float64Var(&fractal.Angle, "angle", angleDefault, "Rotation angle")
	flag.IntVar(&fractal.Step, "step", stepDefault, "Step Size")
	printVersion := flag.Bool("version", false, "Print version")
	flag.IntVar(&fractal.Width, "width", canvasWidthDefault, "Plot width")
	flag.IntVar(&fractal.Height, "height", canvasHeightDefault, "Plot height")

	flag.Parse()

	if *printVersion {
		fmt.Printf("Version:\t%s\nBuild time:\t%s\n", version, buildTime)
		os.Exit(0)
	}

	data.Generate(&fractal)

	fmt.Println(fractal.Data)
}
