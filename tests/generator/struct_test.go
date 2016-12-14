package tests

import (
	"testing"

	"reflect"

	"github.com/drgomesp/gogo/src/generator"
	"github.com/drgomesp/gogo/src/meta"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratingEmptyStruct(t *testing.T) {
	Convey("Given a struct generator instance for an empty struct", t, func() {
		gen := generator.StructGenerator{
			Name: "Foo",
		}

		Convey("Then an empty struct should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, `type Foo struct {}`)
		})
	})
}

func TestGeneratingStructWithFields(t *testing.T) {
	Convey("Given a struct generator instance for a struct with fields", t, func() {
		gen := generator.StructGenerator{
			Name: "Foo",
			Fields: []generator.StructFieldGenerator{
				generator.StructFieldGenerator{
					Name: "FieldA",
					Kind: reflect.Bool,
				},
				generator.StructFieldGenerator{
					Name: "FieldB",
					Kind: reflect.Int,
				},
				generator.StructFieldGenerator{
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
		gen := generator.StructGenerator{
			Name: "Foo",
			Methods: []generator.StructMethodGenerator{
				generator.StructMethodGenerator{
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
