package generator

import "fmt"
import "bytes"

// StructGenerator is a helper for generating structs
type StructGenerator struct {
	Name   string
	Fields []StructFieldGenerator
}

// Generate returns the string representation of a struct
func (g *StructGenerator) Generate() string {
	return fmt.Sprintf(`type %s struct {%s}`, g.Name, (func() string {
		var buffer bytes.Buffer

		if len(g.Fields) > 0 {
			for _, fieldGenerator := range g.Fields {
				buffer.WriteString(fmt.Sprintf(`
	%s`, fieldGenerator.Generate()))
			}

			buffer.WriteString(`
`)
		}

		return buffer.String()
	})())
}
