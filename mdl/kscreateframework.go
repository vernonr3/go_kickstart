package mdl

import (
	"fmt"
	"go_kickstart/expr"
	"go_kickstart/gen"
)

func CreateCodeFramework(d *expr.Design) *Design {
	model := &Model{}
	m := d.Model
	model.Systems = make([]*SoftwareSystem, len(m.Systems))
	for i, sys := range m.Systems {
		model.Systems[i] = createSystemFramework(sys)
	}
	return &Design{
		Name:        d.Name,
		Description: d.Description,
		Version:     d.Version,
		Model:       model,
	}

}

func createSystemFramework(sys *expr.SoftwareSystem) *SoftwareSystem {
	return &SoftwareSystem{
		ID:          sys.ID,
		Name:        sys.Name,
		Description: sys.Description,
		Technology:  sys.Technology,
		Tags:        sys.Tags,
		URL:         sys.URL,
		Properties:  sys.Properties,
		//Relationships: modelizeRelationships(sys.Relationships),
		Location:   LocationKind(sys.Location),
		Containers: createContainerFramework(sys.Containers),
	}
}

func createContainerFramework(cs []*expr.Container) []*Container {
	res := make([]*Container, len(cs))
	for _, c := range cs {
		createComponentFramework(c.Components)
	}
	return res
}

func createComponentFramework(components []*expr.Component) string {
	/*	var structTempls []string
		var structMethodTempls [][]string
		var methodSignatureTemplates [][]string
	*/
	var interfaceCodeStruct []gen.InterfaceCodeStruct
	//var methodTemplates []string
	//var structInterfaceTemplates [][]string
	// create the segments

	// load the template for component - the ComponentName will be the folder.
	// the filename will be arbitrary for now... as the differentiation bu filename adds complexity for little gain..
	// the expectation
	// fill the template
	// generate the code
	//res := make([]*Component, len(cs))

	//iterate over the number of components
	for _, component := range components {
		codeStruct := createStructFramework(component.Structs)
		interfaceCodeStruct = createInterfaceFunctionalFramework(component.Interfaces)
		gen.GenerateComponent(*component, codeStruct, interfaceCodeStruct)
	}

	return ""
}

func getMethodDescriptions(methods []*expr.Method) []string {
	var Descriptions []string
	Descriptions = make([]string, 0)
	for _, method := range methods {
		Descriptions = append(Descriptions, method.Description)
	}
	return Descriptions
}

func createStructFramework(structs []*expr.Struct) []gen.CodeStruct { // (structTempls []string, structMethodTempls [][]string, methodSignatureTempls [][]string, interfaceTempls [][]string) {
	codestructs := make([]gen.CodeStruct, 0)
	singlecodestruct := gen.CodeStruct{}
	//methodTempls := make([][]string, len(structs))
	//structTempls = make([]string, len(structs))
	//structMethodTempls = make([][]string, len(structs))
	//methodSignatureTempls = make([][]string, len(structs))
	//interfaceTempl := make([]gen.InterfaceCodeStruct, len(structs))
	// iterate over the number of structs
	for _, mstruct := range structs {
		//for each struct
		interfaceTempl := createInterfaceFramework(mstruct, mstruct.Interfaces)
		methodSignatureTempl, _ := createMethodSignatures(mstruct.Methods)
		methodTempl := createMethodFramework(mstruct.Methods)
		structTempl := gen.GenerateStruct(*mstruct, methodSignatureTempl)
		structMethodTempl := gen.GenerateStructMethods(*mstruct, getMethodDescriptions(mstruct.Methods), methodTempl)
		// aggregate across all structs
		singlecodestruct.StructTemplate = structTempl
		//singlecodestruct.MockStructTemplate = mockTempl
		//singlecodestruct.ExternalMethods = methodSignatureTempl
		singlecodestruct.InternalMethodSignatures = methodSignatureTempl
		singlecodestruct.InternalMethods = methodTempl
		singlecodestruct.StructMethods = structMethodTempl
		singlecodestruct.InterfaceCodeStructs = interfaceTempl
		//singlecodestruct.MockMethods = mockMethodTempl
		codestructs = append(codestructs, singlecodestruct)
	}

	return codestructs
}

func createInterfaceFunctionalFramework(compInterface []*expr.Interface) []gen.InterfaceCodeStruct {
	interfaceMethodTempls := make([]gen.InterfaceCodeStruct, 0)
	// iterate over the number of interfaces in this component
	for _, minterface := range compInterface {
		interfaceMethodTempl := gen.InterfaceCodeStruct{}
		methodTempl := createMethodFramework(minterface.Methods)
		interfaceMethodTempl.ExternalFunctions = methodTempl
		interfaceMethodTempls = append(interfaceMethodTempls, interfaceMethodTempl)
	}
	return interfaceMethodTempls
}

func createInterfaceFramework(mstruct *expr.Struct, interfaces []*expr.Interface) []gen.InterfaceCodeStruct {
	interfaceMethodTempls := make([]gen.InterfaceCodeStruct, 0)
	// iterate over the number of interfaces in this component
	for _, minterface := range interfaces {
		interfaceMethodTempl := gen.InterfaceCodeStruct{}
		_, methodRawSignatureTempl := createMethodSignatures(minterface.Methods) // i.e. without the body
		methodTempl := createMethodFramework(minterface.Methods)
		structMethodTempl := gen.GenerateStructMethods(*mstruct, getMethodDescriptions(minterface.Methods), methodTempl)
		//methodTempls = append(methodTempls, methodTempl)
		interfaceMethodTempl.StructTemplate = mstruct.Name
		interfaceMethodTempl.ExternalMethods = structMethodTempl
		fmt.Printf("%s has len %d external methods\n", interfaceMethodTempl.StructTemplate, len(interfaceMethodTempl.ExternalMethods))
		interfaceMethodTempl.InterfaceSignature = gen.GenerateInterface(*minterface, methodRawSignatureTempl)
		interfaceMethodTempls = append(interfaceMethodTempls, interfaceMethodTempl)
	}
	return interfaceMethodTempls
}

func createMethodSignatures(methods []*expr.Method) ([]string, []string) {
	methodTemplSignatures := make([]string, 0)
	methodRawTemplSignatures := make([]string, 0)
	// iterate over the number of methods in this struct or interface
	for _, method := range methods {
		inputs := createInputParameterFramework(method.InputParameters)
		returns := createReturnParameterFramework(method.ReturnParameters)
		methodTemplSignatures = append(methodTemplSignatures, gen.GenerateMethodSignature(*method, inputs, returns))
		methodRawTemplSignatures = append(methodRawTemplSignatures, gen.GenerateRawMethodSignature(*method, inputs, returns))
	}
	return methodTemplSignatures, methodRawTemplSignatures
}

func createMethodFramework(methods []*expr.Method) []string {
	methodTempls := make([]string, 0)
	// iterate over the number of methods in this struct or interface
	for _, method := range methods {
		inputs := createInputParameterFramework(method.InputParameters)
		returns := createReturnParameterFramework(method.ReturnParameters)
		methodTempls = append(methodTempls, gen.GenerateMethod(*method, inputs, returns))
	}
	return methodTempls
}

func createInputParameterFramework(inputparameters []*expr.InputParameter) string {
	var parameterlist string
	parameterlist = gen.GenerateInputParameter(inputparameters)
	return parameterlist
}

func createReturnParameterFramework(returnparameters []*expr.ReturnParameter) string {
	var parameterlist string
	parameterlist = gen.GenerateReturnParameter(returnparameters)
	return parameterlist
}
