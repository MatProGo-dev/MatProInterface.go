package problem_test

/*
optimization_problem_test.go
Description:

	Tests for all functions and objects defined in the optimization_problem.go file.
*/

import (
	"fmt"
	"strings"
	"testing"

	"github.com/MatProGo-dev/MatProInterface.go/causeOfProblemNonlinearity"
	"github.com/MatProGo-dev/MatProInterface.go/mpiErrors"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"github.com/MatProGo-dev/MatProInterface.go/problem"
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
)

/*
TestOptimizationProblem_NewProblem1
Description:

	Tests the NewProblem function with a simple name.
	Verifies that the name is set correctly and
	that zero variables and constraints exist in the fresh
	problem.
*/
func TestOptimizationProblem_NewProblem1(t *testing.T) {
	// Constants
	name := "TestProblem1"

	// New Problem
	p1 := problem.NewProblem(name)

	// Check that the name is as expected in the problem.
	if p1.Name != name {
		t.Errorf("expected the name of the problem to be %v; received %v",
			name, p1.Name)
	}

	// Check that the number of variables is zero.
	if len(p1.Variables) != 0 {
		t.Errorf("expected the number of variables to be 0; received %v",
			len(p1.Variables))
	}

	// Check that the number of constraints is zero.
	if len(p1.Constraints) != 0 {
		t.Errorf("expected the number of constraints to be 0; received %v",
			len(p1.Constraints))
	}
}

