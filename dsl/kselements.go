package dsl

import (
	//"fmt"
	"strings"

	gexpr "go_kickstart/expr"

	"goa.design/goa/v3/eval"
)

func Struct(name string, args ...interface{}) *gexpr.Struct {
	component, ok := eval.Current().(*gexpr.Component)
	if !ok {
		eval.IncompatibleDSL()
		return nil
	}
	if strings.Contains(name, "/") {
		eval.ReportError("Component: name cannot include slashes")
	}
	description, technology, dsl, err := parseElementArgs(args...)
	if err != nil {
		eval.ReportError("Component: " + err.Error())
		return nil
	}
	c := &gexpr.Struct{
		Element: &gexpr.Element{
			Name:        name,
			Description: description,
			Technology:  technology,
			DSLFunc:     dsl,
		},
		Component: component,
	}
	return component.AddStruct(c)
}
