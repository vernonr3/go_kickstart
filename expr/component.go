package expr

import (
	"fmt"
	"strings"
)

type (
	// Component represents a component.
	Component struct {
		*Element
		Structs    []*Struct
		Interfaces []*Interface
		Functions  []*Function
		Container  *Container
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

// Struct returns the struct with the given name if any, nil otherwise.
func (c *Component) Struct(name string) *Struct {
	for _, cc := range c.Structs {
		if cc.Name == name {
			return cc
		}
	}
	return nil
}

// Function returns the function with the given name if any, nil otherwise.
func (c *Component) Function(name string) *Function {
	for _, cc := range c.Functions {
		if cc.Name == name {
			return cc
		}
	}
	return nil
}

// AddStruct adds the given structt to the component. If there is already a
// struct with the given name then AddStruct merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddStruct returns the new or merged struct.
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

// Component returns the component with the given name if any, nil otherwise.
func (c *Component) Interface(name string) *Interface {
	for _, cc := range c.Interfaces {
		if cc.Name == name {
			return cc
		}
	}
	return nil
}

// AddInterface adds the given interface to the component. If there is already a
// interface with the given name then AddInterface merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddInterface returns the new or merged struct.
func (c *Component) AddInterface(cmp *Interface) *Interface {
	existing := c.Interface(cmp.Name)
	if existing == nil {
		Identify(cmp)
		c.Interfaces = append(c.Interfaces, cmp)
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

// AddFunction adds the given function to the component. If there is already a
// function with the given name then AddFunction merges both definitions. The
// merge algorithm:
//
//   - overrides the description, technology and URL if provided,
//   - merges any new tag or propery into the existing tags and properties,
//
// AddFunction returns the new or merged function.
func (c *Component) AddFunction(cmp *Function) *Function {
	existing := c.Function(cmp.Name)
	if existing == nil {
		Identify(cmp)
		c.Functions = append(c.Functions, cmp)
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
