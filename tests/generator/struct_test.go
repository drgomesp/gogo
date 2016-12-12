package tests

import (
	"testing"

	"reflect"

	"github.com/drgomesp/gogo/src/generator"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratingEmptyStruct(t *testing.T) {
	Convey("Given a struct generator instance for an empty struct", t, func() {
		g := &generator.StructGenerator{
			Name: "Foo",
		}

		Convey("Then an empty struct should be generated correctly", func() {
			So(g.Generate(), ShouldEqual, `type Foo struct {}`)
		})
	})
}

func TestGeneratingStructWithFields(t *testing.T) {
	Convey("Given a struct generator instance for an empty struct", t, func() {
		g := &generator.StructGenerator{
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
			So(g.Generate(), ShouldEqual, `type Foo struct {
	FieldA bool
	FieldB int
	FieldC string
}`)
		})
	})
}
