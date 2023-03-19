# MatProInterface.go
A common interface used for modeling Mathematical Programs in the language Go.

## To-Dos

* [ ] Create New AddConstr methods which work for vector constraints
* [ ] Mult
  * [ ] General Function (in operators.go)
    * [ ] Methods for
        * [ ] Scalars
            * [ ] Constant
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
* [ ] VectorConstraint
    * [ ] AtVec()