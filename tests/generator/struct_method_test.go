package tests

import (
	"fmt"
	"testing"

	"reflect"

	"github.com/drgomesp/gogo/src/generator"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateEmptyMethod(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := &generator.StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
		}

		Convey("Then a struct method should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, fmt.Sprintf("func (r *ReceiverStructName) MethodName() {}"))
		})
	})
}

func TestGenerateEmptyMethodWithParameters(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := &generator.StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
			Parameters: map[string]reflect.Kind{
				"first":  reflect.Bool,
				"second": reflect.Int,
				"third":  reflect.String,
			},
		}

		Convey("Then a struct method should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, fmt.Sprintf("func (r *ReceiverStructName) MethodName(first bool, second int, third string) {}"))
		})
	})
}

func TestGenerateEmptyMethodWithParametersAndSingleReturn(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := &generator.StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
			Parameters: map[string]reflect.Kind{
				"first":  reflect.Bool,
				"second": reflect.Int,
				"third":  reflect.String,
			},
			Returns: reflect.String,
		}

		Convey("Then a struct method should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, fmt.Sprintf("func (r *ReceiverStructName) MethodName(first bool, second int, third string) string {}"))
		})
	})
}

func TestGenerateEmptyMethodWithParametersAndDoubleReturn(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := &generator.StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
			Parameters: map[string]reflect.Kind{
				"first":  reflect.Bool,
				"second": reflect.Int,
				"third":  reflect.String,
			},
			Returns: map[string]reflect.Kind{
				"first":  reflect.Bool,
				"second": reflect.Int,
			},
		}

		Convey("Then a struct method should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, fmt.Sprintf("func (r *ReceiverStructName) MethodName(first bool, second int, third string) (first bool, second int) {}"))
		})
	})
}
