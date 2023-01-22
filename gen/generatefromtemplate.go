package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"go_kickstart/expr"
	"log"
	"os"
	"regexp"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
)

type InterfaceCodeStruct struct {
	StructTemplate     string
	MockStructTemplate string
	ExternalMethods    string   // these will be represented within the appropriate interface type
	ExternalFunctions  []string // these will be represented within the appropriate interface type
	InterfaceSignature string   // these will be represented within the appropriate interface type
}

type CodeStruct struct {
	StructTemplate           string
	MockStructTemplate       string
	InternalMethodSignatures []string              // these won't be within an interface - but will have access to the struct
	InternalMethods          []string              // the methods within component that don't reference the struct
	StructMethods            string                // the methods for the struct that reference the struct
	MockMethods              []string              // the methods for the mock struct that realise the interface
	InterfaceCodeStructs     []InterfaceCodeStruct // methods reached through the interface
}

type MockCodeStruct struct {
	MockStructTemplate       string
	InternalMethodSignatures []string              // these won't be within an interface - but will have access to the struct
	InternalMethods          []string              // the methods within component that don't reference the struct
	StructMethods            string                // the methods for the struct that reference the struct
	MockMethods              []string              // the methods for the mock struct that realise the interface
	InterfaceCodeStructs     []InterfaceCodeStruct // methods reached through the interface
}

type MockStruct struct {
	MockStructTemplate string
	MockMethods        string                // the methods for the mock struct that realise the interface
	MockInterfaceCodes []InterfaceCodeStruct // methods reached through the interface
}

type StructData struct {
	StructName string
}

type StructMethodData struct {
	MethodTempls       []string
	StructName         string
	MethodDescriptions []string
}

type InterfaceData struct {
	InterfaceName    string
	Methods          []string
	InterfaceDescrip string
}

type MethodData struct {
	StructName        string
	MethodName        string
	MethodDescription string
	Inputs            string
	InputParams       string
	ReturnParams      string
}

type ParamData struct {
	Params []Param
}
type Param struct {
	VarName string
	VarType string
}

type ComponentData struct {
	ComponentName         string
	Codestructs           []CodeStruct
	InterfaceCodeStructs  []InterfaceCodeStruct
	StructInterfaceTempls [][]string
	FunctionTemplates     []string
}

type MockComponentData struct {
	ComponentName         string
	MockCodeStructs       []MockCodeStruct
	InterfaceCodeStructs  []InterfaceCodeStruct
	StructInterfaceTempls [][]string
}

func GenerateComponent(component expr.Component, codestructs []CodeStruct, interfaceCodeStructs []InterfaceCodeStruct, functionTemplates []string) {
	//there should be a set per struct
	mcodestruct := ComponentData{
		ComponentName:        component.Name,
		Codestructs:          codestructs,
		InterfaceCodeStructs: interfaceCodeStructs,
		FunctionTemplates:    functionTemplates,
		//StructInterfaceTempls: structInterfaceTempls,
	}
	fmt.Printf("generating component\n")

	componentString := processTemplate("component.tmpl", "codetemplates/component.tmpl", "tmp/component.txt", mcodestruct)
	fmt.Printf("%s\n", componentString)
}

func GenerateMockComponent(component expr.Component, mockcodestructs []MockCodeStruct) {
	//there should be a set per struct
	mcodestruct := MockComponentData{
		ComponentName:   component.Name,
		MockCodeStructs: mockcodestructs,
	}
	fmt.Printf("generating mock component\n")

	componentString := processTemplate("mockcomponent.tmpl", "codetemplates/mockcomponent.tmpl", "tmp/mockcomponent.txt", mcodestruct)
	fmt.Printf("%s\n", componentString)
}

func GenerateStruct(mstruct expr.Struct, methodTempls []string) string {
	//fmt.Printf("generating struct method signatures\n")
	//fmt.Printf("name %s\n", mstruct.Name)
	data := &StructData{
		StructName: mstruct.Name,
		//Methods:    methodTempls,
	}
	structstring := processTemplate("struct.tmpl", "codetemplates/struct.tmpl", "tmp/struct.txt", data)
	//fmt.Printf("Struct : \n%s\n", structstring)
	return structstring
}

func GenerateMockStruct(mstruct expr.Struct, methodTempls []string) string {
	//fmt.Printf("generating struct method signatures\n")
	//fmt.Printf("name %s\n", mstruct.Name)
	data := &StructMethodData{
		StructName:   "Mock" + mstruct.Name,
		MethodTempls: methodTempls,
	}
	structstring := processTemplate("mockstruct.tmpl", "codetemplates/mockstruct.tmpl", "tmp/mockstruct.txt", data)
	//fmt.Printf("Struct : \n%s\n", structstring)
	return structstring
}

