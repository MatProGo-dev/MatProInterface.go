package problem

import "github.com/MatProGo-dev/SymbolicMath.go/symbolic"

// Objective represents an optimization objective given an expression and
// objective sense (maximize or minimize).
type Objective struct {
	symbolic.Expression
	Sense ObjSense
}

// NewObjective returns a new optimization objective given an expression and
// objective sense
func NewObjective(e symbolic.Expression, sense ObjSense) *Objective {
	return &Objective{e, sense}
}
