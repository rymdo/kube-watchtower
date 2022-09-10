package kubernetes

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func (k *Kubernetes) GetStatefulsets() []Resource {
	result := []Resource{}
	var ListEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.AppsV1().StatefulSets("").List(k.cfg.Ctx, ListEverything)
	if err != nil {
		panic(err.Error())
	}
	for _, s := range res.Items {
		containers := []ResourceContainer{}
		for _, c := range s.Spec.Template.Spec.Containers {
			containers = append(containers, ResourceContainer{
				Name:  c.Name,
				Image: c.Image,
			})
		}
		result = append(result, Resource{
			Type:        Statefulset,
			Name:        s.Name,
			Namespace:   s.Namespace,
			Annotations: s.Annotations,
			Containers:  containers,
		})
	}
	return result
}
