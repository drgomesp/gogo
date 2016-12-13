package generator

import (
	"bytes"
	"fmt"
	"reflect"
)

// StructMethodGenerator is a helper for generating struct methods
type StructMethodGenerator struct {
	Receiver   string
	Name       string
	Parameters map[string]reflect.Kind
	Returns    interface{}
	Body       string
}

// Generate returns the string representation of a struct method
func (g *StructMethodGenerator) Generate() string {
	return fmt.Sprintf("func (r *%s) %s%s%s {}", g.Receiver, g.Name, g.parameters(), g.returns())
}

// Parameters string representation
func (g *StructMethodGenerator) parameters() string {
	var buffer bytes.Buffer

	buffer.WriteString("(")

	counter := 1
	for name, t := range g.Parameters {
		buffer.WriteString(fmt.Sprintf("%s %s", name, t))

		if counter != len(g.Parameters) {
			buffer.WriteString(", ")
		}

		counter++
	}

	buffer.WriteString(")")

	return buffer.String()
}

// Return string representation
func (g *StructMethodGenerator) returns() string {
	if r, ok := g.Returns.(reflect.Kind); ok {
		return fmt.Sprintf(" %s", r.String())
	}

	var buffer bytes.Buffer

	if r, ok := g.Returns.(map[string]reflect.Kind); ok {
		buffer.WriteString(" (")

		counter := 1
		for name, t := range r {
			buffer.WriteString(fmt.Sprintf("%s %s", name, t))

			if counter != len(r) {
				buffer.WriteString(", ")
			}

			counter++
		}

		buffer.WriteString(")")
	}

	return buffer.String()
}
