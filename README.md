[![Go Reference](https://pkg.go.dev/badge/github.com/MatProGo-dev/MatProInterface.go.svg)](https://pkg.go.dev/github.com/MatProGo-dev/MatProInterface.go)
![Coverage](https://img.shields.io/badge/Coverage-65.1%25-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/MatProGo-dev/MatProInterface.go)](https://goreportcard.com/report/github.com/MatProGo-dev/MatProInterface.go)

# MatProInterface.go
A common interface used for modeling Mathematical Programs in the language Go.

## FAQs

> Why are the solvers not bundled into the interface?

The solvers are separated into separate repositories to avoid compilation issues.
A compilation issue would arise, for example, if Gurobi bindings were built into this interace,
but your computer did not have Gurobi installed on it. The same can be said for a number of other
solvers as well. To avoid such issues, ALL SOLVERS should be included in separate pacakages
that implement the `solver` interface in this package.

With this in mind, you should be able to use any solver by installing its associated
MatProGo.dev pacakage and then calling its "Solver" object.

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