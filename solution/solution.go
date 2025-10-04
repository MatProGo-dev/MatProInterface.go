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
