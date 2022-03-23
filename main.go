package main

import (
	"agile-helm/pkg/action"
	"agile-helm/ahelm"
	"fmt"
)

func main() {
	actcfg, err := ahelm.GetActionConfig("agile-helm")
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
