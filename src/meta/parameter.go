package meta

import "reflect"

// Parameter represents a function parameter composed by a name and a type (reflect.Kind)
type Parameter struct {
	Name string
	Type reflect.Kind
}
