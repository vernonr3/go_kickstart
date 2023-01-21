package expr

import (
	"fmt"

	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"

	//model "goa.design/model/pkg"
	model "go_kickstart/pkg"
)

type (
	// Design contains the AST generated from the DSL.
	Design struct {
		Name        string
		Description string
		Version     string
		Model       *Model
		Views       *Views
	}
)

// Root is the design root expression.
var Root = &Design{Model: &Model{}, Views: &Views{}}

// Register design root with eval engine.
func init() {
	eval.Register(Root)
}

// WalkSets iterates over the elements and views.
// Elements DSL cannot be executed on init because all elements must first be
// loaded and their IDs captured in the registry before relationships can be
// built with DSL.
func (d *Design) WalkSets(walk eval.SetWalker) {
	// 1. Model
	walk([]eval.Expression{d.Model})
	// 2. People
	walk(eval.ToExpressionSet(d.Model.People))
	// 3. Systems
	walk(eval.ToExpressionSet(d.Model.Systems))
	// 4. Containers
	for _, s := range d.Model.Systems {
		walk(eval.ToExpressionSet(s.Containers))
	}
	// 5. Components
	for _, s := range d.Model.Systems {
		for _, c := range s.Containers {
			walk(eval.ToExpressionSet(c.Components))
		}
	}
	// 6. Structs & Interfaces
	for _, s := range d.Model.Systems {
		for _, c := range s.Containers {
			for _, d := range c.Components {
				walk(eval.ToExpressionSet(d.Structs))
				walk(eval.ToExpressionSet(d.Interfaces))
			}

		}
	}
	// 7. Methods & interface inside structs
	for _, s := range d.Model.Systems {
		for _, c := range s.Containers {
			for _, d := range c.Components {
				for _, e := range d.Structs {
					walk(eval.ToExpressionSet(e.Interfaces))
					walk(eval.ToExpressionSet(e.Methods))
				}
				for _, e := range d.Interfaces {
					walk(eval.ToExpressionSet(e.Methods))
				}
			}

		}
	}
	// 7. Methods inside interfaces
	for _, s := range d.Model.Systems {
		for _, c := range s.Containers {
			for _, d := range c.Components {
				for _, e := range d.Structs {
					for _, f := range e.Interfaces {
						walk(eval.ToExpressionSet(f.Methods))
					}
				}
			}
		}
	}
	// 6. Functions
	for _, s := range d.Model.Systems {
		for _, c := range s.Containers {
			for _, d := range c.Components {
				walk(eval.ToExpressionSet(d.Functions))
			}

		}
	}
	// 9. Deployment environments
	walkDeploymentNodes(d.Model.DeploymentNodes, walk)
	// 10. Views
	walk([]eval.Expression{d.Views})
}

// Packages returns the import path to the Go packages that make
// up the DSL. This is used to skip frames that point to files
// in these packages when computing the location of errors.
func (d *Design) Packages() []string {
	return []string{
		"goa.design/model/expr",
		"goa.design/model/dsl",
		fmt.Sprintf("goa.design/model@%s/expr", model.Version()),
		fmt.Sprintf("goa.design/model@%s/dsl", model.Version()),
	}
}

// DependsOn tells the eval engine to run the goa DSL first.
func (d *Design) DependsOn() []eval.Root { return []eval.Root{expr.Root} }

// EvalName returns the generic expression name used in error messages.
func (d *Design) EvalName() string { return "root" }

func walkDeploymentNodes(n []*DeploymentNode, walk eval.SetWalker) {
	if n == nil {
		return
	}
	walk(eval.ToExpressionSet(n))
	for _, d := range n {
		walk(eval.ToExpressionSet(d.InfrastructureNodes))
		walk(eval.ToExpressionSet(d.ContainerInstances))
		walkDeploymentNodes(d.Children, walk)
	}
}

// Person returns the person with the given name if any, nil otherwise.
func (d *Design) Person(name string) *Person {
	return d.Model.Person(name)
}

// SoftwareSystem returns the software system with the given name if any, nil
// otherwise.
func (d *Design) SoftwareSystem(name string) *SoftwareSystem {
	return d.Model.SoftwareSystem(name)
}

// DeploymentNode returns the deployment node with the given name in the given
// environment if any, nil otherwise.
func (d *Design) DeploymentNode(env, name string) *DeploymentNode {
	return d.Model.DeploymentNode(env, name)
}
