package main

import (
	"os"
	"testing"
)

func Test_BasicMode(t *testing.T) {
	out := "model.json"
	//b, err := gen("goa.design/model/examples/basic/model", "", true)
	//b, err := gen("github.com/vernonr3/go_kickstart/testdata/basic", "", true)
	b, err := gen("go_kickstart/testdata/basic", "", true, out)
	if err == nil {
		err = os.WriteFile(out, b, 0644)
	}
}

func Test_Stage1(t *testing.T) {
	out := "model1.json"
	b, err := gen("go_kickstart/testdata/stage1", "", true, out)
	if err == nil {
		err = os.WriteFile(out, b, 0644)
	}
}

func Test_Stage2(t *testing.T) {
	out := "model2.json"
	b, err := gen("go_kickstart/testdata/stage2", "", true, out)
	if err == nil {
		err = os.WriteFile(out, b, 0644)
	}
}
