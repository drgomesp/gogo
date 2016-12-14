package generator

import "fmt"
import "bytes"

// StructGenerator is a helper for generating structs
type StructGenerator struct {
	Name    string
	Fields  []StructFieldGenerator
	Methods []StructMethodGenerator
}

// Generate returns the string representation of a struct
func (g *StructGenerator) Generate() string {
	return fmt.Sprintf(`type %s struct {%s}%s`, g.Name, g.fields(), g.methods())
}

func (g *StructGenerator) fields() string {
	if len(g.Fields) > 0 {
		var buffer bytes.Buffer

		for i := 0; i < len(g.Fields); i++ {
			buffer.WriteString(fmt.Sprintf(`
	%s`, g.Fields[i].Generate()))
		}

		buffer.WriteString(`
`)

		return buffer.String()
	}

	return ""
}

func (g *StructGenerator) methods() string {
	if len(g.Methods) > 0 {
		var buffer bytes.Buffer

		buffer.WriteString("\n")

		for i := 0; i < len(g.Methods); i++ {
			if len(g.Methods[i].Receiver) == 0 {
				g.Methods[i].Receiver = g.Name
			}

			buffer.WriteString(fmt.Sprintf("\n%s", g.Methods[i].Generate()))
		}

		return buffer.String()
	}

	return ""
}
