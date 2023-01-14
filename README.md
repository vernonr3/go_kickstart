# go_kickstart
This is to allow code for a Go component/service to be "kick started". 
An author creates a DSL (Domain Specific Language) in Go to describe the component.
It uses goadesign/goa to interpret a DSL (Domain Specific Language) and generate the shell implementation. 
Although in goa terminology it is a plugin(slave); it is the master and uses the goa code as a library.
