package controller

import (
	"github.com/example-inc/killer-operator/pkg/controller/podkiller"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, podkiller.Add)
}
