package optim

/*
model_test.go
Description:
	This script tests the model object.
*/

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
TestModel_NewModel1
Description:

	Tests the new model was initialized in the proper way.
*/
func TestModel_NewModel1(t *testing.T) {
	// Constants
	m := optim.NewModel("test")

	// Algorithm
	if m.ShowLog {
		t.Errorf("ShowLog should be initialized as false.")
	}

	if m.Name != "test" {
		t.Errorf("Expected model's name to be %v; received %v", "test", m.Name)
	}
}