// generate methods that implement and interface for a struct
func GenerateStructMethods(mstruct expr.Struct, methodDescriptions []string, methodTempls []string) string {
	var methodstrings = make([]string, 1)
	if mstruct.Name == "" {
		mstruct.Name = "blank"
	}
	data := &StructMethodData{
		MethodTempls:       methodTempls,
		StructName:         mstruct.Name,
		MethodDescriptions: methodDescriptions,
	}
	fmt.Printf("Generate StructMethod Evaluating for %s\n", mstruct.Name)
	methodstring := processTemplate("structmethod.tmpl", "codetemplates/structmethod.tmpl", "tmp/structmethod.txt", data)
	fmt.Printf("struct method : \n%s\n", methodstring)
	methodstrings[0] = methodstring
	return methodstring
}

// generate methods that implement and interface for a mock struct
func GenerateMockStructMethods(mstruct expr.Struct, methodDescriptions []string, methodTempls []string) string {
	var methodstrings = make([]string, 1)
	if mstruct.Name == "" {
		mstruct.Name = "blank"
	}
	data := &StructMethodData{
		MethodTempls:       methodTempls,
		StructName:         "mock" + mstruct.Name,
		MethodDescriptions: methodDescriptions,
	}
	fmt.Printf("Generate StructMethod Evaluating for %s\n", mstruct.Name)
	methodstring := processTemplate("mockstructmethod.tmpl", "codetemplates/mockstructmethod.tmpl", "tmp/mockstructmethod.txt", data)
	fmt.Printf("struct method : \n%s\n", methodstring)
	methodstrings[0] = methodstring
	return methodstring
}

// generates interface with method signatures
func GenerateInterface(miface expr.Interface, methodTempls []string) string {
	data := &InterfaceData{
		InterfaceName:    miface.Name,
		Methods:          methodTempls,
		InterfaceDescrip: miface.Description,
	}
	interfacestring := processTemplate("interface.tmpl", "codetemplates/interface.tmpl", "tmp/interface.txt", data)
	fmt.Printf("Interface : \n%s\n", interfacestring)
	return interfacestring
}

func GenerateMockInterface(miface expr.Interface, methodTempls []string) string {
	data := &InterfaceData{
		InterfaceName:    miface.Name,
		Methods:          methodTempls,
		InterfaceDescrip: miface.Description,
	}
	interfacestring := processTemplate("mockinterface.tmpl", "codetemplates/mockinterface.tmpl", "tmp/mockinterface.txt", data)
	fmt.Printf("Interface : \n%s\n", interfacestring)
	return interfacestring
}

// generates a method body with a trivial welcome message
func GenerateMethod(method expr.Method, inputparams string, returnparams string) string {
	//fmt.Printf("generating method %s with inputparams %s and returnparams %s\n", method.Name, inputparams, returnparams)
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodstring := processTemplate("method.tmpl", "codetemplates/method.tmpl", "tmp/method.txt", data)
	fmt.Printf("method : \n%s\n", methodstring)
	return methodstring
}

// generates a method body with a trivial welcome message
func GenerateFunction(method expr.Function, inputparams string, returnparams string) string {
	//fmt.Printf("generating method %s with inputparams %s and returnparams %s\n", method.Name, inputparams, returnparams)
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	functionstring := processTemplate("function.tmpl", "codetemplates/function.tmpl", "tmp/function.txt", data)
	fmt.Printf("function : \n%s\n", functionstring)
	return functionstring
}

func GenerateMethod4Struct(method expr.Method, inputparams string, returnparams string) string {
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodstring := processTemplate("method4struct.tmpl", "codetemplates/method4struct.tmpl", "tmp/method4struct.txt", data)
	fmt.Printf("method : \n%s\n", methodstring)
	return methodstring
}

func GenerateMockMethod4Struct(method expr.Method, inputs string, inputparams string, returnparams string) string {
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		Inputs:            inputs,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodstring := processTemplate("mockmethod4struct.tmpl", "codetemplates/mockmethod4struct.tmpl", "tmp/mockmethod4struct.txt", data)
	fmt.Printf("mock method : \n%s\n", methodstring)
	return methodstring
}

