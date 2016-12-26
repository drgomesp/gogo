package generator

import (
	"fmt"
	"reflect"
)

// StructFieldGenerator is a helper for generating struct fields
type StructFieldGenerator struct {
	Name string
	Kind reflect.Kind
}

// Generate returns the string representation of a struct field
func (g *StructFieldGenerator) Generate() string {
	return fmt.Sprintf(`%s %s`, g.Name, g.Kind.String())
}
