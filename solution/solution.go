package solution

import (
	"fmt"

	"github.com/MatProGo-dev/MatProInterface.go/problem"
	solution_status "github.com/MatProGo-dev/MatProInterface.go/solution/status"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

const (
	tinyNum float64 = 0.01
)

// Solution stores the solution of an optimization problem and associated
// metadata
type Solution interface {
	// GetValueMap returns a map from variable ID to its value in the solution
	//
	// The map keys are the IDs of the variables (uint64)
	// The map values are the corresponding values of those variables in the solution (float64)
	//
	// Example:
	// If variable with ID 1 has value 3.5 in the solution, then the map will contain an entry:
	//   1: 3.5
	//
	// This allows easy lookup of variable values by their IDs.
	//
	// Note: Variable IDs can be obtained from the symbolic.Variable.ID field.
	//
	// Example usage:
	//   valMap := solution.GetValueMap()
	//   x1Value := valMap[x1.ID] // where x1 is a symbolic.Variable
	//   fmt.Println("Value of x1 in solution:", x1Value)
	GetValueMap() map[uint64]float64

	// GetStatus returns the status of the solution (e.g., optimal, infeasible, etc.).
	//
	// The returned value is of type solution_status.SolutionStatus, which indicates
	// whether the solution is optimal, infeasible, unbounded, or has another status.
	//
	// Example usage:
	//   status := solution.GetStatus()
	//   fmt.Println("Solution status:", status)
	GetStatus() solution_status.SolutionStatus

	// GetProblem returns the optimization problem that this solution is for
	GetProblem() *problem.OptimizationProblem
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
// and returns the resulting symbolic expression (typically a constant).
func FindValueOfExpression(s Solution, expr symbolic.Expression) (symbolic.Expression, error) {
	// Get all variables in the expression
	vars := expr.Variables()

	// Create a substitution map from variables to their constant values
	subMap := make(map[symbolic.Variable]symbolic.Expression)
	for _, v := range vars {
		val, err := ExtractValueOfVariable(s, v)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to extract value for variable %v: %w",
				v.ID,
				err,
			)
		}
		subMap[v] = symbolic.K(val)
	}

	// Substitute all variables with their values
	resultExpr := expr.SubstituteAccordingTo(subMap)

	return resultExpr, nil
}

// GetOptimalObjectiveValue evaluates the objective function of an optimization problem
// at the solution point. It uses the FindValueOfExpression function to compute the value
// of the objective expression using the variable values from the solution.
func GetOptimalObjectiveValue(sol Solution) (float64, error) {
	// Get the problem from the solution
	prob := sol.GetProblem()
	if prob == nil {
		return 0.0, fmt.Errorf("solution does not have an associated problem")
	}

	// Get the objective expression from the problem
	objectiveExpr := prob.Objective.Expression
	if objectiveExpr == nil {
		return 0.0, fmt.Errorf("problem does not have a defined objective")
	}

	// Use FindValueOfExpression to evaluate the objective at the solution point
	resultExpr, err := FindValueOfExpression(sol, objectiveExpr)
	if err != nil {
		return 0.0, fmt.Errorf("failed to evaluate objective expression: %w", err)
	}

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
