package generator

import (
	"fmt"
	"testing"

	"reflect"

	"github.com/drgomesp/gogo/meta"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateEmptyMethod(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := StructMethodGenerator{
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
		gen := StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
			Parameters: []meta.Parameter{
				meta.Parameter{Name: "first", Type: reflect.Bool},
				meta.Parameter{Name: "second", Type: reflect.Int},
				meta.Parameter{Name: "third", Type: reflect.String},
			},
		}

		Convey("Then a struct method should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, fmt.Sprintf("func (r *ReceiverStructName) MethodName(first bool, second int, third string) {}"))
		})
	})
}

func TestGenerateEmptyMethodWithParametersAndSingleReturn(t *testing.T) {
	Convey("Given a struct field generator instance", t, func() {
		gen := StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
			Parameters: []meta.Parameter{
				meta.Parameter{Name: "first", Type: reflect.Bool},
				meta.Parameter{Name: "second", Type: reflect.Int},
				meta.Parameter{Name: "third", Type: reflect.String},
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
		gen := StructMethodGenerator{
			Receiver: "ReceiverStructName",
			Name:     "MethodName",
			Parameters: []meta.Parameter{
				meta.Parameter{Name: "first", Type: reflect.Bool},
				meta.Parameter{Name: "second", Type: reflect.Int},
				meta.Parameter{Name: "third", Type: reflect.String},
			},
			Returns: []meta.Parameter{
				meta.Parameter{Name: "first", Type: reflect.Bool},
				meta.Parameter{Name: "second", Type: reflect.Int},
			},
		}

		Convey("Then a struct method should be generated correctly", func() {
			So(gen.Generate(), ShouldEqual, fmt.Sprintf("func (r *ReceiverStructName) MethodName(first bool, second int, third string) (first bool, second int) {}"))
		})
	})
}
