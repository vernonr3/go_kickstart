package expr

import (
	"fmt"
	//"goa.design/model/expr"
)

type (
	// Method represents a method.
	Method struct {
		*Element
		//Parameters  Parameters	 // what it contains
		Struct    *Struct    // pointer to "pseudo-parent
		Interface *Interface // pointer to "pseudo-parent
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
