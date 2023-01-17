package dsl

import (
	//"fmt"
	"fmt"
	"strings"

	"go_kickstart/expr"

	"goa.design/goa/v3/eval"
)

// Struct defines a golang struct.
//
// Struct must appear in a Design expression.
//
// Struct takes 1 to 3 arguments. The first argument is the struct
// name and the last argument a function that contains the expressions
// that defines the content of the struct. An optional description may be given
// after the name.
//
// The valid syntax for Struct is thus:
//
//	Struct("<name>")
//
//	Struct("<name>", "[description]")
//
//	Struct("<name>", func())
//
//	Struct("<name>", "[description]", func())
func Struct(name string, args ...interface{}) *expr.Struct {
	component, ok := eval.Current().(*expr.Component)
	if !ok {
		eval.IncompatibleDSL()
		return nil
	}
	if strings.Contains(name, "/") {
		eval.ReportError("Struct: name cannot include slashes")
	}
	description, technology, dsl, err := parseElementArgs(args...)
	if err != nil {
		eval.ReportError("Struct: " + err.Error())
		return nil
	}
	c := &expr.Struct{
		Element: &expr.Element{
			Name:        name,
			Description: description,
			Technology:  technology,
			DSLFunc:     dsl,
		},
		Component: component,
	}
	return component.AddStruct(c)
}

func Interface(name string, args ...interface{}) *expr.Interface {
	component, ok := eval.Current().(*expr.Component)
	if !ok {
		eval.IncompatibleDSL()
		return nil
	}
	if strings.Contains(name, "/") {
		eval.ReportError("Interface: name cannot include slashes")
	}
	description, technology, dsl, err := parseElementArgs(args...)
	if err != nil {
		eval.ReportError("Interface: " + err.Error())
		return nil
	}
	c := &expr.Interface{
		Element: &expr.Element{
			Name:        name,
			Description: description,
			Technology:  technology,
			DSLFunc:     dsl,
		},
		Component: component,
	}
	return component.AddInterface(c)
}

// Method defines a golang method.
//
// Method must appear in a Struct, Interface or component expression.
//
// Method takes 1 to 3 arguments. The first argument is the method
// name and the last argument a function that contains the expressions
// that defines the parameters of the method. An optional description may be given
// after the name.
//
// The valid syntax for Method is thus:
//
//	Method("<name>")
//
//	Method("<name>", "[description]")
//
//	Method("<name>", func())
//
//	Method("<name>", "[description]", func())
func Method(name string, args ...interface{}) *expr.Method {
	var bIsStruct bool = false
	var bIsInterface bool = false
	var mstruct *expr.Struct
	var minterface *expr.Interface
	var ok bool
	mstruct, ok = eval.Current().(*expr.Struct)
	if ok {
		bIsStruct = true
	} else {
		minterface, ok = eval.Current().(*expr.Interface)
		if ok {
			bIsInterface = true
		} else {
			fmt.Printf("Method not in correct context\n")
		}
	}
	if !ok {
		eval.IncompatibleDSL()
		return nil
	}
	if strings.Contains(name, "/") {
		eval.ReportError("Method: name cannot include slashes")
	}
	description, technology, dsl, err := parseElementArgs(args...)
	if err != nil {
		eval.ReportError("Method: " + err.Error())
		return nil
	}
	if bIsStruct {
		c := &expr.Method{
			Element: &expr.Element{
				Name:        name,
				Description: description,
				Technology:  technology,
				DSLFunc:     dsl,
			},
			Struct: mstruct,
		}
		return mstruct.AddMethod(c)
	}
	if bIsInterface {
		c := &expr.Method{
			Element: &expr.Element{
				Name:        name,
				Description: description,
				Technology:  technology,
				DSLFunc:     dsl,
			},
			Interface: minterface,
		}
		return minterface.AddMethod(c)
	}
	return nil
}

// Method defines a golang method.
//
// Method must appear in a Struct, Interface or component expression.
//
// Method takes 1 to 3 arguments. The first argument is the method
// name and the last argument a function that contains the expressions
// that defines the parameters of the method. An optional description may be given
// after the name.
//
// The valid syntax for Method is thus:
//
//	Method("<name>")
//
//	Method("<name>", "[description]")
//
//	Method("<name>", func())
//
//	Method("<name>", "[description]", func())
func InputParameter(name string, args ...interface{}) *expr.InputParameter {
	var bIsStruct bool = false
	//var bIsInterface bool = false
	var method *expr.Method
	//var minterface *expr.Interface
	var ok bool
	method, ok = eval.Current().(*expr.Method)
	if ok {
		bIsStruct = true
	} else {
		fmt.Printf("InputParameter not in correct context\n")
	}

	if !ok {
		eval.IncompatibleDSL()
		return nil
	}
	if strings.Contains(name, "/") {
		eval.ReportError("InputParameter: name cannot include slashes")
	}
	description, technology, dsl, err := parseElementArgs(args...)
	if err != nil {
		eval.ReportError("InputParameter: " + err.Error())
		return nil
	}
	if bIsStruct {
		c := &expr.InputParameter{
			Element: &expr.Element{
				Name:        name,
				Description: description,
				Technology:  technology,
				DSLFunc:     dsl,
			},
			Method: method,
		}
		return method.AddInputParameter(c)
	}
	return nil
}

func ReturnParameter(name string, args ...interface{}) *expr.ReturnParameter {
	var bIsStruct bool = false
	//var bIsInterface bool = false
	var method *expr.Method
	//var minterface *expr.Interface
	var ok bool
	method, ok = eval.Current().(*expr.Method)
	if ok {
		bIsStruct = true
	} else {
		fmt.Printf("ReturnParameter not in correct context\n")
	}

	if !ok {
		eval.IncompatibleDSL()
		return nil
	}
	if strings.Contains(name, "/") {
		eval.ReportError("ReturnParameter: name cannot include slashes")
	}
	description, technology, dsl, err := parseElementArgs(args...)
	if err != nil {
		eval.ReportError("ReturnParameter: " + err.Error())
		return nil
	}
	if bIsStruct {
		c := &expr.ReturnParameter{
			Element: &expr.Element{
				Name:        name,
				Description: description,
				Technology:  technology,
				DSLFunc:     dsl,
			},
			Method: method,
		}
		return method.AddReturnParameter(c)
	}
	return nil
}
