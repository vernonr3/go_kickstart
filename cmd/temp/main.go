// Code generated by goa v3.10.2, DO NOT EDIT.
//
// Code Generator
//
// Command:
// goa

package main

import (
	"encoding/json"
	"fmt"
	//"goa.design/model/mdl"
	"go_kickstart/mdl"
	//"model/mdl"
	_ "go_kickstart/testdata/stage3"
	"os"
)

func main() {
	// Retrieve output path
	out := os.Args[1]

	// Run the model DSL
	w, err := mdl.RunDSL()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	b, err := json.MarshalIndent(w, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode into JSON: %s", err.Error())
		os.Exit(1)
	}
	if err := os.WriteFile(out, b, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %s", err.Error())
		os.Exit(1)
	}
}
