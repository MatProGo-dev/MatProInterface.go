package solver

import (
	op "github.com/MatProGo-dev/MatProInterface.go/problem"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
solver.go
Description:
	Defines the new interface Solver which should define
*/

type Solver interface {
	ShowLog(tf bool) error
	SetTimeLimit(timeLimit float64) error
	AddVariable(varIn symbolic.Variable) error
	AddVariables(varSlice []symbolic.Variable) error
	AddConstraint(constrIn symbolic.Constraint) error
	SetObjective(objectiveIn op.Objective) error
	Optimize() (op.Solution, error)
	DeleteSolver() error
}
