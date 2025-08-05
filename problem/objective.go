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

/*
IsLinear
Description:

	This method returns true if the objective is linear, false otherwise.
*/
func (o *Objective) IsLinear() bool {
	return symbolic.IsLinear(o.Expression)
}

/*
SubstituteAccordingTo
Description:

	Substitutes the variables in the objective according to the replacement map.
*/
func (o *Objective) SubstituteAccordingTo(replacementMap map[symbolic.Variable]symbolic.Expression) *Objective {
	newExpression := o.Expression.SubstituteAccordingTo(replacementMap)
	return &Objective{newExpression, o.Sense}
}
