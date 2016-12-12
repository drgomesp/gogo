package generator

// Generator defines a basic interface for code generation
type Generator interface {
	Generate() string
}

