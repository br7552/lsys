package lsystem

import "strings"

type Lsystem struct {
	replacer *strings.Replacer
	current  string
}

func NewLsystem(axiom string, rules ...string) *Lsystem {
	return &Lsystem{
		replacer: strings.NewReplacer(rules...),
		current:  axiom,
	}
}

func (ls *Lsystem) String() string {
	return ls.current
}

func (ls *Lsystem) Grow() {
	ls.current = ls.replacer.Replace(ls.current)
}
