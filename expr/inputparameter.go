package expr

import (
	"fmt"
	//"goa.design/model/expr"
)

type (
	// Method represents a method.
	InputParameter struct {
		*Element
		Method *Method // pointer to "pseudo-parent
	}

	// Components is a slice of components that can be easily converted into
	// a slice of ElementHolder.
	InputParameters []*InputParameter
)

// EvalName returns the generic expression name used in error messages.
func (p *InputParameter) EvalName() string {
	if p.Name == "" {
		return "unnamed method"
	}
	return fmt.Sprintf("method %q", p.Name)
}

// Finalize adds the 'Component' tag ands finalizes relationships.
func (p *InputParameter) Finalize() {
	p.PrefixTags("Element", "Method")
	p.Element.Finalize()
}

// Elements returns a slice of ElementHolder that contains the elements of c.
func (ps InputParameters) Elements() []ElementHolder {
	res := make([]ElementHolder, len(ps))
	for i, p := range ps {
		res[i] = p
	}
	return res
}
