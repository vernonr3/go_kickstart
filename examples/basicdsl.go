package design

import (
	. "go_kickstart/dsl"

	. "goa.design/model/dsl"
)

var _ = Design("Model Usage", "Not a software architecture but a diagram illustrating how to use Model.", func() {
	SoftwareSystem("Model Usage", func() {
		Container("Model Usage", func() {
			Component("Design", "Go package containing Model DSL that describes the system architecture", func() {
				Uses("Visual Editor", "mdl serve")
				Uses("Design JSON", "mdl gen")
				// Uses("Static HTML", "mdl gen")
				Uses("Structurizr workspace JSON", "stz gen")
				Tag("Design")
				Package("MyPackage", func() {
					Tag("Pack")
				})
			})
		})
	})
})

var _ = Package("My Package", func() {

})
