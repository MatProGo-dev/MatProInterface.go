package solution

import (
	"fmt"

	solution_status "github.com/MatProGo-dev/MatProInterface.go/solution/status"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

const (
	tinyNum float64 = 0.01
)

// Solution stores the solution of an optimization problem and associated
// metadata
type Solution interface {
	GetOptimalValue() float64
	GetValueMap() map[uint64]float64

	// GetStatus
	//
	GetStatus() solution_status.SolutionStatus
}

func ExtractValueOfVariableWithID(s Solution, idx uint64) (float64, error) {
	val, ok := s.GetValueMap()[idx]
	if !ok {
		return 0.0, fmt.Errorf(
			"The idx \"%v\" was not in the variable map for the solution.",
			idx,
		)
	}
	return val, nil
}

func ExtractValueOfVariable(s Solution, v symbolic.Variable) (float64, error) {
	idx := v.ID // Extract index of v
	return ExtractValueOfVariableWithID(s, idx)
}

// FindValueOfExpression evaluates a symbolic expression using the values from a solution.
// It substitutes all variables in the expression with their values from the solution
// and returns the resulting scalar value.
func FindValueOfExpression(s Solution, expr symbolic.Expression) (float64, error) {
	// Get all variables in the expression
	vars := expr.Variables()

	// Create a substitution map from variables to their constant values
	subMap := make(map[symbolic.Variable]symbolic.Expression)
	for _, v := range vars {
		val, err := ExtractValueOfVariable(s, v)
		if err != nil {
			return 0.0, fmt.Errorf(
				"failed to extract value for variable %v: %w",
				v.ID,
				err,
			)
		}
		subMap[v] = symbolic.K(val)
	}

	// Substitute all variables with their values
	resultExpr := expr.SubstituteAccordingTo(subMap)

	// Type assert to K (constant) to extract the float64 value
	resultK, ok := resultExpr.(symbolic.K)
	if !ok {
		return 0.0, fmt.Errorf(
			"expected substituted expression to be a constant, got type %T",
			resultExpr,
		)
	}

	return float64(resultK), nil
}
