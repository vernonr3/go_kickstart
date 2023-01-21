package expr

import (
	"fmt"
	"strings"
	//"goa.design/model/expr"
)

type (
	// Component represents a component.
	Interface struct {
		*Element
		Methods   Methods    // what it contains
		Component *Component // pointer to "pseudo-parent
		Struct    *Struct    // pointer to "pseudo-parent
	}

	// Components is a slice of components that can be easily converted into
	// a slice of ElementHolder.
	Interfaces []*Interface
)

// EvalName returns the generic expression name used in error messages.
func (s *Interface) EvalName() string {
	if s.Name == "" {
		return "unnamed interface"
	}
	return fmt.Sprintf("interface %q", s.Name)
}

// Finalize adds the 'Component' tag ands finalizes relationships.
func (s *Interface) Finalize() {
	s.PrefixTags("Element", "Interface")
	s.Element.Finalize()
}

// Elements returns a slice of ElementHolder that contains the elements of c.
func (cs Interfaces) Elements() []ElementHolder {
	res := make([]ElementHolder, len(cs))
	for i, cc := range cs {
		res[i] = cc
	}
	return res
}

func (s *Interface) Method(name string) *Method {
	for _, cc := range s.Methods {
		if cc.Name == name {
			return cc
		}
	}
	return nil
}

// AddMethod adds the given method to the struct. If there is already a
// method with the given name then AddMethod merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddMethod returns the new or merged method.
func (s *Interface) AddMethod(cmp *Method) *Method {
	existing := s.Method(cmp.Name)
	if existing == nil {
		Identify(cmp)
		s.Methods = append(s.Methods, cmp)
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
