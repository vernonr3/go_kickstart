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
	var interfaceCodeStruct []gen.InterfaceCodeStruct
	//iterate over the number of components
	for _, component := range components {
		codeStruct := createStructFramework(component.Structs)
		interfaceCodeStruct = createInterfaceFunctionalFramework(component.Interfaces)
		functionStruct := createFuncFramework(component.Functions)
		gen.GenerateComponent(*component, codeStruct, interfaceCodeStruct, functionStruct)
		mockcodeStruct := createMockStructFramework(component.Structs)
		gen.GenerateMockComponent(*component, mockcodeStruct)
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

func createStructFramework(structs []*expr.Struct) []gen.CodeStruct {
	codestructs := make([]gen.CodeStruct, 0)
	singlecodestruct := gen.CodeStruct{}
	// iterate over the number of structs
	for _, mstruct := range structs {
		//for each struct
		interfaceTempl := createInterfaceFramework(mstruct, mstruct.Interfaces)
		methodSignatureTempl, _ := createMethodSignatures(mstruct.Methods)
		methodTempl := createStructMethodFramework(mstruct.Methods)
		internalmethodTempl := createMethodFramework(mstruct.Methods)
		structTempl := gen.GenerateStruct(*mstruct, methodSignatureTempl)
		structMethodTempl := gen.GenerateStructMethods(*mstruct, getMethodDescriptions(mstruct.Methods), methodTempl)
		// aggregate across all structs
		singlecodestruct.StructTemplate = structTempl
		//singlecodestruct.MockStructTemplate = mockTempl
		//singlecodestruct.ExternalMethods = methodSignatureTempl
		singlecodestruct.InternalMethodSignatures = methodSignatureTempl
		singlecodestruct.InternalMethods = internalmethodTempl
		singlecodestruct.StructMethods = structMethodTempl
		singlecodestruct.InterfaceCodeStructs = interfaceTempl
		//singlecodestruct.MockMethods = mockMethodTempl
		codestructs = append(codestructs, singlecodestruct)
	}

	return codestructs
}

func createMockStructFramework(structs []*expr.Struct) []gen.MockCodeStruct {
	codestructs := make([]gen.MockCodeStruct, 0)
	singlecodestruct := gen.MockCodeStruct{}
	// iterate over the number of structs
	for _, mstruct := range structs {
		// copy the structures so we don't change the original expr.Struct
		interfaceTempl := createMockInterfaceFramework(mstruct, mstruct.Interfaces)
		methodSignatureTempl, _ := createMockMethodSignatures(mstruct.Interfaces) // mstruct.Methods)
		spoofmethodTempl := createSpoofMockMethods(mstruct, mstruct.Interfaces)
		//		internalmethodTempl := createMethodFramework(mstruct.Methods)
		structTempl := gen.GenerateMockStruct(*mstruct, methodSignatureTempl)
		//		structMethodTempl := gen.GenerateStructMethods(*mstruct, getMethodDescriptions(mstruct.Methods), methodTempl)
		// 		aggregate across all structs
		singlecodestruct.MockStructTemplate = structTempl
		singlecodestruct.MockMethods = spoofmethodTempl
		//		singlecodestruct.InternalMethodSignatures = methodSignatureTempl
		//		singlecodestruct.InternalMethods = internalmethodTempl
		//		singlecodestruct.StructMethods = structMethodTempl
		singlecodestruct.InterfaceCodeStructs = interfaceTempl
		//		singlecodestruct.MockMethods = mockMethodTempl
		codestructs = append(codestructs, singlecodestruct)
	}

	return codestructs
}

func createFuncFramework(functions []*expr.Function) []string {
	funcTempls := make([]string, 0)
	// iterate over the number of methods in this struct or interface
	for _, function := range functions {
		inputs := createInputParameterFramework(function.InputParameters)
		returns := createReturnParameterFramework(function.ReturnParameters)
		funcTempls = append(funcTempls, gen.GenerateFunction(*function, inputs, returns))
	}
	return funcTempls

}

func createInterfaceFunctionalFramework(compInterface []*expr.Interface) []gen.InterfaceCodeStruct {
	interfaceMethodTempls := make([]gen.InterfaceCodeStruct, 0)
	// iterate over the number of interfaces in this component
	for _, minterface := range compInterface {
		interfaceMethodTempl := gen.InterfaceCodeStruct{}
		methodTempl := createMethodFramework(minterface.Methods)
		_, methodRawSignatureTempl := createMethodSignatures(minterface.Methods) // i.e. without the body
		interfaceMethodTempl.InterfaceSignature = gen.GenerateInterface(*minterface, methodRawSignatureTempl)
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
		methodTempl := createStructMethodFramework(minterface.Methods)
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

func createMockInterfaceFramework(mstruct *expr.Struct, interfaces []*expr.Interface) []gen.InterfaceCodeStruct {
	interfaceMethodTempls := make([]gen.InterfaceCodeStruct, 0)
	// iterate over the number of interfaces in this component
	for _, minterface := range interfaces {
		interfaceMethodTempl := gen.InterfaceCodeStruct{}
		_, methodRawSignatureTempl := createMethodSignatures(minterface.Methods) // i.e. without the body
		methodTempl := createMockStructMethodFramework(minterface.Methods)
		structMethodTempl := gen.GenerateMockStructMethods(*mstruct, getMethodDescriptions(minterface.Methods), methodTempl)
		//methodTempls = append(methodTempls, methodTempl)
		interfaceMethodTempl.StructTemplate = mstruct.Name
		interfaceMethodTempl.ExternalMethods = structMethodTempl
		fmt.Printf("%s has len %d external methods\n", interfaceMethodTempl.StructTemplate, len(interfaceMethodTempl.ExternalMethods))
		interfaceMethodTempl.InterfaceSignature = gen.GenerateMockInterface(*minterface, methodRawSignatureTempl)
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

func createMockMethodSignatures(interfaces []*expr.Interface) ([]string, []string) {
	methodTemplSignatures := make([]string, 0)
	methodRawTemplSignatures := make([]string, 0)
	// iterate over the number of methods in this struct or interface
	for _, minterface := range interfaces {
		for _, method := range minterface.Methods {
			inputs := createInputParameterFramework(method.InputParameters)
			returns := createReturnParameterFramework(method.ReturnParameters)
			methodTemplSignatures = append(methodTemplSignatures, gen.GenerateMockMethodSignature(*method, inputs, returns))
			methodRawTemplSignatures = append(methodRawTemplSignatures, gen.GenerateRawMethodSignature(*method, inputs, returns))
		}
	}
	return methodTemplSignatures, methodRawTemplSignatures
}

func createSpoofMockMethods(mstruct *expr.Struct, interfaces []*expr.Interface) []string {
	spoofmethods := make([]string, 0)
	for _, minterface := range interfaces {
		for _, method := range minterface.Methods {
			inputs := createInputParameterFramework(method.InputParameters)
			returns := createReturnParameterFramework(method.ReturnParameters)
			spoofmethods = append(spoofmethods, gen.GenerateSpoofMethod(mstruct.Name, *method, inputs, returns))
		}
	}
	return spoofmethods
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

func createStructMethodFramework(methods []*expr.Method) []string {
	methodTempls := make([]string, 0)
	// iterate over the number of methods in this struct or interface
	for _, method := range methods {
		inputs := createInputParameterFramework(method.InputParameters)
		returns := createReturnParameterFramework(method.ReturnParameters)
		methodTempls = append(methodTempls, gen.GenerateMethod4Struct(*method, inputs, returns))
	}
	return methodTempls
}

func createMockStructMethodFramework(methods []*expr.Method) []string {
	methodTempls := make([]string, 0)
	// iterate over the number of methods in this struct or interface
	for _, method := range methods {
		inputs := createInputFramework(method.InputParameters)
		inputparams := createInputParameterFramework(method.InputParameters)
		returnparams := createReturnParameterFramework(method.ReturnParameters)
		methodTempls = append(methodTempls, gen.GenerateMockMethod4Struct(*method, inputs, inputparams, returnparams))
	}
	return methodTempls
}

func createInputParameterFramework(inputparameters []*expr.InputParameter) string {
	var parameterlist string
	parameterlist = gen.GenerateInputParameter(inputparameters)
	return parameterlist
}

func createInputFramework(inputparameters []*expr.InputParameter) string {
	var parameterlist string
	parameterlist = gen.GenerateInputs(inputparameters)
	return parameterlist
}

func createReturnParameterFramework(returnparameters []*expr.ReturnParameter) string {
	var parameterlist string
	parameterlist = gen.GenerateReturnParameter(returnparameters)
	return parameterlist
}
