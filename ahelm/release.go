package ahelm

import (
	"agile-helm/pkg/action"
	"agile-helm/pkg/release"
	"context"
	log "github.com/golang/glog"
)

type Chart struct{}

func (*Chart) ListRelease( ns string) ([]*release.Release, error) {
	actionCfg, errRes := GetActionConfig( ns)
	if errRes != nil {
		return nil, errRes
	}

	allNamespaces := false
	if ns == "" {
		allNamespaces = true
	}
	listAction := action.NewList(actionCfg)
	listAction.AllNamespaces = allNamespaces
	list, err := listAction.Run()
	if err != nil {
		log.Error( err)
		return nil, err
	}
	return list,err
}

func (*Chart) InfoRelase(ns, rlsName string) (*release.Release, error) {
	actionCfg, err := GetActionConfig(ns)
	if err != nil {
		return nil, err
	}

	rls, err := action.NewGet(actionCfg).Run(rlsName)
	if err != nil {
		return nil,err
	}
	return rls,nil
}

func (*Chart) UninstallRelease(ns, rlsName string) error {
	actionCfg, errRes := GetActionConfig( ns)
	if errRes != nil {
		return errRes
	}

	if _, err := action.NewUninstall(actionCfg).Run(rlsName); err != nil {
		log.Error( err)
		return err
	}
	return nil
}

func (*Chart) Args(ctx context.Context,  ns, rlsName string, allArg bool) (map[string]interface{}, error) {
	actionCfg, err := GetActionConfig( ns)
	if err != nil {
		return nil, err
	}

	argAction := action.NewGetValues(actionCfg)
	argAction.AllValues = allArg
	args, err := argAction.Run(rlsName)
	if err != nil {
		log.Error( err)
		return nil, err
	}
	return args, nil
}

