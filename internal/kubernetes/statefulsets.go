package kubernetes

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func (k *Kubernetes) GetStatefulsets() []StatefulSet {
	result := []StatefulSet{}
	var listEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.AppsV1().StatefulSets("").List(k.cfg.Ctx, listEverything)
	if err != nil {
		panic(err.Error())
	}
	pods := k.GetPods()
	for _, statefulset := range res.Items {
		item := StatefulSet{
			Name:        statefulset.Name,
			Namespace:   statefulset.Namespace,
			Annotations: statefulset.Annotations,
		}

		// Filter non-active statefulsets
		if *statefulset.Spec.Replicas == 0 {
			continue
		}

		// Add pods
		for _, pod := range pods {
			if pod.Owner.Name != statefulset.Name {
				continue
			}
			if pod.Owner.Kind != "StatefulSet" {
				continue
			}
			item.Pods = append(item.Pods, pod)
		}

		result = append(result, item)
	}
	return result
}
