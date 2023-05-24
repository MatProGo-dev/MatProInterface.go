[![Go Reference](https://pkg.go.dev/badge/github.com/MatProGo-dev/MatProInterface.go.svg)](https://pkg.go.dev/github.com/MatProGo-dev/MatProInterface.go)
![Coverage](https://img.shields.io/badge/Coverage-67.1%25-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/MatProGo-dev/MatProInterface.go)](https://goreportcard.com/report/github.com/MatProGo-dev/MatProInterface.go)

# MatProInterface.go
A common interface used for modeling Mathematical Programs in the language Go.

| ![](images/scalar-range-optimization1.png) |
|:------------------------------------------:|
|  Effectively Model Mathematical Programs   |

## How to Install

```
go get github.com/The-Velo-Network/MatProInterface.go
```

The interface is very useful on its own, but typically you won't want to install it alone.
You should use it with a solver that can address the problems specified
by your model.

## Available Solvers

- [Gurobi](https://github.com/MatProGo-dev/Gurobi.go)

## Modeling the Mathematical Program Above
For example, to model the program above one would write the following code:
```
// Constants
modelName := "mpg-qp1"
m := optim.NewModel(modelName)
x := m.AddVariableVector(2)

gs := mpgSolver.NewGurobiSolver(
    fmt.Sprintf("solvertest-%v", modelName),
)

// Create Vector Constants
c1 := optim.KVector(
    *mat.NewVecDense(2, []float64{0.0, 1.0}),
)

c2 := optim.KVector(
    *mat.NewVecDense(2, []float64{2.0, 3.0}),
)

// Use these to create constraints.

vc1, err := x.LessEq(c2)
if err != nil {
    t.Errorf("There was an issue creating the proper vector constraint: %v", err)
}

vc2, err := x.GreaterEq(c1)
if err != nil {
    t.Errorf("There was an issue creating the proper vector constraint: %v", err)
}

// Create objective
Q1 := optim.Identity(x.Len())
Q1.Set(0, 1, 0.25)
Q1.Set(1, 0, 0.25)
Q1.Set(1, 1, 0.25)

obj := optim.ScalarQuadraticExpression{
    Q: Q1,
    X: x,
    L: *mat.NewVecDense(x.Len(), []float64{0, -0.97}),
    C: 2.0,
}

// Add Constraints
constraints := []optim.Constraint{vc1, vc2}
for _, constr := range constraints {
    err = gs.AddConstraint(constr)
    if err != nil {
        t.Errorf("There was an issue adding the vector constraint to the model: %v", err)
    }
}

// Add objective
err = m.SetObjective(optim.Objective{obj, optim.SenseMinimize})
if err != nil {
    t.Errorf("There was an issue setting the objective of the Gurobi solver model: %v", err)
}

// Solve!
sol, err := m.Optimize()
if err != nil {
    t.Errorf("There was an issue optimizing the QP: %v", err)
}
```

## FAQs

> Why are the solvers not bundled into the interface?

The solvers are separated into separate repositories to avoid compilation issues.
A compilation issue would arise, for example, if Gurobi bindings were built into this interace,
but your computer did not have Gurobi installed on it. The same can be said for a number of other
solvers as well. To avoid such issues, ALL SOLVERS should be included in separate pacakages
that implement the `solver` interface in this package.

With this in mind, you should be able to use any solver by installing its associated
MatProGo.dev pacakage and then calling its "Solver" object.

> I feel like things can be done more efficiently in this library.
> Why did you avoid using things like pointer receivers?

This project was written to make it easier for first-time Go contributors/users
to easily understand. For this reason, I've avoided making some optimizations
that might improve speed but might confuse a less experienced Go programmer.

Still, there might be behaviors that occur that confuse you.
For example, you might realize that when
you manipulate certain variables with this library, the
objects are passed by reference and not by value. Feel free to
ask if this is intentional by creating an issue.

## To-Dos

* [X] Create New AddConstr methods which work for vector constraints
* [ ] Mult
  * [ ] General Function (in operators.go)
    * [ ] Methods for
        * [ ] Scalars
            * [X] Constant
            * [ ] Var
            * [ ] ScalarLinearExpression
            * [ ] QuadraticExpression
        * [ ] Vectors
            * [ ] Vector Constant
            * [ ] VarVector
            * [ ] VectorLinearExpression
* [ ] Plus
    * [ ] General Function (in operators.go)
* [ ] Introducing Optional Input for Variable Name to Var/VarVector
* [ ] Consider renaming VarVector to VectorVar
* [ ] VarVector
    * [ ] Plus
    * [ ] Multiply
    * [ ] LessEq
    * [ ] GreaterEq
    * [ ] Eq
    * [ ] Len
* [ ] Decide whether or not we really need the Coeffs() method (What is it doing?)
* [ ] Create function for easily creating MatDense:
    * [ ] ones matrices
* [ ] Create function for:
    * [ ] IsScalar()
    * [ ] IsVector()
* [X] VectorConstraint
    * [X] AtVec()
* [ ] Write changes to all AtVec() methods to output both elements AND errors (so we can detect out of length calls)
* [ ] Investigate where logrus was used and plan around it.