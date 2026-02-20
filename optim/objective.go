package optim

// Objective represents an optimization objective given an expression and
// objective sense (maximize or minimize).
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type Objective struct {
	ScalarExpression
	Sense ObjSense
}

// NewObjective returns a new optimization objective given an expression and
// objective sense.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func NewObjective(e ScalarExpression, sense ObjSense) *Objective {
	return &Objective{e, sense}
}

// ObjSense represents whether an optimization objective is to be maximized or
// minimized. This implementation conforms to the Gurobi encoding.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type ObjSense int

// Objective senses (minimize and maximize) encoding using Gurobi's standard.
const (
	// SenseMinimize indicates that the objective should be minimized.
	//
	// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
	SenseMinimize ObjSense = 1
	// SenseMaximize indicates that the objective should be maximized.
	//
	// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
	SenseMaximize = -1
)
