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
				Struct("StructA", "A custom struct", "Go only", func() {
					Interface("StructA_InterfaceA", "IASA first interface", "Go only", func() {
						Method("MethodA_InterfaceA_StructA", "MAIASA ABC", "Go only", func() {
							InputParameter("IVar1", "int32")
							ReturnParameter("bvalue", "bool")
						})
						Method("MethodB_InterfaceA_StructA", "MBIASA DEF", "Go only", func() {
							InputParameter("IVar3", "int")
							ReturnParameter("fvalue", "bool")
						})
					})
					Interface("StructA_InterfaceB", "IBSA second interface", "Go only", func() {
						Method("MethodA_InterfaceB_StructA", "MAIBSA GHI", "Go only", func() {
							InputParameter("IVar2", "int64")
							ReturnParameter("bvalue", "string")
						})
					})
					Method("km_struct_internal_method", "maSA MNO", "Go only", func() {
						InputParameter("intVar1", "int32")
						ReturnParameter("cvalue", "int")
					})

				})
				Interface("FunctionInterfaceA", "FIA first functional interface", "Go only", func() {
					Method("MethodA_FunctionInterfaceA", "MAFA PQS", "Go only", func() {
						InputParameter("Var1", "string")
						ReturnParameter("dvalue", "int64")
					})
					Method("MethodB_FunctionInterfaceA", "MBFA XYZ", "Go only", func() {
						InputParameter("Var2", "string")
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
