package expr

import (
	"fmt"
	"strings"
	//"goa.design/model/expr"
)

type (
	// Function represents a function.
	Function struct {
		*Element
		InputParameters  []*InputParameter  // what it contains
		ReturnParameters []*ReturnParameter // what it contains
		Component        *Component         // pointer to "pseudo-parent
	}

	// Functions is a slice of functions that can be easily converted into
	// a slice of ElementHolder.
	Functions []*Function
)

// EvalName returns the generic expression name used in error messages.
func (s *Function) EvalName() string {
	if s.Name == "" {
		return "unnamed function"
	}
	return fmt.Sprintf("function %q", s.Name)
}

// Finalize adds the 'Function' tag ands finalizes relationships.
func (s *Function) Finalize() {
	s.PrefixTags("Element", "Function")
	s.Element.Finalize()
}

// Elements returns a slice of ElementHolder that contains the elements of c.
func (cs Functions) Elements() []ElementHolder {
	res := make([]ElementHolder, len(cs))
	for i, cc := range cs {
		res[i] = cc
	}
	return res
}

func (m *Function) FindInputParameterName(name string) *InputParameter {
	for _, ip := range m.InputParameters {
		if ip.Name == name {
			return ip
		}
	}
	return nil
}

func (m *Function) FindReturnParameterName(name string) *ReturnParameter {
	for _, ip := range m.ReturnParameters {
		if ip.Name == name {
			return ip
		}
	}
	return nil
}

// AddInputParameter adds the given Function to the component. If there is already a
// function with the given name then AddInputParameter merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddInputParameter returns the new or merged function.
func (m *Function) AddInputParameter(cmp *InputParameter) *InputParameter {
	existing := m.FindInputParameterName(cmp.Name)
	if existing == nil {
		Identify(cmp)
		m.InputParameters = append(m.InputParameters, cmp)
		return cmp
	}
	if cmp.Description != "" {
		existing.Description = cmp.Description
	}
	if cmp.Technology != "" {
		existing.Technology = cmp.Technology
	}
	if cmp.URL != "" {
		existing.URL = cmp.URL
	}
	existing.MergeTags(strings.Split(cmp.Tags, ",")...)
	if olddsl := existing.DSLFunc; olddsl != nil {
		existing.DSLFunc = func() { olddsl(); cmp.DSLFunc() }
	}
	return existing
}

// AddInputParameter adds the given function to the component. If there is already a
// function with the given name then AddInputParameter merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddInputParameter returns the new or merged function.
func (m *Function) AddReturnParameter(cmp *ReturnParameter) *ReturnParameter {
	existing := m.FindReturnParameterName(cmp.Name)
	if existing == nil {
		Identify(cmp)
		m.ReturnParameters = append(m.ReturnParameters, cmp)
		return cmp
	}
	if cmp.Description != "" {
		existing.Description = cmp.Description
	}
	if cmp.Technology != "" {
		existing.Technology = cmp.Technology
	}
	if cmp.URL != "" {
		existing.URL = cmp.URL
	}
	existing.MergeTags(strings.Split(cmp.Tags, ",")...)
	if olddsl := existing.DSLFunc; olddsl != nil {
		existing.DSLFunc = func() { olddsl(); cmp.DSLFunc() }
	}
	return existing
}