/*
TestOptimizationProblem_AddVariable1
Description:

	Tests the AddVariable function with a simple problem.
*/
func TestOptimizationProblem_AddVariable1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddVariable()

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Continuous {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Continuous, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddRealVariable1
Description:

	Tests the AddRealVariable function with a simple problem.
*/
func TestOptimizationProblem_AddRealVariable1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddRealVariable()

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Continuous {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Continuous, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddVariableClassic1
Description:

	Tests the AddVariableClassic function with a simple problem.
*/
func TestOptimizationProblem_AddVariableClassic1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddVariableClassic(0, 1, symbolic.Binary)

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Binary {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Binary, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddBinaryVariable1
Description:

	Tests the AddBinaryVariable function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariable1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddBinaryVariable()

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Binary {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Binary, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddVariableVector1
Description:

	Tests the AddVariableVector function with a simple problem.
*/
func TestOptimizationProblem_AddVariableVector1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	p1.AddVariableVector(dim)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Continuous, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddVariableVectorClassic1
Description:

	Tests the AddVariableVectorClassic function with a simple problem.
*/
func TestOptimizationProblem_AddVariableVectorClassic1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	p1.AddVariableVectorClassic(dim, 0, 1, symbolic.Binary)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Binary {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddBinaryVariableVector1
Description:

	Tests the AddBinaryVariableVector function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariableVector1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	p1.AddBinaryVariableVector(dim)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Binary {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddVariableMatrix1
Description:

	Tests the AddVariableMatrix function with a simple problem.
*/
func TestOptimizationProblem_AddVariableMatrix1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	rows := 5
	cols := 5

	// Algorithm
	p1.AddVariableMatrix(rows, cols, 0, 1, symbolic.Binary)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != rows*cols {
		t.Errorf("expected the number of variables to be %v; received %v",
			rows*cols, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddBinaryVariableMatrix1
Description:

	Tests the AddBinaryVariableMatrix function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariableMatrix1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	rows := 5
	cols := 5

	// Algorithm
	p1.AddBinaryVariableMatrix(rows, cols)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != rows*cols {
		t.Errorf("expected the number of variables to be %v; received %v",
			rows*cols, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_SetObjective1
Description:

	Tests the SetObjective function with a simple linear objective.
*/
func TestOptimizationProblem_SetObjective1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SetObjective1")
	v1 := p1.AddVariable()
	v2 := p1.AddVariable()

	// Algorithm
	err := p1.SetObjective(v1.Plus(v2), problem.SenseMaximize)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the objective is as expected.
	if p1.Objective.Sense != problem.SenseMaximize {
		t.Errorf("expected the sense of the objective to be %v; received %v",
			problem.SenseMaximize, p1.Objective.Sense)
	}
}

/*
TestOptimizationProblem_SetObjective2
Description:

	Tests the SetObjective function with a vector objective
	which should cause an error.
*/
func TestOptimizationProblem_SetObjective2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SetObjective2")
	v1 := p1.AddVariableVector(5)

	// Algorithm
	err := p1.SetObjective(v1, problem.SenseMaximize)
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			"trouble parsing input expression:",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_ToSymbolicConstraint1
Description:

	Tests the ToSymbolicConstraint function with a simple problem.
*/
func TestOptimizationProblem_ToSymbolicConstraint1(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := model1.AddVariable()
	v2 := model1.AddVariable()
	v3 := model1.AddVariable()

	// Algorithm
	sum, err := v1.Plus(v2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	constr1, err := sum.LessEq(v3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	constr1prime, err := problem.ToSymbolicConstraint(constr1)

	// Check that constr1prime is a VectorConstraint
	if _, ok := constr1prime.(symbolic.ScalarConstraint); !ok {
		t.Errorf("expected the type of constr1prime to be %T; received %T",
			symbolic.VectorConstraint{}, constr1prime)
	}

}

/*
TestOptimizationProblem_ToSymbolicConstraint2
Description:

	Tests the ToSymbolicConstraint function with a simple problem
	that has a vector constraint. This vector constraint
	will be a GreaterThanEqual vector constraint between
	a vector variable and a vector variable.
*/
func TestOptimizationProblem_ToSymbolicConstraint2(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := model1.AddVariableVector(5)
	v2 := model1.AddVariableVector(5)

	// Algorithm
	constr1, err := v1.GreaterEq(v2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	constr1prime, err := problem.ToSymbolicConstraint(constr1)

	// Check that constr1prime is a VectorConstraint
	if _, ok := constr1prime.(symbolic.VectorConstraint); !ok {
		t.Errorf("expected the type of constr1prime to be %T; received %T",
			symbolic.VectorConstraint{}, constr1prime)
	}
}

/*
TestOptimizationProblem_ToSymbolicConstraint3
Description:

	Tests the ToSymbolicConstraint function with a simple problem
	that has a LeftHandSide that is not well-defined (in this case,
	a variable). This should cause an error.
*/
func TestOptimizationProblem_ToSymbolicConstraint3(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := model1.AddVariable()
	v2 := optim.Variable{Lower: 0, Upper: -1}

	// Algorithm
	constr1 := optim.ScalarConstraint{
		LeftHandSide:  v2,
		RightHandSide: v1,
		Sense:         optim.SenseLessThanEqual,
	}
	_, err := problem.ToSymbolicConstraint(constr1)

	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			v2.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_ToSymbolicConstraint4
Description:

	Tests the ToSymbolicConstraint function with a simple constraint
	that has a RightHandSide that is not well-defined (in this case,
	a variable). This should cause an error.
*/
func TestOptimizationProblem_ToSymbolicConstraint4(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := optim.Variable{Lower: 0, Upper: -1}
	v2 := model1.AddVariable()

	// Algorithm
	constr1 := optim.ScalarConstraint{
		LeftHandSide:  v2,
		RightHandSide: v1,
		Sense:         optim.SenseLessThanEqual,
	}
	_, err := problem.ToSymbolicConstraint(constr1)

	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			v1.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From1
Description:

	Tests the From function with a simple
	model that doesn't have an objective.
*/
func TestOptimizationProblem_From1(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From1",
	)

	N := 5
	for ii := 0; ii < N; ii++ {
		model.AddVariable()
	}

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			"the input model has no objective function!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From2
Description:

	Tests the From function with a simple
	model that doesn't have an objective.
*/
func TestOptimizationProblem_From2(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From2",
	)

	N := 5
	var tempVar optim.Variable
	for ii := 0; ii < N; ii++ {
		tempVar = model.AddVariable()
	}

	model.SetObjective(tempVar, optim.SenseMaximize)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != 5 {
		t.Errorf("expected the number of variables to be %v; received %v",
			5, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}
}

/*
TestOptimizationProblem_From3
Description:

	Tests the From function with a convex optimization
	model that has a quadratic objective and
	at least two constraints.
*/
func TestOptimizationProblem_From3(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From3",
	)

	N := 5
	var tempVar optim.Variable
	for ii := 0; ii < N; ii++ {
		tempVar = model.AddVariable()
	}

	// Add a quadratic objective
	obj, err := tempVar.Multiply(tempVar)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = model.SetObjective(obj, optim.SenseMaximize)
	if err != nil {
		t.Errorf("error while setting objective! %v", err)
	}

	// Add a constraint
	constr1, err := tempVar.LessEq(1.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr1)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != N {
		t.Errorf("expected the number of variables to be %v; received %v",
			N, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}

	// Check that the number of constraints is as expected.
	if len(problem1.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(problem1.Constraints))
	}
}

/*
TestOptimizationProblem_From4
Description:

	Tests the From function with a convex optimization
	problem that has a linear objective and
	a vector inequality constraint.
*/
func TestOptimizationProblem_From4(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From4",
	)

	N := 5
	vv1 := model.AddVariableVector(N)
	vle2 := optim.VectorLinearExpr{
		L: *mat.NewDense(2, N, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		X: vv1,
		C: *mat.NewVecDense(2, []float64{21, 22}),
	}

	// Add a linear objective
	obj, err := vv1.Elements[0].Plus(vv1.Elements[2])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.SetObjective(obj, optim.SenseMaximize)

	// Add a vector constraint
	constr1, err := vle2.LessEq(*mat.NewVecDense(2, []float64{11, 12}))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr1)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != N {
		t.Errorf("expected the number of variables to be %v; received %v",
			N, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}

	// Check that the number of constraints is as expected.
	if len(problem1.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(problem1.Constraints))
	}
}

/*
TestOptimizationProblem_From5
Description:

	Tests the From function with a convex optimization
	problem that has a linear objective and
	two vector inequality constraints.
*/
func TestOptimizationProblem_From5(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From5",
	)

	N := 5
	vv1 := model.AddVariableVector(N)
	vle2 := optim.VectorLinearExpr{
		L: *mat.NewDense(2, N, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		X: vv1,
		C: *mat.NewVecDense(2, []float64{21, 22}),
	}

	vle3 := optim.VectorLinearExpr{
		L: *mat.NewDense(2, N, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		X: vv1,
		C: *mat.NewVecDense(2, []float64{31, 32}),
	}

	// Add a linear objective
	obj, err := vv1.Elements[0].Plus(vv1.Elements[2])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.SetObjective(obj, optim.SenseMaximize)

	// Add a vector constraint
	constr1, err := vle2.LessEq(*mat.NewVecDense(2, []float64{11, 12}))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr1)

	// Add another vector constraint
	constr2, err := vle3.LessEq(*mat.NewVecDense(2, []float64{41, 42}))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr2)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != N {
		t.Errorf("expected the number of variables to be %v; received %v",
			N, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}
}

/*
TestOptimizationProblem_From6
Description:

	Tests the From function properly produces an error
	when the input model is not well-defined.
*/
func TestOptimizationProblem_From6(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From6",
	)

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	} else {
		if !strings.Contains(
			err.Error(),
			"the model has no variables!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From7
Description:

	Tests the From function properly produces an error
	when the input model has an improperly defined objective.
*/
func TestOptimizationProblem_From7(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From7",
	)

	// Add a variable
	model.AddVariable()

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	} else {
		if !strings.Contains(
			err.Error(),
			"the input model has no objective function!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From8
Description:

	Tests the From function properly produces an error
	when the input model has an objective function that
	is not well-defined.
*/
func TestOptimizationProblem_From8(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From8",
	)

	// Add a variable
	v1 := model.AddVariable()

	// Add an objective
	model.SetObjective(
		optim.ScalarLinearExpr{
			L: *mat.NewVecDense(2, []float64{1, 2}),
			X: optim.VarVector{Elements: []optim.Variable{v1}},
			C: 1.2,
		},
		optim.SenseMaximize,
	)

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error, received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			"the length of L",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From9
Description:

	Tests that the From function properly produces an error
	when a constraint has been added to the problem that is not well-defined.
*/
func TestOptimizationProblem_From9(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From9",
	)

	// Add a variable
	v1 := model.AddVariable()

	// Add an objective
	model.SetObjective(v1, optim.SenseMaximize)

	// Add a constraint
	model.AddConstraint(optim.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: optim.Variable{Lower: 1, Upper: 0},
		Sense:         optim.SenseLessThanEqual,
	})

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error, received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf("there was a problem creating the %v-th constraint", 0),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From10
Description:

	Tests that the From function properly produces an error
	when the objective is not well-formed.
*/
func TestOptimizationProblem_From10(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From10",
	)

	// Add a variable
	v1 := model.AddVariable()

	// Add an objective
	model.SetObjective(optim.Variable{Lower: 0, Upper: -1}, optim.SenseMaximize)

	// Add a constraint
	model.AddConstraint(optim.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: optim.K(1.2),
		Sense:         optim.SenseLessThanEqual,
	})

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error, received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			model.Obj.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_Check1
Description:

	Tests the Check function with a simple problem
	that has one variable, one constraint and an objective
	that is not well-defined.
*/
func TestOptimizationProblem_Check1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_Check1")
	v1 := p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create bad objective
	p1.Objective = *problem.NewObjective(
		symbolic.Variable{}, problem.SenseMaximize,
	)

	// Algorithm
	err := p1.Check()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			p1.Objective.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_Check2
Description:

	Tests the Check function with a simple problem
	that has one variable, one well-defined objective
	and a set of constraints containing one bad constraint.
*/
func TestOptimizationProblem_Check2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_Check2")
	v1 := p1.AddVariable()
	c1 := symbolic.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: symbolic.Variable{},
		Sense:         symbolic.SenseLessThanEqual,
	}

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		v1, problem.SenseMaximize,
	)

	// Algorithm
	err := p1.Check()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			c1.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_Check3
Description:

	Tests the Check function with a simple problem
	that has one variable and no objective defined.
	The mpiErrors.NoObjectiveDefinedError should be created.
*/
func TestOptimizationProblem_Check3(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_Check3")
	v1 := p1.AddVariable()

	// Add variable to problem
	p1.Variables = append(p1.Variables, v1)

	// Algorithm
	err := p1.Check()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		expectedError := mpiErrors.NoObjectiveDefinedError{}
		if err.Error() != expectedError.Error() {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_Check4
Description:

	Tests the Check function with a simple problem
	that has:
	- objective defined
	- two variables (one is NOT well-defined)
	- and no constraints defined.
	The result should throw an error relating to the bad variable.
*/
func TestOptimizationProblem_Check4(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_Check4")
	v1 := p1.AddVariable()
	v2 := symbolic.Variable{}

	// Add variables to problem
	// p1.Variables = append(p1.Variables, v1) // Already added
	p1.Variables = append(p1.Variables, v2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		v1, problem.SenseMaximize,
	)

	// Algorithm
	err := p1.Check()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			v2.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_IsLinear1
Description:

	Tests the IsLinear function with a simple problem
	that has a constant objective and a single, linear constraint.
*/
func TestOptimizationProblem_IsLinear1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_IsLinear1")
	v1 := p1.AddVariable()
	c1 := v1.LessEq(1.0)
	k1 := symbolic.K(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		k1, problem.SenseFind,
	)

	// Algorithm
	isLinear := p1.IsLinear()
	if !isLinear {
		t.Errorf("expected the problem to be linear; received non-linear")
	}
}

/*
TestOptimizationProblem_IsLinear2
Description:

	Tests the IsLinear function with a simple problem
	that has a linear objective containing 3 variables and two lienar constraints,
	each containing one variable.
*/
func TestOptimizationProblem_IsLinear2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_IsLinear2")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).LessEq(1.0)
	c2 := vv1.AtVec(1).LessEq(1.0)

	// Add constraints
	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		vv1.Transpose().Multiply(symbolic.OnesVector(3)),
		problem.SenseMinimize,
	)

	// Algorithm
	isLinear := p1.IsLinear()
	if !isLinear {
		t.Errorf("expected the problem to be linear; received non-linear")
	}
}

/*
TestOptimizationProblem_IsLinear3
Description:

	Tests the IsLinear function with a simple problem
	that has a quadratic objective containing 3 variables and two lienar constraints,
	each containing one variable.
*/
func TestOptimizationProblem_IsLinear3(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_IsLinear3")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).LessEq(1.0)
	c2 := vv1.AtVec(1).LessEq(1.0)

	// Add constraints
	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		vv1.Transpose().Multiply(vv1),
		problem.SenseMaximize,
	)

	// Algorithm
	isLinear := p1.IsLinear()
	if isLinear {
		t.Errorf("expected the problem to be non-linear; received linear")
	}
}

/*
TestOptimizationProblem_IsLinear4
Description:

	Tests the IsLinear function with a simple problem
	that has a constant objective and a single, quadratic constraint.
	The problem should be non-linear.
*/
func TestOptimizationProblem_IsLinear4(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_IsLinear4")
	vv1 := p1.AddVariableVector(3)

	// Add constraints
	p1.Constraints = append(
		p1.Constraints,
		vv1.Transpose().Multiply(vv1).Plus(vv1).LessEq(1.0),
	)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	isLinear := p1.IsLinear()
	if isLinear {
		t.Errorf("expected the problem to be non-linear; received linear")
	}
}

/*
TestOptimizationProblem_IsLinear5
Description:

	Tests the IsLinear function with an optimization problem
	that is NOT well-defined.
	The function should cause a panic.
*/
func TestOptimizationProblem_IsLinear5(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_IsLinear5")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).LessEq(1.0)
	c2 := symbolic.ScalarConstraint{
		LeftHandSide: vv1.AtVec(1),
		RightHandSide: symbolic.Monomial{
			Coefficient:     1.0,
			VariableFactors: []symbolic.Variable{vv1.AtVec(1).(symbolic.Variable)},
			Exponents:       []int{1, 2},
		},
		Sense: symbolic.SenseLessThanEqual,
	}

	// Add constraints
	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		vv1.Transpose().Multiply(symbolic.OnesVector(3)),
		problem.SenseMinimize,
	)

	// Algorithm should panic with a not well-defined error
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("expected a panic; received none")
		}

		err, tf := r.(error)
		if !tf {
			t.Errorf("expected a panic of type error; received %T", r)
		}

		// Create the expected error
		expectedError := p1.Check()
		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf("expected a not well-defined panic; received %v", r)
		}

	}()

	p1.IsLinear()
	t.Errorf("expected a panic; received none")
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices1
Description:

	Tests the LinearInequalityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and a single linear inequality constraint.
	The result should be a matrix with 1 row and 2 columns.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices1")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearInequalityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 1 {
		t.Errorf("expected the number of rows to be %v; received %v",
			1, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 2 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 1 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			1, len(b))
	}
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices2
Description:

	Tests the LinearInequalityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and two scalar linear inequality constraints.
	The result should be a matrix with 2 rows and 2 columns.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices2")
	vv1 := p1.AddVariableVector(2)
	c1 := vv1.AtVec(0).LessEq(1.0)
	c2 := vv1.AtVec(1).LessEq(2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearInequalityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 2 {
		t.Errorf("expected the number of rows to be %v; received %v",
			2, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 2 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 2 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			2, len(b))
	}
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices3
Description:

	Tests the LinearInequalityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 3 variables,
	- and a single vector linear inequality constraint.
	The result should be a matrix with 3 rows and 3 columns.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices3(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices3")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.LessEq(symbolic.OnesVector(3))

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearInequalityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 3 {
		t.Errorf("expected the number of columns to be %v; received %v",
			3, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices4
Description:

	Tests the LinearInequalityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 3 variables,
	- and two vector linear inequality constraints.
	The result should be a matrix with 6 rows and 3 columns.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices4(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices4")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).Plus(symbolic.OnesVector(3)).LessEq(symbolic.OnesVector(3))
	c2 := vv1.AtVec(1).Plus(symbolic.OnesVector(3)).GreaterEq(symbolic.OnesVector(3))

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearInequalityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 6 {
		t.Errorf("expected the number of rows to be %v; received %v",
			6, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 3 {
		t.Errorf("expected the number of columns to be %v; received %v",
			3, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 6 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			6, len(b))
	}
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices5
Description:

	Tests the LinearInequalityConstraintMatrices function with a simple problem
	that looks like the one in TestOptimizationProblem_LinearInequalityConstraintMatrices1.
	The problem will have:
	- a constant objective
	- 2 variables,
	- a single linear inequality constraint,
	- and a single linear equality constraint.
	The result should be a matrix with 1 row and 2 columns.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices5(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices5")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.LessEq(1.0)
	c2 := v1.Eq(2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearInequalityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 1 {
		t.Errorf("expected the number of rows to be %v; received %v",
			1, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 2 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 1 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			1, len(b))
	}
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices6
Description:

	Tests the LinearInequalityConstraintMatrices function with a simple problem
	that contains a mixture of scalar and vector inequality constraints.
	The problem will have:
	- a constant objective
	- 3 variables,
	- a single vector linear inequality constraint,
	- and a single scalar linear inequality constraint.
	The result should be a matrix with 4 rows and 3 columns.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices6(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices6")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.LessEq(symbolic.OnesVector(3))
	c2 := vv1.AtVec(0).LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearInequalityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 4 {
		t.Errorf("expected the number of rows to be %v; received %v",
			4, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 3 {
		t.Errorf("expected the number of columns to be %v; received %v",
			3, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 4 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			4, len(b))
	}
}

/*
TestOptimizationProblem_LinearInequalityConstraintMatrices7
Description:

	Tests that the LinearInequalityConstraintMatrices function
	properly produces an error when the problem has NO inequality constraints.
*/
func TestOptimizationProblem_LinearInequalityConstraintMatrices7(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearInequalityConstraintMatrices7")
	v1 := p1.AddVariable()
	v2 := p1.AddVariable()

	// Create good objective
	p1.Objective = *problem.NewObjective(
		v1.Plus(v2.Multiply(0.5)),
		problem.SenseMaximize,
	)

	// Create a linear equality constraint
	p1.Constraints = append(p1.Constraints, v1.Plus(v2).Eq(1.0))

	// Algorithm
	_, _, err := p1.LinearInequalityConstraintMatrices()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			mpiErrors.NoInequalityConstraintsFoundError{}.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices1
Description:

	Tests the LinearEqualityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and a single linear equality constraint.
	The result should be a matrix with 1 row and 2 columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices1")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.Eq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 1 {
		t.Errorf("expected the number of rows to be %v; received %v",
			1, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 2 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 1 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			1, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices2
Description:

	Tests the LinearEqualityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and two scalar linear equality constraints.
	The result should be a matrix with 2 rows and 2 columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices2")
	vv1 := p1.AddVariableVector(2)
	c1 := vv1.AtVec(0).Eq(1.0)
	c2 := vv1.AtVec(1).Eq(2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 2 {
		t.Errorf("expected the number of rows to be %v; received %v",
			2, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 2 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 2 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			2, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices3
Description:

	Tests the LinearEqualityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 3 variables,
	- and a single vector linear equality constraint.
	The result should be a matrix with 3 rows and 3 columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices3(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices3")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.Eq(symbolic.OnesVector(3))

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 3 {
		t.Errorf("expected the number of columns to be %v; received %v",
			3, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices4
Description:

	Tests the LinearEqualityConstraintMatrices function with a simple problem.
	The problem will have:
	- a constant objective
	- 3 variables,
	- and two vector linear equality constraints.
	The result should be a matrix with 6 rows and 3 columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices4(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices4")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).Plus(symbolic.OnesVector(3)).Eq(symbolic.OnesVector(3))
	c2 := vv1.AtVec(1).Plus(symbolic.OnesVector(3)).Eq(symbolic.OnesVector(3))

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 6 {
		t.Errorf("expected the number of rows to be %v; received %v",
			6, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 3 {
		t.Errorf("expected the number of columns to be %v; received %v",
			3, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 6 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			6, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices5
Description:

	Tests the LinearEqualityConstraintMatrices function with a simple problem
	that looks like the one in TestOptimizationProblem_LinearEqualityConstraintMatrices1.
	The problem will have:
	- a constant objective
	- 2 variables,
	- a single linear equality constraint,
	- and a single linear inequality constraint.
	The result should be a matrix with 1 row and 2 columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices5(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices5")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.Eq(1.0)
	c2 := v1.LessEq(2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 1 {
		t.Errorf("expected the number of rows to be %v; received %v",
			1, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 2 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 1 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			1, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices6
Description:

	Tests the LinearEqualityConstraintMatrices function with a simple problem
	that contains a mixture of scalar and vector equality constraints.
	The problem will have:
	- a constant objective
	- 3 variables,
	- a single vector linear equality constraint,
	- and a single scalar linear equality constraint.
	The result should be a matrix with 4 rows and 3 columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices6(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices6")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.Eq(symbolic.OnesVector(3))
	c2 := vv1.AtVec(0).Eq(1.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	A, b, err := p1.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 4 {
		t.Errorf("expected the number of rows to be %v; received %v",
			4, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 3 {
		t.Errorf("expected the number of columns to be %v; received %v",
			3, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 4 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			4, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices7
Description:

	Tests the LinearEqualityConstraintMatrices function with a problem
	that led to panics in the field.
	The problem is Problem3 from our examples file.
	The problem will have:
	- a linear objective
	- 3 variables,
	- and a single linear VECTOR equality constraint (n_ineq = 1).
	Because each variable can be positive or negative, the resulting
	linear equality constraint matrix should have 3 rows and 3*2+n_ineq columns.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices7(t *testing.T) {
	// Constants
	p1 := problem.GetExampleProblem3()

	// Transform p1 into the standard form
	p1Standard, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Attempt to Call LinearEqualityConstraintMatrices
	A, b, err := p1Standard.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	nVariables1 := len(p1.Variables)
	nInequalityConstraints1 := p1.Constraints[0].Left().Dims()[0]
	if A.Dims()[1] != 2*nVariables1+nInequalityConstraints1 {
		t.Errorf("expected the number of columns to be %v; received %v",
			2*nVariables1+nInequalityConstraints1, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices8
Description:

	Tests the LinearEqualityConstraintMatrices function properly produces
	an NoEqualityConstraintsFoundError when the problem has NO equality constraints.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices8(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_LinearEqualityConstraintMatrices8")
	v1 := p1.AddVariable()
	v2 := p1.AddVariable()

	// Create good objective
	p1.Objective = *problem.NewObjective(
		v1.Plus(v2.Multiply(0.5)),
		problem.SenseMaximize,
	)

	// Create a linear inequality constraint
	p1.Constraints = append(p1.Constraints, v1.Plus(v2).LessEq(1.0))

	// Algorithm
	_, _, err := p1.LinearEqualityConstraintMatrices()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			mpiErrors.NoEqualityConstraintsFoundError{}.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices9
Description:

	Tests the LinearEqualityConstraintMatrices function with a problem
	that led to panics in the field.
	The problem is Problem4 from our examples file.
	The problem will have:
	- a linear objective
	- 3 variables (each labeled as positive via the LB input to AddVariableVectorClassic),
	- and a single linear VECTOR equality constraint.
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices9(t *testing.T) {
	// Constants
	p1 := problem.GetExampleProblem4()

	// Transform p1 into the standard form
	p1Standard, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Attempt to Call LinearEqualityConstraintMatrices
	A, b, err := p1Standard.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	nVariables1 := len(p1.Variables)
	nInequalityConstraints1 := p1.Constraints[0].Left().Dims()[0]
	if A.Dims()[1] != nVariables1+nInequalityConstraints1 {
		t.Errorf("expected the number of columns to be %v; received %v",
			nVariables1+nInequalityConstraints1, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_LinearEqualityConstraintMatrices10
Description:

	Tests the LinearEqualityConstraintMatrices function with a problem
	that led to panics in the field.
	The problem is Problem4 from our examples file.
	The problem will have:
	- a linear objective
	- 3 variables (each labeled as positive implicitly via the final 3 constraints),
	- and a single linear VECTOR equality constraint.
	The resulting problem should have 6 constraints (3 coming from inequality constraint)
	and 3 from the positivity constraints.
	The resulting linear equality constraint matrix should have 3 rows and 3+3 columns
	(3 variables, 3 inequality constraints, and 3 positivity constraints).
*/
func TestOptimizationProblem_LinearEqualityConstraintMatrices10(t *testing.T) {
	// Constants
	p1 := problem.GetExampleProblem5()

	// Transform p1 into the standard form
	p1Standard, slackVariables, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Attempt to Call LinearEqualityConstraintMatrices
	A, b, err := p1Standard.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.

	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	nVariables1 := len(p1.Variables)
	nSlackVariables1 := len(slackVariables)
	expectedNumVariables := nVariables1 + nSlackVariables1
	if A.Dims()[1] != expectedNumVariables {
		t.Errorf("expected the number of columns to be %v; received %v",
			expectedNumVariables, A.Dims()[1])
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_ToProblemWithAllPositiveVariables1
Description:

	Tests the ToProblemWithAllPositiveVariables function with a simple problem
	that has:
	- a constant objective
	- 2 variables,
	- and a single linear inequality constraint.
	The result should be a problem with 4 variables and 1 constraint.
*/
func TestOptimizationProblem_ToProblemWithAllPositiveVariables1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToProblemWithAllPositiveVariables1")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, err := p1.ToProblemWithAllPositiveVariables()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(p2.Variables) != 4 {
		t.Errorf("expected the number of variables to be %v; received %v",
			4, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(p2.Constraints))
	}

	// Verify that the new constraint contains two variables in the left hand side
	if len(p2.Constraints[0].Left().Variables()) != 2 {
		t.Errorf("expected the number of variables in the left hand side to be %v; received %v",
			2, len(p2.Constraints[0].Left().Variables()))
	}
}

/*
TestOptimizationProblem_ToProblemWithAllPositiveVariables2
Description:

	Tests the ToProblemWithAllPositiveVariables function with a simple problem
	that has:
	- a constant objective
	- 2 variables,
	- and two scalar linear inequality constraints.
	One of the variables is purely positive, while the other is purely negative.
	The result should be a problem with 2 variables and 2 constraints.
*/
func TestOptimizationProblem_ToProblemWithAllPositiveVariables2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToProblemWithAllPositiveVariables2")
	vv1 := p1.AddVariableVector(2)
	// Add constraints
	c1 := vv1.AtVec(0).GreaterEq(1.0)
	c2 := vv1.AtVec(1).LessEq(-2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, err := p1.ToProblemWithAllPositiveVariables()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(p2.Variables) != 2 {
		t.Errorf("expected the number of variables to be %v; received %v",
			2, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 2 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			2, len(p2.Constraints))
	}

	// Verify that the new constraints contain two variables in the left hand side
	for _, c := range p2.Constraints {
		if len(c.Left().Variables()) != 1 {
			t.Errorf("expected the number of variables in the left hand side to be %v; received %v",
				1, len(c.Left().Variables()))
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_1
Description:

	Tests the ToLPStandardForm function with a simple problem
	that contains:
	- a constant objective
	- 1 variable,
	- and a single linear inequality constraint (SenseGreaterThanEqual).
	The result should be a problem with 2 variables and 1 constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_1")
	v1 := p1.AddVariable()
	c1 := v1.GreaterEq(-1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	expectedNumVariables := 0
	expectedNumVariables += 2 * len(p1.Variables) // original variables (positive and negative halfs)
	expectedNumVariables += len(p1.Constraints)   // slack variables
	if len(p2.Variables) != expectedNumVariables {
		t.Errorf("expected the number of variables to be %v; received %v",
			expectedNumVariables, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(p2.Constraints))
	}

	// Verify that all constraints are equality constraints
	for _, c := range p2.Constraints {
		if c.ConstrSense() != symbolic.SenseEqual {
			t.Errorf("expected the constraint to be an equality constraint; received %v",
				c.ConstrSense())
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_2
Description:

	Tests the ToLPStandardForm function with a simple problem
	that contains:
	- a constant objective
	- 3 variables,
	- and a single vector linear inequality constraint (SenseGreaterThanEqual) of 5 dimensions.
	The result should be a problem with 3*2+5 = 11 variables and 1 constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_2")
	vv1 := p1.AddVariableVector(3)
	A2 := getKMatrix.From([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
		{10.0, 11.0, 12.0},
		{13.0, 14.0, 15.0},
	})
	c1 := A2.Multiply(vv1).GreaterEq(symbolic.OnesVector(5))

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	expectedNumVariables := 0
	expectedNumVariables += 2 * len(p1.Variables) // original variables (positive and negative halfs)
	p1FirstConstraint := p1.Constraints[0]
	p1FirstConstraintAsVC, ok := p1FirstConstraint.(symbolic.VectorConstraint)
	if !ok {
		t.Errorf("expected the first constraint to be a vector constraint; received %T", p1FirstConstraint)
	}
	expectedNumVariables += p1FirstConstraintAsVC.Dims()[0] // slack variables
	if len(p2.Variables) != expectedNumVariables {
		t.Errorf("expected the number of variables to be %v; received %v",
			expectedNumVariables, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 5 {
		t.Errorf(
			"expected the number of constraints to be %v; received %v",
			5,
			len(p2.Constraints),
		)
	}

	// Verify that all constraints are equality constraints
	for _, c := range p2.Constraints {
		if c.ConstrSense() != symbolic.SenseEqual {
			t.Errorf("expected the constraint to be an equality constraint; received %v",
				c.ConstrSense())
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_3
Description:

	Tests the ToLPStandardForm function with a simple problem
	that contains:
	- a constant objective
	- 3 variables,
	- and a single vector linear inequality constraint (SenseLessThanEqual) of 5 dimensions.
	The result should be a problem with 3*2+5 = 11 variables and 1 constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_3(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_3")
	vv1 := p1.AddVariableVector(3)
	A2 := getKMatrix.From([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
		{10.0, 11.0, 12.0},
		{13.0, 14.0, 15.0},
	})
	c1 := A2.Multiply(vv1).LessEq(symbolic.OnesVector(5))

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	expectedNumVariables := 0
	expectedNumVariables += 2 * len(p1.Variables) // original variables (positive and negative halfs)
	p1FirstConstraint := p1.Constraints[0]
	p1FirstConstraintAsVC, ok := p1FirstConstraint.(symbolic.VectorConstraint)
	if !ok {
		t.Errorf("expected the first constraint to be a vector constraint; received %T", p1FirstConstraint)
	}
	expectedNumVariables += p1FirstConstraintAsVC.Dims()[0] // slack variables
	if len(p2.Variables) != expectedNumVariables {
		t.Errorf("expected the number of variables to be %v; received %v",
			expectedNumVariables, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != A2.Dims()[0] {
		t.Errorf(
			"expected the number of constraints to be %v; received %v",
			A2.Dims()[0],
			len(p2.Constraints),
		)
	}

	// Verify that all constraints are equality constraints
	for _, c := range p2.Constraints {
		if c.ConstrSense() != symbolic.SenseEqual {
			t.Errorf(
				"expected the constraint to be an equality constraint; received %v",
				c.ConstrSense())
		}
	}

}

/*
TestOptimizationProblem_ToLPStandardForm1_4
Description:

	This test verifies that the ToLPStandardForm function throws an error
	when called on a problem that is not linear.
	In this case, we will define a problem with a quadratic objective function.
	The problem will have:
	- a quadratic objective
	- 2 variables,
	- and a single linear inequality constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_4(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_4")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create a quadratic objective
	p1.Objective = *problem.NewObjective(
		v1.Multiply(v1),
		problem.SenseMaximize,
	)

	// Algorithm
	_, _, _, err := p1.ToLPStandardForm1()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		expectedError := mpiErrors.ProblemNotLinearError{
			ProblemName:     p1.Name,
			Cause:           causeOfProblemNonlinearity.Objective,
			ConstraintIndex: -2,
		}
		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_5
Description:

	This test verifies that the ToLPStandardForm function throws an error
	when called on a problem that is not linear.
	In this case, we will define a problem with a quadratic constraint.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and a single quadratic inequality constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_5(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_5")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.Multiply(v1).LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	_, _, _, err := p1.ToLPStandardForm1()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		expectedError := mpiErrors.ProblemNotLinearError{
			ProblemName:     p1.Name,
			Cause:           causeOfProblemNonlinearity.Constraint,
			ConstraintIndex: 0,
		}
		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_6
Description:

	This test verifies that the ToLPStandardForm function properly handles
	a simple problem with a single, scalar linear inequality constraint.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and a single scalar linear inequality constraint (SenseLessThanEqual).
	The result should be a problem with 2*2+1 = 5 variables and 1 constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_6(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_6")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	expectedNumVariables := 0
	expectedNumVariables += 2 * len(p1.Variables) // original variables (positive and negative halfs)
	expectedNumVariables += len(p1.Constraints)   // slack variables
	if len(p2.Variables) != expectedNumVariables {
		t.Errorf("expected the number of variables to be %v; received %v",
			expectedNumVariables, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			expectedNumVariables, len(p2.Constraints))
	}

	// Verify that all constraints are equality constraints
	for _, c := range p2.Constraints {
		if c.ConstrSense() != symbolic.SenseEqual {
			t.Errorf("expected the constraint to be an equality constraint; received %v",
				c.ConstrSense())
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_7
Description:

	This test verifies that the ToLPStandardForm function properly handles
	a simple problem with a single, scalar equality constraint.
	The problem will have:
	- a constant objective
	- 2 variables,
	- and a single scalar linear equality constraint (SenseEqual).
	The result should be a problem with 2*2 = 4 variables and 1 constraint.
*/
func TestOptimizationProblem_ToLPStandardForm1_7(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_7")
	v1 := p1.AddVariable()
	p1.AddVariable()
	c1 := v1.Eq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p2, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	expectedNumVariables := 0
	expectedNumVariables += 3 // original variables (one with positive and negative halfs, the other with only the positive part)
	if len(p2.Variables) != expectedNumVariables {
		t.Errorf("expected the number of variables to be %v; received %v",
			expectedNumVariables, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			expectedNumVariables, len(p2.Constraints))
	}

	// Verify that all constraints are equality constraints
	for _, c := range p2.Constraints {
		if c.ConstrSense() != symbolic.SenseEqual {
			t.Errorf("expected the constraint to be an equality constraint; received %v",
				c.ConstrSense())
		}
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_8
Description:

	Tests the LinearEqualityConstraintMatrices function properly produces
	a matrix:
		[ 1 -1 0 0  0 0  1 0 0 ]
	C = [ 0 0  1 -1 0 0  0 1 0 ]
		[ 0 0  0 0  1 -1 0 0 1 ]
	and
		b = [ 1 2 3 ]
	By creating a problem with 3 variables and 3 linear inequality constraints.
	The results should produce equality constraint matrix C
	with 3 rows and 6 columns and a vector b with 3 elements.
*/
func TestOptimizationProblem_ToLPStandardForm1_8(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_8")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).LessEq(1.0)
	c2 := vv1.AtVec(1).LessEq(2.0)
	c3 := vv1.AtVec(2).LessEq(3.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)
	p1.Constraints = append(p1.Constraints, c3)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p1Prime, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	A, b, err := p1Prime.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 9 {
		t.Errorf("expected the number of columns to be %v; received %v",
			6, A.Dims()[1])
	}

	// Check that each of the entries in A is as expected.
	expectedA := getKMatrix.From([][]float64{
		{1.0, -1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, -1.0, 0.0, 0.0, 0.0, 1.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 1.0, -1.0, 0.0, 0.0, 1.0},
	})
	AAsDense := A.ToDense()
	expectedAAsDense := expectedA.ToDense()
	if !mat.EqualApprox(
		&AAsDense,
		&expectedAAsDense,
		1e-10,
	) {
		t.Errorf("expected A to be %v; received %v", expectedA, A)
		t.Errorf("Variables in A: %v", p1Prime.Variables)
		t.Errorf("Linear coefficient for constraint 1: %v", p1Prime.Constraints[0].Left().(symbolic.ScalarExpression).LinearCoeff(p1Prime.Variables))
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_9
Description:

	Tests the LinearEqualityConstraintMatrices function properly produces
	a matrix:
		[ 1 -1 0 0  0 0  -1 0  0 ]
	C = [ 0 0  1 -1 0 0  0  -1 0 ]
		[ 0 0  0 0  1 -1 0  0  -1 ]
	and
		b = [ -1 -2 -3 ]
	By creating a problem with 3 variables and 3 linear inequality constraints.
	The results should produce equality constraint matrix C
	with 3 rows and 6 columns and a vector b with 3 elements.
*/
func TestOptimizationProblem_ToLPStandardForm1_9(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_9")
	vv1 := p1.AddVariableVector(3)
	c1 := vv1.AtVec(0).GreaterEq(-1.0)
	c2 := vv1.AtVec(1).GreaterEq(-2.0)
	c3 := vv1.AtVec(2).GreaterEq(-3.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)
	p1.Constraints = append(p1.Constraints, c3)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	p1Prime, _, _, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	A, b, err := p1Prime.LinearEqualityConstraintMatrices()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of rows is as expected.
	if A.Dims()[0] != 3 {
		t.Errorf("expected the number of rows to be %v; received %v",
			3, A.Dims()[0])
	}

	// Check that the number of columns is as expected.
	if A.Dims()[1] != 9 {
		t.Errorf("expected the number of columns to be %v; received %v",
			6, A.Dims()[1])
	}

	// Check that each of the entries in A is as expected.
	expectedA := getKMatrix.From([][]float64{
		{1.0, -1.0, 0.0, 0.0, 0.0, 0.0, -1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, -1.0, 0.0, 0.0, 0.0, -1.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 1.0, -1.0, 0.0, 0.0, -1.0},
	})
	AAsDense := A.ToDense()
	expectedAAsDense := expectedA.ToDense()
	if !mat.EqualApprox(
		&AAsDense,
		&expectedAAsDense,
		1e-10,
	) {
		t.Errorf("expected A to be %v; received %v", expectedA, A)
		t.Errorf("Variables in A: %v", p1Prime.Variables)
		t.Errorf("Linear coefficient for constraint 1: %v", p1Prime.Constraints[0].Left().(symbolic.ScalarExpression).LinearCoeff(p1Prime.Variables))
	}

	// Check that the number of elements in b is as expected.
	if len(b) != 3 {
		t.Errorf("expected the number of elements in b to be %v; received %v",
			3, len(b))
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_10
Description:

	This method verifies that the method will return an error
	if the optimization problem is not well-defined.
	In this case, we will create a problem with a constraint
	that has mismatched dimensions.
*/
func TestOptimizationProblem_ToLPStandardForm1_10(t *testing.T) {
	// Setup
	N := 10

	// Create objective function
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_10")
	x := p1.AddVariableVector(N)
	p1.Constraints = append(
		p1.Constraints,
		symbolic.VectorConstraint{
			LeftHandSide:  x,
			RightHandSide: symbolic.VecDenseToKVector(symbolic.OnesVector(N + 1)), // Mismatched dimensions
			Sense:         symbolic.SenseLessThanEqual,
		},
	)

	p1.Objective = *problem.NewObjective(
		x.Transpose().Multiply(symbolic.OnesVector(x.Len())),
		problem.SenseMinimize,
	)

	// Call the ToLPStandardForm1
	_, _, _, err := p1.ToLPStandardForm1()
	if err == nil {
		t.Errorf("expected an error; received nil")
	}

	// Create the expected error
	expectedError := p1.MakeNotWellDefinedError()
	if err.Error() != expectedError.Error() {
		t.Errorf("unexpected error: %v", err)
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_11
Description:

	This method verifies that the map from original variables to
	standard form variables is correct for a small problem.
	In this problem, we will have:
	- a constant objective
	- 1 variable,
	- and a single linear inequality constraint.
	The resulting map should contain 1 entry, mapping the original
	variable to the positive half and negative half variables in the standard form.
*/
func TestOptimizationProblem_ToLPStandardForm1_11(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_11")
	v1 := p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	_, _, varMap, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of entries in the map is as expected.
	if len(varMap) != 1 {
		t.Errorf("expected the number of entries in the map to be %v; received %v",
			1, len(varMap))
	}

	// Check that the entry in the map contains two variables.
	v1Expr, ok := varMap[v1]
	if !ok {
		t.Errorf("expected the map to contain an entry for variable %v; it does not", v1)
	}

	if len(v1Expr.Variables()) != 2 {
		t.Errorf("expected the entry in the map to contain %v variables; received %v",
			2, len(v1Expr.Variables()))
	}

	if len(symbolic.UniqueVars(v1Expr.Variables())) != 2 {
		t.Errorf("expected the entry in the map to contain %v unique variables; received %v",
			2, len(symbolic.UniqueVars(v1Expr.Variables())))
	}
}

/*
TestOptimizationProblem_ToLPStandardForm1_12
Description:

	This method verifies that the map from original variables to
	standard form variables is correct for a small problem.
	In this problem, we will have:
	- a constant objective
	- 2 variables,
	- and two linear inequality constraints.
	One of the variables is purely positive, while the other is purely negative.
	The resulting map should contain 2 entries, one mapping the purely positive
	variable to itself, and the other mapping the purely negative variable to
	the negative half variables in the standard form.
*/
func TestOptimizationProblem_ToLPStandardForm1_12(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm1_12")
	vv1 := p1.AddVariableVector(2)
	// Add constraints
	c1 := vv1.AtVec(0).GreaterEq(1.0)
	c2 := vv1.AtVec(1).LessEq(-2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	_, _, varMap, err := p1.ToLPStandardForm1()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of entries in the map is as expected.
	if len(varMap) != 2 {
		t.Errorf("expected the number of entries in the map to be %v; received %v",
			2, len(varMap))
	}

	// Check that the entry in the map for the purely positive variable contains one variable.
	v1 := vv1.AtVec(0).(symbolic.Variable)
	v1Expr, ok := varMap[v1]
	if !ok {
		t.Errorf("expected the map to contain an entry for variable %v; it does not", v1)
	}

	if len(v1Expr.Variables()) != 1 {
		t.Errorf("expected the entry in the map to contain %v variables; received %v",
			1, len(v1Expr.Variables()))
	}

	// Check that the entry in the map for the purely negative variable contains one variable.
	v2 := vv1.AtVec(1).(symbolic.Variable)
	v2Expr, ok := varMap[v2]
	if !ok {
		t.Errorf("expected the map to contain an entry for variable %v; it does not", v2)
	}

	if len(v2Expr.Variables()) != 1 {
		t.Errorf("expected the entry in the map to contain %v variables; received %v",
			1, len(v2Expr.Variables()))
	}

	// TODO: Uncomment these checks once we have Polynomial's AsSimplifiedExpression method better tuned for polynomials.
	// v2ExprAsMonomial, ok := v2Expr.(symbolic.Monomial)
	// if !ok {
	// 	t.Errorf("expected the entry in the map to be a Monomial; received %T", v2Expr)
	// }

	// coeff2 := v2ExprAsMonomial.LinearCoeff()
	// if coeff2.AtVec(0) != -1.0 {
	// 	t.Errorf("expected the entry in the map to have coefficient %v; received %v",
	// 		-1.0, coeff2.AtVec(0))
	// }
}

/*
TestOptimizationProblem_CheckIfLinear1
Description:

	This test verifies that the CheckIfLinear function properly identifies
	a NOT well-defined problem is not linear.
	The problem will have a vector constraint with mismatched dimensions.
*/
func TestOptimizationProblem_CheckIfLinear1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_CheckIfLinear1")
	vv1 := p1.AddVariableVector(3)
	A2 := getKMatrix.From([][]float64{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
	})
	c1 := symbolic.VectorConstraint{
		LeftHandSide:  A2.Multiply(vv1).(symbolic.VectorExpression),
		RightHandSide: getKVector.From(symbolic.OnesVector(5)),
		Sense:         symbolic.SenseLessThanEqual,
	}

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Algorithm
	err := p1.CheckIfLinear()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		expectedError := mpiErrors.ProblemNotLinearError{
			ProblemName:     p1.Name,
			Cause:           causeOfProblemNonlinearity.NotWellDefined,
			ConstraintIndex: -1,
		}
		if !strings.Contains(
			err.Error(),
			expectedError.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_SimplifyConstraints1
Description:

	This test verifies that the SimplifyConstraints function properly simplifies
	a problem with redundant constraints.
	The problem will have:
	- a constant objective
	- 1 variable1,
	- and 2 linear inequality constraints:
		x1 <= 1
		x1 <= 2
	The second constraint is redundant and should be removed.
	The result should be a problem with 1 variable and 1 constraint.
*/
func TestOptimizationProblem_SimplifyConstraints1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SimplifyConstraints1")
	v1 := p1.AddVariable()
	c1 := v1.LessEq(1.0)
	c2 := v1.LessEq(2.0)

	p1.Constraints = append(p1.Constraints, c1)
	p1.Constraints = append(p1.Constraints, c2)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Create a copy of p1 to compare against after simplification
	p2 := p1.Copy()

	// Algorithm
	p1.SimplifyConstraints()

	// Check that the number of variables is as expected.
	if len(p2.Variables) != len(p1.Variables) {
		t.Errorf("expected the number of variables to be %v; received %v",
			len(p1.Variables), len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p1.Constraints) != 1 {
		t.Errorf("expected the new number of constraints to be %v; received %v",
			1, len(p1.Constraints))
	}

	// Check that the remaining constraint is the expected one.
	p1FirstConstraint := p1.Constraints[0]
	L1 := p1FirstConstraint.Left().(symbolic.ScalarExpression).LinearCoeff(p1.Variables)
	p2FirstConstraint := p2.Constraints[0]
	L2 := p2FirstConstraint.Left().(symbolic.ScalarExpression).LinearCoeff(p2.Variables)

	if L1.AtVec(0) != L2.AtVec(0) {
		t.Errorf("expected the remaining constraint's coefficient to be %v; received %v",
			L1, L2)
	}
}

/*
TestOptimizationProblem_SimplifyConstraints2
Description:

	This test verifies that the SimplifyConstraints function properly handles
	a problem with no constraints.
	The problem will have:
	- a constant objective
	- 1 variable,
	- and no constraints.
	The result should be a problem with 1 variable and no constraints.
*/
func TestOptimizationProblem_SimplifyConstraints2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SimplifyConstraints2")

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Create a copy of p1 to compare against after simplification
	p2 := p1.Copy()

	// Algorithm
	p1.SimplifyConstraints()

	// Check that the number of variables is as expected.
	if len(p2.Variables) != len(p1.Variables) {
		t.Errorf("expected the number of variables to be %v; received %v",
			len(p1.Variables), len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p1.Constraints) != 0 {
		t.Errorf("expected the new number of constraints to be %v; received %v",
			0, len(p1.Constraints))
	}
}

/*
TestOptimizationProblem_SimplifyConstraints3
Description:

	This test verifies that the SimplifyConstraints function properly handles
	a problem with a single constraint that is not redundant.
	The problem will have:
	- a constant objective
	- 1 variable,
	- and a single linear inequality constraint.
	The result should be a problem with 1 variable and 1 constraint.
*/
func TestOptimizationProblem_SimplifyConstraints3(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SimplifyConstraints3")
	v1 := p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		symbolic.K(3.14),
		problem.SenseMaximize,
	)

	// Create a copy of p1 to compare against after simplification
	p2 := p1.Copy()

	// Algorithm
	p1.SimplifyConstraints()

	// Check that the number of variables is as expected.
	if len(p2.Variables) != len(p1.Variables) {
		t.Errorf("expected the number of variables to be %v; received %v",
			len(p1.Variables), len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p1.Constraints) != 1 {
		t.Errorf("expected the new number of constraints to be %v; received %v",
			1, len(p1.Constraints))
	}

	// Check that the remaining constraint is the expected one.
	p1FirstConstraint := p1.Constraints[0]
	L1 := p1FirstConstraint.Left().(symbolic.ScalarExpression).LinearCoeff(p1.Variables)
	p2FirstConstraint := p2.Constraints[0]
	L2 := p2FirstConstraint.Left().(symbolic.ScalarExpression).LinearCoeff(p2.Variables)

	if L1.AtVec(0) != L2.AtVec(0) {
		t.Errorf("expected the remaining constraint's coefficient to be %v; received %v",
			L1, L2)
	}
}

/*
TestOptimizationProblem_ToLPStandardForm2_1
Description:

	Tests the ToLPStandardForm2 function with a simple problem
	that contains:
	- a linear objective,
	- a MINIMIZATION sense
	- 1 variable,
	- and a single linear inequality constraint (SenseGreaterThanEqual).
	The result should be a problem with 2 variables and 1 constraint.
	The sense of the resulting problem should be MAXIMIZATION.
*/
func TestOptimizationProblem_ToLPStandardForm2_1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_ToLPStandardForm2_1")
	v1 := p1.AddVariable()
	c1 := v1.GreaterEq(-1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		v1,
		problem.SenseMinimize,
	)

	// Algorithm
	p2, _, _, err := p1.ToLPStandardForm2()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	expectedNumVariables := 0
	expectedNumVariables += 2 * len(p1.Variables) // original variables (positive and negative halfs)
	expectedNumVariables += len(p1.Constraints)   // slack variables
	if len(p2.Variables) != expectedNumVariables {
		t.Errorf("expected the number of variables to be %v; received %v",
			expectedNumVariables, len(p2.Variables))
	}

	// Check that the number of constraints is as expected.
	if len(p2.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(p2.Constraints))
	}

	// Verify that all constraints are equality constraints
	for _, c := range p2.Constraints {
		if c.ConstrSense() != symbolic.SenseEqual {
			t.Errorf("expected the constraint to be an equality constraint; received %v",
				c.ConstrSense())
		}
	}

	// Verify that the sense of the objective is MAXIMIZATION
	if p2.Objective.Sense != problem.SenseMaximize {
		t.Errorf("expected the sense of the objective to be %v; received %v",
			problem.SenseMaximize, p2.Objective.Sense)
	}
}

/*
TestOptimizationProblem_ToLPStandardForm2_2
Description:

	Tests the ToLPStandardForm2 function with the challenging
	problem described in gonum issue:
		https://github.com/gonum/gonum/issues/1914
	The problem is fairly large with 57 variables.
	We will verify that the function correctly creates:
	- an objective that contains BOTH positive and negative coefficients,
	- a MAXIMIZATION sense
*/
func TestOptimizationProblem_ToLPStandardForm2_2(t *testing.T) {
	// Create the problem described in the gonum issue
	pGonum := problem.GetGonumBugLPProblem()

	// Convert to standard form
	p2, _, _, err := pGonum.ToLPStandardForm2()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify that the sense of the objective is MAXIMIZATION
	if p2.Objective.Sense != problem.SenseMaximize {
		t.Errorf("expected the sense of the objective to be %v; received %v",
			problem.SenseMaximize, p2.Objective.Sense)
	}

	// Verify that the objective contains both positive and negative coefficients
	hasPositive := false
	hasNegative := false
	objAsPolynomial, ok := p2.Objective.Expression.(symbolic.Polynomial)
	if !ok {
		t.Errorf("expected the objective to be a Polynomial; received %T",
			p2.Objective.Expression)
	}

	coeffs := objAsPolynomial.LinearCoeff(p2.Variables)

	for _, coeff := range coeffs.RawVector().Data {
		if coeff > 0 {
			hasPositive = true
		}
		if coeff < 0 {
			hasNegative = true
		}
	}
	if !hasPositive {
		t.Errorf("expected the objective to have at least one positive coefficient; it does not")
	}
	if !hasNegative {
		t.Errorf("expected the objective to have at least one negative coefficient; it does not")
	}
}

/*
TestOptimizationProblem_CopyVariable1
Description:

	This method tests that the CopyVariable method for OptimizationProblem
	properly creates a copy of one variable in an optimization problem.
	Check that:
	- the new variable has a different ID than the one that is copied
	- the new variable has a slightly different name than the one that is copied
*/
func TestOptimizationProblem_CopyVariable1(t *testing.T) {
	p1 := problem.NewProblem("TestOptimizationProblem_CopyVariable1")
	v1 := p1.AddVariable()

	// Create a copy of the variable
	v2 := p1.CopyVariable(v1)

	// Check that the new variable has a different ID
	if v1.ID == v2.ID {
		t.Errorf(
			"expected the new variable to have a different ID; received %v",
			v2.ID,
		)
	}

	// Check that the new variable has a slightly different name
	if v1.Name == v2.Name {
		t.Errorf(
			"expected the new variable to have a slightly different name; received %v",
			v2.Name,
		)
	}
}

/*
TestOptimizationProblem_String1
Description:

	Tests that a small optimization problem with all scalar constraints gets represented
	as a string with:
	- Minimize sense of objective
	- The objective expression is completely contained in the string
	- the string describes that there are 0 vector constraints and 0 matrix constraints
*/
func TestOptimizationProblem_String1(t *testing.T) {
	// Create Optimization Problem
	p := problem.NewProblem("TestOptimizationProblem_String1")

	N := 2
	x := p.AddVariableVector(N)
	c := symbolic.OnesVector(N)
	objExpr := x.Transpose().Multiply(c)
	p.SetObjective(objExpr, problem.SenseMinimize)

	p.Constraints = append(p.Constraints, x.AtVec(0).LessEq(1.2))
	p.Constraints = append(p.Constraints, x.AtVec(1).GreaterEq(3.14))

	// Create String
	pAsString := fmt.Sprintf("%s", p)

	// Check that the string has "Minimize" in it
	if !strings.Contains(pAsString, "Minimize") {
		t.Errorf(
			"Problem string does not contain the string \"Minimize\".",
		)
	}

	if !strings.Contains(pAsString, objExpr.String()) {
		t.Errorf(
			"Problem string does not contain the expression in the objective %s",
			objExpr,
		)
	}

	if !strings.Contains(pAsString, "0 vector constraints") {
		t.Errorf(
			"Problem string does not contain \"0 vector constraints\".",
		)
	}

	if !strings.Contains(pAsString, "0 matrix constraints") {
		t.Errorf(
			"Problem string does not contain \"0 matrix constraints\".",
		)
	}
}

/*
TestOptimizationProblem_String2
Description:

	Tests that a small optimization problem with all scalar constraints gets represented
	as a string with:
	- Maximize sense of objective
	- the string describes that there are 2 vector constraints and 1 matrix constraints
*/
func TestOptimizationProblem_String2(t *testing.T) {
	// Create Optimization Problem
	p := problem.NewProblem("TestOptimizationProblem_String1")

	N := 2
	x := p.AddVariableVector(N)
	y := p.AddVariableMatrix(N, N, 0.0, 200.0, symbolic.Continuous)
	c := symbolic.OnesVector(N)
	objExpr := x.Transpose().Multiply(c)
	p.SetObjective(objExpr, problem.SenseMaximize)

	p.Constraints = append(p.Constraints, x.LessEq(c))
	p.Constraints = append(p.Constraints, x.GreaterEq(symbolic.ZerosVector(N)))
	p.Constraints = append(p.Constraints, y.Eq(symbolic.ZerosMatrix(N, N)))

	// Create String
	pAsString := fmt.Sprintf("%s", p)

	// Check that the string has "Maximize" in it
	if !strings.Contains(pAsString, "Maximize") {
		t.Errorf(
			"Problem string does not contain the string \"Maximize\".",
		)
	}

	if !strings.Contains(pAsString, "2 vector constraints") {
		t.Errorf(
			"Problem string does not contain \"2 vector constraints\".",
		)
	}

	if !strings.Contains(pAsString, "1 matrix constraints") {
		t.Errorf(
			"Problem string does not contain \"1 matrix constraints\".",
		)
	}
}
