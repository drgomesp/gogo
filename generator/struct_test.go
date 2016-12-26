package generator

import (
	"testing"

	"reflect"

	"github.com/drgomesp/gogo/meta"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratingEmptyStruct(t *testing.T) {
	Convey("Given a struct generator instance for an empty struct", t, func() {
		gen := StructGenerator{
			Name: "Foo",
		}

		Convey("Then an empty struct should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, `type Foo struct {}`)
		})
	})
}

func TestGeneratingStructWithFields(t *testing.T) {
	Convey("Given a struct generator instance for a struct with fields", t, func() {
		gen := StructGenerator{
			Name: "Foo",
			Fields: []StructFieldGenerator{
				StructFieldGenerator{
					Name: "FieldA",
					Kind: reflect.Bool,
				},
				StructFieldGenerator{
					Name: "FieldB",
					Kind: reflect.Int,
				},
				StructFieldGenerator{
					Name: "FieldC",
					Kind: reflect.String,
				},
			},
		}

		Convey("Then an empty struct should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, `type Foo struct {
	FieldA bool
	FieldB int
	FieldC string
}`)
		})
	})
}

func TestGeneratingStructWithMethods(t *testing.T) {
	Convey("Given a struct generator instance for struct with methods", t, func() {
		gen := StructGenerator{
			Name: "Foo",
			Methods: []StructMethodGenerator{
				StructMethodGenerator{
					Name: "Bar",
					Parameters: []meta.Parameter{
						meta.Parameter{Name: "first", Type: reflect.Bool},
						meta.Parameter{Name: "second", Type: reflect.Int},
						meta.Parameter{Name: "third", Type: reflect.String},
					},
					Returns: []meta.Parameter{
						meta.Parameter{Name: "first", Type: reflect.Bool},
						meta.Parameter{Name: "second", Type: reflect.Int},
					},
				},
			},
		}

		Convey("Then an empty struct should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, `type Foo struct {}

func (r *Foo) Bar(first bool, second int, third string) (first bool, second int) {}`)
		})
	})
}

func TestGeneratingStructWithFieldsAndMethods(t *testing.T) {
	Convey("Given a struct generator instance for struct with fields and methods", t, func() {
		gen := StructGenerator{
			Name: "Foo",
			Fields: []StructFieldGenerator{
				StructFieldGenerator{
					Name: "FieldA",
					Kind: reflect.Bool,
				},
				StructFieldGenerator{
					Name: "FieldB",
					Kind: reflect.Int,
				},
				StructFieldGenerator{
					Name: "FieldC",
					Kind: reflect.String,
				},
			},
			Methods: []StructMethodGenerator{
				StructMethodGenerator{
					Name: "Bar",
					Parameters: []meta.Parameter{
						meta.Parameter{Name: "first", Type: reflect.Bool},
						meta.Parameter{Name: "second", Type: reflect.Int},
						meta.Parameter{Name: "third", Type: reflect.String},
					},
					Returns: []meta.Parameter{
						meta.Parameter{Name: "first", Type: reflect.Bool},
						meta.Parameter{Name: "second", Type: reflect.Int},
					},
				},
			},
		}

		Convey("Then an empty struct should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, `type Foo struct {
	FieldA bool
	FieldB int
	FieldC string
}

func (r *Foo) Bar(first bool, second int, third string) (first bool, second int) {}`)
		})
	})
}
