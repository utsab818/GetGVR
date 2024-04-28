package main

import (
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	// res := "pods"

	var res string
	flag.StringVar(&res, "res", "", "resource that you want to interact with")
	flag.Parse()

	configFlag := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchVersionFlags := cmdutil.NewMatchVersionFlags(configFlag)
	m, err := cmdutil.NewFactory(matchVersionFlags).ToRESTMapper()
	if err != nil {
		fmt.Printf("getting rest mapper from newfactory %s\n", err.Error())
	}

	gvr, err := m.ResourceFor(schema.GroupVersionResource{
		Resource: res,
	})
	if err != nil {
		fmt.Printf("getting GVR for resource %s\n", err.Error())
	}
	fmt.Printf("Complete GVR is, group %s, version %s, resource %s\n", gvr.Group, gvr.Version, gvr.Resource)
}
