package problem

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
)

/*
GetExampleProblem3
Description:

	Returns the LP from this youtube video:
		https://www.youtube.com/watch?v=QAR8zthQypc&t=483s
	It should look like this:
		Maximize	4 x1 + 3 x2 + 5 x3
		Subject to
			x1 + 2 x2 + 2 x3 <= 4
			3 x1 + 4 x3 <= 6
			2 x1 + x2 + 4 x3 <= 8
			x1 >= -1.0
			x2 >= -1.0
			x3 >= -1.0
*/
func GetExampleProblem3() *OptimizationProblem {
	// Setup
	out := NewProblem("TestProblem3")

	// Create variables
	x := out.AddVariableVectorClassic(
		3,
		-1.0,
		symbolic.Infinity.Constant(),
		symbolic.Continuous,
	)

	// Create Basic Objective
	c := getKVector.From([]float64{4.0, 3.0, 5.0})
	out.SetObjective(
		c.Transpose().Multiply(x),
		SenseMaximize,
	)

	// Create Constraints (using one big matrix)
	A := getKMatrix.From([][]float64{
		{1.0, 2.0, 2.0},
		{3.0, 0.0, 4.0},
		{2.0, 1.0, 4.0},
	})
	b := getKVector.From([]float64{4.0, 6.0, 8.0})
	out.Constraints = append(out.Constraints, A.Multiply(x).LessEq(b))

	// TODO(kwesi): Figure out how to add non-negativity constraints
	// // Add non-negativity constraints
	// for _, varII := range x {
	// 	out.Constraints = append(
	// 		out.Constraints,
	// 		varII.GreaterEq(0.0),
	// 	)
	// }
	return out
}

/*
GetExampleProblem4
Description:

	Returns the LP from this youtube video:
		https://www.youtube.com/watch?v=QAR8zthQypc&t=483s
	It should look like this:
		Maximize	4 x1 + 3 x2 + 5 x3
		Subject to
			x1 + 2 x2 + 2 x3 <= 4
			3 x1 + 4 x3 <= 6
			2 x1 + x2 + 4 x3 <= 8
			x1 >= 0
			x2 >= 0
			x3 >= 0
*/
func GetExampleProblem4() *OptimizationProblem {
	// Setup
	out := NewProblem("TestProblem3")

	// Create variables
	x := out.AddVariableVectorClassic(
		3,
		0.0, // Use this line to implement non-negativity constraints
		symbolic.Infinity.Constant(),
		symbolic.Continuous,
	)

	// Create Basic Objective
	c := getKVector.From([]float64{4.0, 3.0, 5.0})
	out.SetObjective(
		c.Transpose().Multiply(x),
		SenseMaximize,
	)

	// Create Constraints (using one big matrix)
	A := getKMatrix.From([][]float64{
		{1.0, 2.0, 2.0},
		{3.0, 0.0, 4.0},
		{2.0, 1.0, 4.0},
	})
	b := getKVector.From([]float64{4.0, 6.0, 8.0})
	out.Constraints = append(out.Constraints, A.Multiply(x).LessEq(b))

	return out
}

/*
GetExampleProblem3
Description:

	Returns the LP from this youtube video:
		https://www.youtube.com/watch?v=QAR8zthQypc&t=483s
	It should look like this:
		Maximize	4 x1 + 3 x2 + 5 x3
		Subject to
			x1 + 2 x2 + 2 x3 <= 4
			3 x1 + 4 x3 <= 6
			2 x1 + x2 + 4 x3 <= 8
			x1 >= 0
			x2 >= 0
			x3 >= 0
*/
func GetExampleProblem5() *OptimizationProblem {
	// Setup
	out := NewProblem("TestProblem3")

	// Create variables
	x := out.AddVariableVector(3)

	// Create Basic Objective
	c := getKVector.From([]float64{4.0, 3.0, 5.0})
	out.SetObjective(
		c.Transpose().Multiply(x),
		SenseMaximize,
	)

	// Create Constraints (using one big matrix)
	A := getKMatrix.From([][]float64{
		{1.0, 2.0, 2.0},
		{3.0, 0.0, 4.0},
		{2.0, 1.0, 4.0},
	})
	b := getKVector.From([]float64{4.0, 6.0, 8.0})
	out.Constraints = append(out.Constraints, A.Multiply(x).LessEq(b))

	// Add non-negativity constraints
	for _, varII := range x {
		out.Constraints = append(
			out.Constraints,
			varII.GreaterEq(0.0),
		)
	}
	return out
}
