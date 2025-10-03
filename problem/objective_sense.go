package problem

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
)

// ObjSense represents whether an optimization objective is to be maximized or
// minimized. This implementation conforms to the Gurobi encoding
type ObjSense string

// Objective senses (minimize and maximize) encoding using Gurobi's standard
const (
	SenseMinimize ObjSense = "Minimize"
	SenseMaximize          = "Maximize"
	SenseFind     ObjSense = "Find"
)

/*
ToObjSense
Description:

	This method converts an input optim.ObjSense to a problem.ObjSense.
*/
func ToObjSense(sense optim.ObjSense) ObjSense {
	if sense == optim.SenseMinimize {
		return SenseMinimize
	} else {
		return SenseMaximize
	}
}