func GenerateSpoofMethod(mstructName string, method expr.Method, inputparams string, returnparams string) string {
	data := &MethodData{
		StructName:        mstructName,
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodstring := processTemplate("spoofmethod.tmpl", "codetemplates/spoofmethod.tmpl", "tmp/spoofmethod.txt", data)
	fmt.Printf("spoof method : \n%s\n", methodstring)
	return methodstring
}

func GenerateMockMethodSignature(method expr.Method, inputparams string, returnparams string) string {
	//fmt.Printf("generating method signature %s with inputparams %s and returnparams %s\n", method.Name, inputparams, returnparams)
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodsignature := processTemplate("mockmethodsignature.tmpl", "codetemplates/mockmethodsignature.tmpl", "tmp/mockmethodsignature.txt", data)
	fmt.Printf("methodsignature : \n%s\n", methodsignature)
	return methodsignature
}

// generates a function/method signature of the type that can be included in an interface definition
func GenerateMethodSignature(method expr.Method, inputparams string, returnparams string) string {
	//fmt.Printf("generating method signature %s with inputparams %s and returnparams %s\n", method.Name, inputparams, returnparams)
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodsignature := processTemplate("methodsignature.tmpl", "codetemplates/methodsignature.tmpl", "tmp/methodsignature.txt", data)
	fmt.Printf("methodsignature : \n%s\n", methodsignature)
	return methodsignature
}

// generates a function/method signature of the type that can be included as a method prefixed by a struct/struct pointer
func GenerateRawMethodSignature(method expr.Method, inputparams string, returnparams string) string {
	//fmt.Printf("generating method signature %s with inputparams %s and returnparams %s\n", method.Name, inputparams, returnparams)
	data := &MethodData{
		MethodName:        method.Name,
		MethodDescription: method.Description,
		InputParams:       inputparams,
		ReturnParams:      returnparams,
	}
	methodsignature := processTemplate("rawmethodsignature.tmpl", "codetemplates/rawmethodsignature.tmpl", "tmp/rawmethodsignature.txt", data)
	fmt.Printf("raw methodsignature : \n%s\n", methodsignature)
	return methodsignature
}
func GenerateInputs(iparams []*expr.InputParameter) string {
	var params []Param
	//fmt.Printf("generating input parameters\n")
	for _, param := range iparams {
		params = append(params, Param{
			VarName: param.Name,
		})

	}
	data := &ParamData{Params: params}
	inputParams := processTemplate("inputs.tmpl", "codetemplates/inputs.tmpl", "tmp/inputs.txt", data)

	return inputParams
}

// generate input parameters
func GenerateInputParameter(iparams []*expr.InputParameter) string {
	var params []Param
	//fmt.Printf("generating input parameters\n")
	for _, param := range iparams {
		//fmt.Printf("Input param : Name %s Type %s\n", param.Name, param.Description)
		params = append(params, Param{
			VarName: param.Name,
			VarType: param.Description,
		})

	}
	data := &ParamData{Params: params}
	inputParams := processTemplate("inputparameters.tmpl", "codetemplates/inputparameters.tmpl", "tmp/inputparameters.txt", data)

	return inputParams
}

// generate return parameters
func GenerateReturnParameter(rparams []*expr.ReturnParameter) string {
	var params []Param
	//fmt.Printf("generating return parameters\n")
	for _, param := range rparams {
		//fmt.Printf("Return Param Name %s Type %s\n", param.Name, param.Description)
		params = append(params, Param{
			VarName: param.Name,
			VarType: param.Description,
		})
	}
	data := &ParamData{Params: params}
	returnParams := processTemplate("returnparameters.tmpl", "codetemplates/returnparameters.tmpl", "tmp/returnparameters.txt", data)
	return returnParams
}

func processTemplate(templateName string, fileName string, outputFile string, data any) string {
	//mtemplate, err :=
	rootdir, err := findroot()
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(rootdir + fileName))
	var processed bytes.Buffer
	err = tmpl.ExecuteTemplate(&processed, templateName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	if len(outputFile) > 0 {
		fmt.Println("Writing file: ", rootdir+outputFile)
		f, _ := os.Create(rootdir + outputFile)
		w := bufio.NewWriter(f)
		w.WriteString(processed.String())
		w.Flush()

	}
	return processed.String()
}

func findroot() (string, error) {
	var relativedirs string = ""
	patt := regexp.MustCompile(`(.*)go_kickstart\/([^\/]*)(\/*.*)`)
	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	results := patt.FindAllStringSubmatch(workdir, -1)
	fmt.Printf("%q\n", results)
	if len(results) > 0 {
		if len(results[0][3]) > 0 {
			relativedirs = "../../"
		} else {
			relativedirs = "../"
		}

	}
	return relativedirs, nil
}
