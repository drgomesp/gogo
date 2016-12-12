package tests

import (
	"reflect"
	"testing"

	"github.com/drgomesp/gogo/src/generator"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateField(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := &generator.StructFieldGenerator{
			Name: "FieldA",
			Kind: reflect.Bool,
		}

		Convey("Then a struct field should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, "FieldA bool")
		})
	})
}
