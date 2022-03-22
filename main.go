package main

import (
	"agile-helm/helm/v3/pkg/action"
	"agile-helm/pkg"
	"fmt"
)

func main() {
	actcfg, err := pkg.GetActionConfig("agile-helm")
	if err != nil {
		fmt.Println(err)
		return
	}
	list, err := action.NewList(actcfg).Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
