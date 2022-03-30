package main

import (
	"agile-helm/pkg/action"
	"agile-helm/pkg/actioncfg"
	"fmt"
)

func main() {
	actcfg, err := actioncfg.InClusterActionCfg("helm-test")
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
