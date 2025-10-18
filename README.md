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

- [simplex](https://github.com/MatProGo-dev/simplex)
- [Gurobi](https://github.com/MatProGo-dev/Gurobi.go)

## Modeling the Mathematical Program Above
For example, to model the program above one would write the following code:
```

import (
	...
    "github.com/MatProGo-dev/MatProInterface.go/problem"
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
    ...
)

// Constants
problemName := "mpg-qp1"
p1 := problem.NewProblem(problemName)
x := p1.AddVariableVector(2)

// Create Vector Constants
c1 := getKVector.From([]float64{0.0, 1.0})
c2 := getKVector.From([]float64{2.0, 3.0})

// Use these to create constraints.
vc1 := x.LessEq(c2)
vc2 := x.GreaterEq(c1)

p1.Constraints = append(p1.Constraints, vc1)
p1.Constraints = append(p1.Constraints, vc2)

// Create objective
Q1 := symbolic.Identity(x.Len())
Q1.Set(0, 1, 0.25)
Q1.Set(1, 0, 0.25)
Q1.Set(1, 1, 0.25)

p1.Objective = *problem.NewObjective(
    x.Transpose().Multiply(Q).Multiply(x),
    problem.SenseMinimize,
)

// Solve using the solver of your choice!
```

## FAQs

> Why are the solvers not bundled into `MatProInterface.go`?

The solvers are not bundled into this module to avoid having a large number of dependencies. In other words, to avoid including a large number of solvers in the dependencies (e.g., Gurobi, SNOPT), we will require developers to implement those in separate modules. This should hopefully keep this module "lightweight".

> Why did you avoid using things like pointer receivers in your implementations?

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
in numerical Go programs. One is to raise an exception/create a fatal flag which terminates the program.
The other is to share error messages to the user using Go's built-in `error` type (or extensions of it) during
most function calls. In most cases, these error messages will be `nil` indicating
that no error occurred, but occasionally they will contain valuable information.

The second approach is used here because it may be helpful for the library to explain to the user what is going on
and how to use certain functions through a direct message. 
Sometimes, the first method of error handling can point users to unintuitive/difficult to read parts of
code. Hopefully, this is avoided using this format.

> I have a question for you. How can I ask it?

Feel free to create an issue (see the Issues tab above) in this repository!

> I have noticed an error in the code. What should I do?

You can either:
- Create an issue describing the error and how to reproduce it (see the Issues tab above)
- Fork the repository and create a Pull Request

## Design Philosophies

* Share error information
* Composability
  * It should be possible to compose any math operation with any other (Assuming there are no dimension mismatch errors).

## To-Dos

* [X] Create New AddConstr methods which work for vector constraints
* [ ] Deprecate the `optim` package
* [ ] Introducing Optional Input for Variable Name to Var/VarVector
* [ ] Add ToSymbolic() Method for ALL expressions