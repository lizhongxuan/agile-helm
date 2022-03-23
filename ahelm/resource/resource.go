package resource

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
import "k8s.io/cli-runtime/pkg/resource"

type Resource struct {

}

func (r *Resource)Create(info *resource.Info) error {
	obj, err := resource.NewHelper(info.Client, info.Mapping).Create(info.Namespace, true, info.Object)
	if err != nil {
		return err
	}
	return info.Refresh(obj, true)
}

func (r *Resource)Delete(info *resource.Info) error {
	policy := metav1.DeletePropagationBackground
	opts := &metav1.DeleteOptions{PropagationPolicy: &policy}
	_, err := resource.NewHelper(info.Client, info.Mapping).DeleteWithOptions(info.Namespace, info.Name, opts)
	return err
}


func (r *Resource)Update( info *resource.Info) error {
	_, err := resource.NewHelper(info.Client, info.Mapping).Replace(info.Namespace, info.Name, true, info.Object)
	if err != nil {
		return errors.Wrap(err, "failed to replace object")
	}
	return nil
}
