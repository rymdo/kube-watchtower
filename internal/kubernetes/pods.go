package kubernetes

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func (k *Kubernetes) GetPods() []Pod {
	result := []Pod{}
	var listEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.CoreV1().Pods("").List(k.cfg.Ctx, listEverything)
	if err != nil {
		panic(err.Error())
	}
	for _, p := range res.Items {
		item := Pod{
			Name:        p.Name,
			Namespace:   p.Namespace,
			Annotations: p.Annotations,
		}

		// Add containers
		for _, c := range p.Spec.Containers {
			item.Containers = append(item.Containers, Container{
				Name:    c.Name,
				Image:   c.Image,
				ImageID: getContainerImageID(p, c),
			})
		}

		// Add owner
		if len(p.OwnerReferences) > 0 {
			item.Owner.Name = p.OwnerReferences[0].Name
			item.Owner.Kind = p.OwnerReferences[0].Kind
		}

		result = append(result, item)
	}
	return result
}
