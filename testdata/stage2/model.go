package design

//import . "goa.design/model/dsl"

import . "go_kickstart/dsl"

/*
Nesting to the level envisaged in the DSL created for the model..
Extended to include the struct, interface and method levels..
doesn't include parameters
*/

var _ = Design("Getting Started", "This is a model of vernon's software system.", func() {
	SoftwareSystem("Software System", "Vernon's software system.", func() {
		Container("KMContainer", "KM container with real content", "docker & golang", func() {
			// can have Tag, URL, Uses, Delivers inside this
			Component("KM component", "A custom component", "Go and Python", func() {
				Struct("KM Struct", "A custom struct", "Go only", func() {
					Method("KM method", "Means of introspection", "Go only", func() {})
				})
				Interface("KM Interface", "A custom interface", "Go only", func() {
					Method("KM method", "Means of reflection", "Go only", func() {})
				})

			})
		})
	})

})
