[![Go Reference](https://pkg.go.dev/badge/github.com/MatProGo-dev/MatProInterface.go.svg)](https://pkg.go.dev/github.com/MatProGo-dev/MatProInterface.go)
[![codecov](https://codecov.io/gh/MatProGo-dev/MatProInterface.go/branch/main/graph/badge.svg?token=RCIN5S1AK7)](https://codecov.io/gh/MatProGo-dev/MatProInterface.go)
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
    err = m.AddConstraint(constr)
    if err != nil {
        t.Errorf("There was an issue adding the vector constraint to the model: %v", err)
    }
}

// Add objective
err = m.SetObjective(optim.Objective{obj, optim.SenseMinimize})
if err != nil {
    t.Errorf("There was an issue setting the objective of the Gurobi solver model: %v", err)
}

// Solve using the solver of your choice!
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

> Why do most functions return `error` values?

There are two dominant approaches for handling errors/problems
during numerical Go programs. One is to raise an exception/create a fatal flag which terminates the program.
The other is to share error messages to the user using Go's build-in `error` type (or extensions of it) during
most function calls. In most cases, these error messages will be `nil` indicating
that no error occurred, but occasionally they will contain valuable information.

The second approach is used here because it may be helpful for the library to explain to the user what is going on
and how to use certain functions through a direct message. 
Sometimes, the first method of error handling can point users to unintuitive/difficult to read parts of
code. Hopefully, this is avoided using this format.

## Design Philosophies

* Share error information
* Composability
  * It should be possible to compose any math operation with any other (Assuming there are no dimension mismatch errors).

## To-Dos

* [X] Create New AddConstr methods which work for vector constraints
* [ ] Mult
  * [ ] General Function (in operators.go)
* [ ] Plus
    * [ ] General Function (in operators.go)
* [ ] Introducing Optional Input for Variable Name to Var/VarVector
* [ ] Consider renaming VarVector to VectorVar
* [ ] Decide whether or not we really need the Coeffs() method (What is it doing?)
* [ ] Write changes to all AtVec() methods to output both elements AND errors (so we can detect out of length calls)
* [ ] Determine whether or not to keep the Solution and Solver() interfaces in this module. It seems like they can be solver-specific.
* [ ] Add Check() to:
  * [ ] Expression
  * [ ] ScalarExpression
  * [ ] VectorExpression interfaces
* [ ] Add ToSymbolic() Method for ALL expressions