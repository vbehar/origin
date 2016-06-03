// conversion-gen is a tool for auto-generating Conversion functions.
//
// Structs in the input directories with the below line in their comments
// will be ignored during generation.
// // +genconversion=false
package main

import (
	"strings"

	"github.com/golang/glog"

	"k8s.io/kubernetes/cmd/libs/go2idl/args"
	"k8s.io/kubernetes/cmd/libs/go2idl/conversion-gen/generators"
	"k8s.io/kubernetes/cmd/libs/go2idl/generator"
)

func main() {
	arguments := args.Default()

	// Override defaults. These are Kubernetes specific input locations.
	arguments.InputDirs = []string{
		"k8s.io/kubernetes/pkg/api/v1",
		"k8s.io/kubernetes/pkg/api",
		"k8s.io/kubernetes/pkg/runtime",
		"k8s.io/kubernetes/pkg/conversion",
		"github.com/openshift/origin/pkg/authorization/api/v1",
		"github.com/openshift/origin/pkg/authorization/api",
		"github.com/openshift/origin/pkg/build/api/v1",
		"github.com/openshift/origin/pkg/build/api",
		"github.com/openshift/origin/pkg/deploy/api/v1",
		"github.com/openshift/origin/pkg/deploy/api",
		"github.com/openshift/origin/pkg/image/api/v1",
		"github.com/openshift/origin/pkg/image/api",
		"github.com/openshift/origin/pkg/oauth/api/v1",
		"github.com/openshift/origin/pkg/oauth/api",
		"github.com/openshift/origin/pkg/project/api/v1",
		"github.com/openshift/origin/pkg/project/api",
		"github.com/openshift/origin/pkg/route/api/v1",
		"github.com/openshift/origin/pkg/route/api",
		"github.com/openshift/origin/pkg/sdn/api/v1",
		"github.com/openshift/origin/pkg/sdn/api",
		"github.com/openshift/origin/pkg/template/api/v1",
		"github.com/openshift/origin/pkg/template/api",
		"github.com/openshift/origin/pkg/user/api/v1",
		"github.com/openshift/origin/pkg/user/api",
		"github.com/openshift/origin/pkg/security/api/v1",
		"github.com/openshift/origin/pkg/security/api",
	}

	arguments.GoHeaderFilePath = "hack/boilerplate.txt"

	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		func(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
			pkgs := generators.Packages(context, arguments)
			var include generator.Packages
			for _, pkg := range pkgs {
				if strings.HasPrefix(pkg.Path(), "k8s.io/") {
					continue
				}
				include = append(include, pkg)
			}
			return include
		},
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.Info("Completed successfully.")
}
