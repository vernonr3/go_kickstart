/*
package MDL is designed to incorporate the model in goa.design/model.
The following types here are drawn from the mdl folder:

  - Person
  - Software System
  - Container

Component is extended to refer to structs and interfaces.
The following are additional types
  - Struct
  - Interface
  - Method
  - Function
*/
package mdl

type (
	Struct struct {
		// ID of element.
		ID string `json:"id"`
		// Name of element - not applicable to ContainerInstance.
		Name string `json:"name,omitempty"`
		// Description of element if any.
		Description string `json:"description,omitempty"`
		// Technology used by element if any - not applicable to Person.
		Technology string `json:"technology,omitempty"`
		// Tags attached to element as comma separated list if any.
		Tags string `json:"tags,omitempty"`
		// URL where more information about this element can be found.
		URL string `json:"url,omitempty"`
		// Set of arbitrary name-value properties (shown in diagram tooltips).
		Properties map[string]string `json:"properties,omitempty"`
		// Relationships is the set of relationships from this element to other
		// elements.
		Relationships []*Relationship `json:"relationships,omitempty"`

		Methods []*Method
	}

	Interface struct {
		// ID of element.
		ID string `json:"id"`
		// Name of element - not applicable to ContainerInstance.
		Name string `json:"name,omitempty"`
		// Description of element if any.
		Description string `json:"description,omitempty"`
		// Technology used by element if any - not applicable to Person.
		Technology string `json:"technology,omitempty"`
		// Tags attached to element as comma separated list if any.
		Tags string `json:"tags,omitempty"`
		// URL where more information about this element can be found.
		URL string `json:"url,omitempty"`
		// Set of arbitrary name-value properties (shown in diagram tooltips).
		Properties map[string]string `json:"properties,omitempty"`
		// Relationships is the set of relationships from this element to other
		// elements.
		Relationships []*Relationship `json:"relationships,omitempty"`

		Methods []*Method
	}
	Method struct {
		// ID of element.
		ID string `json:"id"`
		// Name of element - not applicable to ContainerInstance.
		Name string `json:"name,omitempty"`
		// Description of element if any.
		Description string `json:"description,omitempty"`
		// Technology used by element if any - not applicable to Person.
		Technology string `json:"technology,omitempty"`
		// Tags attached to element as comma separated list if any.
		Tags string `json:"tags,omitempty"`
		// URL where more information about this element can be found.
		URL string `json:"url,omitempty"`
		// Set of arbitrary name-value properties (shown in diagram tooltips).
		Properties map[string]string `json:"properties,omitempty"`
		// Relationships is the set of relationships from this element to other
		// elements.
		Relationships []*Relationship `json:"relationships,omitempty"`
		// Input parameters are the variables supplied to this method
		InputParameters []*InputParameter
		// Return parameters are the variables returned by this method
		ReturnParameters []*ReturnParameter
	}
	InputParameter struct {
		// ID of element.
		ID string `json:"id"`
		// Name of element - not applicable to ContainerInstance.
		Name string `json:"name,omitempty"`
		// Description of element if any.
		Description string `json:"type,omitempty"`
		// Technology used by element if any - not applicable to Person.
		Technology string `json:"technology,omitempty"`
		// Tags attached to element as comma separated list if any.
		Tags string `json:"tags,omitempty"`
		// URL where more information about this element can be found.
		URL string `json:"url,omitempty"`
		// Set of arbitrary name-value properties (shown in diagram tooltips).
		Properties map[string]string `json:"properties,omitempty"`
		// Relationships is the set of relationships from this element to other
		// elements.
		Relationships []*Relationship `json:"relationships,omitempty"`
	}

	ReturnParameter struct {
		// ID of element.
		ID string `json:"id"`
		// Name of element - not applicable to ContainerInstance.
		Name string `json:"name,omitempty"`
		// Description of element if any.
		Description string `json:"type,omitempty"`
		// Technology used by element if any - not applicable to Person.
		Technology string `json:"technology,omitempty"`
		// Tags attached to element as comma separated list if any.
		Tags string `json:"tags,omitempty"`
		// URL where more information about this element can be found.
		URL string `json:"url,omitempty"`
		// Set of arbitrary name-value properties (shown in diagram tooltips).
		Properties map[string]string `json:"properties,omitempty"`
		// Relationships is the set of relationships from this element to other
		// elements.
		Relationships []*Relationship `json:"relationships,omitempty"`
	}
)
