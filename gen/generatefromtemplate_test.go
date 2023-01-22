package gen

import (
	"go_kickstart/expr"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindRoot(t *testing.T) {
	findroot()
}

func Test_GenerateStructMethod(t *testing.T) {
	var methodTempls []string
	var methodDescriptions []string
	e := &expr.Element{
		Name: "FooBar",
	}
	mstruct := expr.Struct{
		Element: e,
	}
	methodDescriptions = make([]string, 2)
	methodDescriptions[0] = "Spoof allows us to do something useless"
	methodDescriptions[1] = "DoeMethod allows us to do something altogether more worthwhile"
	method, inputparams, returnparams := makemethod("SpoofMethod")
	methodtempl1 := GenerateMethod(method, inputparams, returnparams)
	method, inputparams, returnparams = makemethod("DoeMethod")
	methodTempl2 := GenerateMethod(method, inputparams, returnparams)
	methodTempls = append(methodTempls, methodtempl1, methodTempl2)
	result := GenerateStructMethods(mstruct, methodDescriptions, methodTempls)
	assert.Equal(t, "type FooInterface interface {func SpoofMethod(var1 string,var2 bool)(string,bool)\nfunc DoeMethod(var1 string,var2 bool)(string,bool)\n}", result)
}

func Test_GenerateInterface(t *testing.T) {
	//method expr.Method,  string, returnparams string
	var methodTempls []string
	e := &expr.Element{
		Name: "FooInterface",
	}
	miface := expr.Interface{
		Element: e,
	}
	method, inputparams, returnparams := makemethod("SpoofMethod")
	methodtempl1 := GenerateMethodSignature(method, inputparams, returnparams)
	method, inputparams, returnparams = makemethod("DoeMethod")
	methodTempl2 := GenerateMethodSignature(method, inputparams, returnparams)
	methodTempls = append(methodTempls, methodtempl1, methodTempl2)
	result := GenerateInterface(miface, methodTempls)
	assert.Equal(t, "type FooInterface interface {func SpoofMethod(var1 string,var2 bool)(string,bool)\nfunc DoeMethod(var1 string,var2 bool)(string,bool)\n}", result)
}

func makemethod(name string) (expr.Method, string, string) {
	inputparams := GenerateInputParameter(makeinputparams())
	returnparams := GenerateReturnParameter(makereturnparams())
	//method expr.Method,  string, returnparams string
	e := &expr.Element{
		Name: name,
	}
	method := expr.Method{
		Element: e,
	}
	return method, inputparams, returnparams
}

func Test_GenerateMethod(t *testing.T) {
	method, inputparams, returnparams := makemethod("SpoofMethod")
	result := GenerateMethod(method, inputparams, returnparams)
	assert.Equal(t, "func SpoofMethod(var1 string,var2 bool)(string,bool){\n    fmt.Printf(\"Welcome to SpoofMethod\\n\")\n}", result)
}

func Test_GenerateMethodSignature(t *testing.T) {
	inputparams := GenerateInputParameter(makeinputparams())
	returnparams := GenerateReturnParameter(makereturnparams())
	//method expr.Method,  string, returnparams string
	e := &expr.Element{
		Name: "SpoofMethod",
	}
	method := expr.Method{
		Element: e,
	}
	result := GenerateMethodSignature(method, inputparams, returnparams)
	assert.Equal(t, "func SpoofMethod(var1 string,var2 bool)(string,bool)", result)
}

func Test_GenerateRawMethodSignature(t *testing.T) {
	inputparams := GenerateInputParameter(makeinputparams())
	returnparams := GenerateReturnParameter(makereturnparams())
	//method expr.Method,  string, returnparams string
	e := &expr.Element{
		Name: "SpoofMethod",
	}
	method := expr.Method{
		Element: e,
	}
	result := GenerateRawMethodSignature(method, inputparams, returnparams)
	assert.Equal(t, "SpoofMethod(var1 string,var2 bool)(string,bool)", result)
}

func makeinputparams() []*expr.InputParameter {
	var iparams []*expr.InputParameter
	iparams = make([]*expr.InputParameter, 2)
	element := &expr.Element{
		Name:        "var1",
		Description: "string",
	}
	element2 := &expr.Element{
		Name:        "var2",
		Description: "bool",
	}
	iparams[0] = &expr.InputParameter{
		Element: element,
	}
	iparams[1] = &expr.InputParameter{
		Element: element2,
	}
	return iparams
}

func Test_GenerateInputParameter(t *testing.T) {
	iparams := makeinputparams()
	result := GenerateInputParameter(iparams)
	assert.Equal(t, "(var1 string,var2 bool)", result)
}

func makereturnparams() []*expr.ReturnParameter {
	var iparams []*expr.ReturnParameter
	iparams = make([]*expr.ReturnParameter, 2)
	element := &expr.Element{
		Name:        "var1",
		Description: "string",
	}
	element2 := &expr.Element{
		Name:        "var2",
		Description: "bool",
	}
	iparams[0] = &expr.ReturnParameter{
		Element: element,
	}
	iparams[1] = &expr.ReturnParameter{
		Element: element2,
	}
	return iparams
}

func Test_GenerateReturnParameter(t *testing.T) {
	iparams := makereturnparams()
	result := GenerateReturnParameter(iparams)
	assert.Equal(t, "(string,bool)", result)
}
