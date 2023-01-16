package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"goa.design/goa/v3/codegen"
	"golang.org/x/tools/go/packages"
)

/*
func gen(pkg string, debug bool) ([]byte, error) {
	var bytes []byte
	fmt.Printf("Generate code\n")
	return bytes, nil
}*/

const tmpDirPrefix = "mdl--"

func gen(pkg string, dir string, debug bool, outputfile string) ([]byte, error) {
	// Validate package import path
	if _, err := packages.Load(&packages.Config{Mode: packages.NeedName, Dir: dir}, pkg); err != nil {
		return nil, err
	}

	// Write program that generates JSON
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	tmpDir, err := os.MkdirTemp(cwd, tmpDirPrefix)
	if err != nil {
		return nil, err
	}
	defer func() { os.RemoveAll(tmpDir) }()
	var sections []*codegen.SectionTemplate
	{
		imports := []*codegen.ImportSpec{
			codegen.SimpleImport("fmt"),
			codegen.SimpleImport("encoding/json"),
			codegen.SimpleImport("os"),
			//codegen.SimpleImport("goa.design/model/mdl"),
			//codegen.SimpleImport("goa.design/model/mdl"),
			codegen.SimpleImport("go_kickstart/dsl"),
			codegen.SimpleImport("go_kickstart/mdl"),
			codegen.NewImport("_", pkg),
		}
		sections = []*codegen.SectionTemplate{
			codegen.Header("Code Generator", "main", imports),
			{Name: "main", Source: mainT},
		}
	}
	cf := &codegen.File{Path: "main.go", SectionTemplates: sections}
	fmt.Printf("Render code for package %s\n", pkg)
	if _, err := cf.Render(tmpDir); err != nil {
		return nil, err
	}
	// Compile program
	fmt.Printf("Find go program\n")
	gobin, err := exec.LookPath("go")
	if err != nil {
		return nil, fmt.Errorf(`failed to find a go compiler, looked in "%s"`, os.Getenv("PATH"))
	}
	fmt.Printf("Compile the program\n")
	//if _, err := runCmd(gobin, tmpDir, "build", "-o", "mdl", "-pkgdir", tmpDir); err != nil {
	if _, err := runCmd(gobin, tmpDir, "build", "-o", "mdl"); err != nil {
		fmt.Printf("Build error : %s\n", err.Error())
		return nil, err
	}

	// Run program
	fmt.Printf("Run the program\n")
	o, err := runCmd(path.Join(tmpDir, "mdl"), tmpDir, outputfile)
	if debug {
		fmt.Fprintln(os.Stderr, o)
	}
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path.Join(tmpDir, outputfile))

}

func runCmd(path, dir string, args ...string) (string, error) {
	_ = os.Setenv("GO111MODULE", "on")
	args = append([]string{path}, args...) // args[0] becomes exec path
	fmt.Printf("runCmd %s args %s Dir %s\n", path, args, dir)
	c := exec.Cmd{Path: path, Args: args, Dir: dir}
	fmt.Printf("Get the combined output\n")
	b, err := c.CombinedOutput()
	if err != nil {
		if len(b) > 0 {
			return "", fmt.Errorf(string(b))
		}
		return "", fmt.Errorf("failed to run command %q in directory %q: %s", path, dir, err)
	}
	return string(b), nil
}

// mainT is the template for the generator main.
const mainT = `func main() {
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
`
