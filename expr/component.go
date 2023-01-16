package expr

import (
	"fmt"
)

type (
	// Struct represents a Struct.
	Component struct {
		Element
		//Structs   Structs
		Container *Container
	}

	// Components is a slice of components that can be easily converted into
	// a slice of ElementHolder.
	Components []*Component
)

// EvalName returns the generic expression name used in error messages.
func (c *Component) EvalName() string {
	if c.Name == "" {
		return "unnamed component"
	}
	return fmt.Sprintf("component %q", c.Name)
}

// Finalize adds the 'Component' tag ands finalizes relationships.
func (c *Component) Finalize() {
	c.PrefixTags("Element", "Component")
	c.Element.Finalize()
}

// Elements returns a slice of ElementHolder that contains the elements of c.
func (cs Components) Elements() []ElementHolder {
	res := make([]ElementHolder, len(cs))
	for i, cc := range cs {
		res[i] = cc
	}
	return res
}

/*
// Component returns the component with the given name if any, nil otherwise.
func (c *Component) Struct(name string) *Struct {
	for _, cc := range c.Structs {
		if cc.Name == name {
			return cc
		}
	}
	return nil
}

// AddComponent adds the given component to the container. If there is already a
// component with the given name then AddComponent merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddComponent returns the new or merged component.
func (c *Component) AddStruct(cmp *Struct) *Struct {
	existing := c.Struct(cmp.Name)
	if existing == nil {
		Identify(cmp)
		c.Structs = append(c.Structs, cmp)
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
*/
