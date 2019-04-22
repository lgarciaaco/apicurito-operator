package controller

import (
	"github.com/apicurito-operator/pkg/controller/apicurito"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, apicurito.Add)
}
