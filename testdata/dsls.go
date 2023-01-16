package testdata

import . "goa.design/goa/v3/dsl"

// pinch design & elements from model plugin...
/*
// in model/dsl/design.go
var _ = Design("My Design", "A great architecture.", func() {
// in model/dsl/elements.go
//	    SoftwareSystem("My Software System",["description"],func(){
			URL("https://goa.design/docs/mysystem") // where more information may be found about this system
			External() // external location from this system
			Prop("name", "value") // N of these may occur
			Uses("Other System", "Uses", "gRPC", Synchronous)
	        Delivers("Customer", "Delivers emails to", "SMTP", Synchronous)
			Container("<name>", "[description]", "[technology]", func(){
				// can have Tag, URL, Uses, Delivers inside this
				Component(Container, "My component", "A component", "Go and Goa", func() {
//	                Tag("bill processing")
//	                URL("https://goa.design/mysystem")
//	                Uses("Other Component", "Uses", "gRPC", Synchronous)
//	                Delivers("Customer", "Delivers emails to", "SMTP", Synchronous)
//	            })
			})

		})
// in model/dsl/deployment.go
		DeploymentEnvironment("production", func() {
//	         DeploymentNode("AppServer", "Application server", "Go and Goa v3")
//	         InfrastructureNote("Router", "External traffic router", "AWS Route 53")
//	         ContainerInstance(Container,func(){
				HealthCheck("check", func() {
//	                Interval(10)
					Timeout(1000)
//	            })
			 })
//	     })
//	})


*/

var MAlias = func() {
	var _ = Type("Alias", String, func() {
		MinLength(10)
	})
}
