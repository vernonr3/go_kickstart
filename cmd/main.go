package main

//go:generate esc -o webapp.go -pkg main -prefix webapp/dist webapp/dist/

import (
	//"encoding/json"
	"flag"
	"fmt"
	Pkg "go_kickstart/pkg"
	"os"
	"path/filepath"
	"strings"

	"goa.design/goa/v3/codegen"
	//"goa.design/model/mdl"
	//model "goa.design/model/pkg"
)

func processArgs() (cmd string, pkg string, debug *bool, out *string, pdir string, showUsage func()) {
	var (
		gset       = flag.NewFlagSet("global", flag.ExitOnError)
		help, h    *bool
		genset     = flag.NewFlagSet("gen", flag.ExitOnError)
		genfullset = flag.NewFlagSet("genfull", flag.ExitOnError)
	)
	dir := genset.String("dir", codegen.Gendir, "set output directory used by editor to save SVG files")
	pdir = *dir
	out = genset.String("out", "design.json", "set path to generated JSON representation")
	showUsage = func() { printUsage(genfullset, genset, gset) }
	addGlobals := func(set *flag.FlagSet) {
		debug = set.Bool("debug", false, "print debug output")
		help = set.Bool("help", false, "print this information")
		h = set.Bool("h", false, "print this information")
	}
	idx := 1
	var ()
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			break
		} else if cmd == "" {
			cmd = arg
		} else if pkg == "" {
			pkg = arg
		} else {
			addGlobals(gset)
			showUsage()
		}
		idx++
	}
	switch cmd {
	case "gen":
		addGlobals(genset)
		genset.Parse(os.Args[idx:])
	case "genfull":
		addGlobals(genfullset)
		genfullset.Parse(os.Args[idx:])
	default:
		addGlobals(gset)
		genset.Parse(os.Args[idx:])
	}
	if *h || *help {
		showUsage()
		os.Exit(0)
	}

	return cmd, pkg, debug, out, pdir, showUsage
}

func main() {
	cmd, pkg, debug, out, dir, showUsage := processArgs()
	var err error
	switch cmd {
	case "gen":
		if pkg == "" {
			fail(`missing PACKAGE argument, use "--help" for usage`)
		}
		var b []byte
		dir, _ = filepath.Abs(dir)
		if err := os.MkdirAll(dir, 0777); err != nil {
			fail(err.Error())
		}
		b, err = gen(pkg, dir, *debug, *out)
		if err == nil {
			err = os.WriteFile(*out, b, 0644)
		}
	case "version":
		fmt.Printf("%s %s\n", os.Args[0], Pkg.Version())
	case "", "help":
		showUsage()
	default:
		fail(`unknown command %q, use "--help" for usage`, cmd)
	}
	if err != nil {
		fail(err.Error())
	}

}

func printUsage(fss ...*flag.FlagSet) {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintf(os.Stderr, "  %s serve PACKAGE [FLAGS].\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "    Start a HTTP server that serves a graphical editor for the design described in PACKAGE.\n")
	fmt.Fprintf(os.Stderr, "  %s gen PACKAGE [FLAGS].\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "    Generate a JSON representation of the design described in PACKAGE.\n")
	fmt.Fprintf(os.Stderr, "\nPACKAGE must be the import path to a Go package containing Model DSL.\n\n")
	fmt.Fprintf(os.Stderr, "FLAGS:\n")
	for _, fs := range fss {
		fs.PrintDefaults()
	}
}
func fail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
