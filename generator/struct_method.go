package generator

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/drgomesp/gogo/meta"
)

// StructMethodGenerator is a helper for generating struct methods
type StructMethodGenerator struct {
	Receiver   string
	Name       string
	Parameters []meta.Parameter
	Returns    interface{}
	Body       string
}

// Generate returns the string representation of a struct method
func (g *StructMethodGenerator) Generate() string {
	return fmt.Sprintf("func (r *%s) %s(%s)%s {}", g.Receiver, g.Name, g.parameters(), g.returns())
}

// Parameters string representation
func (g *StructMethodGenerator) parameters() string {
	if len(g.Parameters) == 0 {
		return ""
	}

	var buffer bytes.Buffer

	for i := 0; i < len(g.Parameters); i++ {
		buffer.WriteString(fmt.Sprintf("%s %s", g.Parameters[i].Name, g.Parameters[i].Type))

		if i != len(g.Parameters)-1 {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}

// Return string representation
func (g *StructMethodGenerator) returns() string {
	if r, ok := g.Returns.(reflect.Kind); ok {
		return fmt.Sprintf(" %s", r.String())
	}

	var buffer bytes.Buffer

	if returns, ok := g.Returns.([]meta.Parameter); ok {
		buffer.WriteString(" (")

		for i := 0; i < len(returns); i++ {
			buffer.WriteString(fmt.Sprintf("%s %s", returns[i].Name, returns[i].Type))

			if i != len(returns)-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString(")")
	}

	return buffer.String()
}
