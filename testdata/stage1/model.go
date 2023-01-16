package design

//import . "goa.design/model/dsl"

import . "go_kickstart/dsl"

/*
uses nesting to the level envisaged in the DSL created for the model..
Next stage is to extend this beyond into the levels I want..
*/

var _ = Design("Getting Started", "This is a model of vernon's software system.", func() {
	SoftwareSystem("Software System", "Vernon's software system.", func() {
		Container("KMContainer", "KM container with real content", "docker & golang", func() {
			// can have Tag, URL, Uses, Delivers inside this
			Component("KM component", "A custom component", "Go and Python", func() {
				Struct("KM Struct", "A custom struct", "Go only", func() {

				})
				Interface("KM Interface", "A custom interface", "Go only", func() {

				})

			})
		})
	})

})
