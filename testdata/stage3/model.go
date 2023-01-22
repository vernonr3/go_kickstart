package design

//import . "goa.design/model/dsl"

import . "go_kickstart/dsl"

/*
Nesting to the level envisaged in the DSL created for the model..
Extended to include the struct, interface and method levels..
Extended further to include parameters
*/

var _ = Design("Getting Started", "This is a model of vernon's software system.", func() {
	SoftwareSystem("Software System", "Vernon's software system.", func() {
		Container("KMContainer", "KM container with real content", "docker & golang", func() {
			// can have Tag, URL, Uses, Delivers inside this
			Component("KM_component", "A custom component", "Go and Python", func() {
				Struct("KM_Struct", "A custom struct", "Go only", func() {
					Interface("KM_Method Interface", "A custom interface", "Go only", func() {
						Method("KM_struct_method", "Means of introspection", "Go only", func() {
							InputParameter("IVar1", "int64")
							ReturnParameter("bvalue", "bool")
						})
					})
					Method("km_struct_internal_method", "Means of introspection", "Go only", func() {
						InputParameter("intVar1", "int32")
						ReturnParameter("cvalue", "int")
					})

				})
				Interface("KM_Functional_Interface", "A custom interface", "Go only", func() {
					Method("KM_Function", "Means of reflection", "Go only", func() {
						InputParameter("Var1", "string")
						ReturnParameter("dvalue", "int64")
					})
				})
			})
		})
	})

})

/*
	Method("km_struct_internal_method", "Means of introspection", "Go only", func() {
		InputParameter("intVar1", "int32")
		ReturnParameter("cvalue", "int")
	})

*/
