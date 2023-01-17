package expr

import (
	"fmt"
	"strings"
	//"goa.design/model/expr"
)

type (
	// Method represents a method.
	Method struct {
		*Element
		InputParameters  []*InputParameter  // what it contains
		ReturnParameters []*ReturnParameter // what it contains
		Struct           *Struct            // pointer to "pseudo-parent
		Interface        *Interface         // pointer to "pseudo-parent
	}

	// Components is a slice of components that can be easily converted into
	// a slice of ElementHolder.
	Methods []*Method
)

// EvalName returns the generic expression name used in error messages.
func (s *Method) EvalName() string {
	if s.Name == "" {
		return "unnamed method"
	}
	return fmt.Sprintf("method %q", s.Name)
}

// Finalize adds the 'Component' tag ands finalizes relationships.
func (s *Method) Finalize() {
	s.PrefixTags("Element", "Method")
	s.Element.Finalize()
}

// Elements returns a slice of ElementHolder that contains the elements of c.
func (cs Methods) Elements() []ElementHolder {
	res := make([]ElementHolder, len(cs))
	for i, cc := range cs {
		res[i] = cc
	}
	return res
}

func (m *Method) FindInputParameterName(name string) *InputParameter {
	for _, ip := range m.InputParameters {
		if ip.Name == name {
			return ip
		}
	}
	return nil
}

func (m *Method) FindReturnParameterName(name string) *ReturnParameter {
	for _, ip := range m.ReturnParameters {
		if ip.Name == name {
			return ip
		}
	}
	return nil
}

// AddInputParameter adds the given method to the struct. If there is already a
// method with the given name then AddInputParameter merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddInputParameter returns the new or merged method.
func (m *Method) AddInputParameter(cmp *InputParameter) *InputParameter {
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

// AddInputParameter adds the given method to the struct. If there is already a
// method with the given name then AddInputParameter merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddInputParameter returns the new or merged method.
func (m *Method) AddReturnParameter(cmp *ReturnParameter) *ReturnParameter {
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
