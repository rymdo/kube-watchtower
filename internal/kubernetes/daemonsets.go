package kubernetes

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func (k *Kubernetes) GetDaemonsets() []DaemonSet {
	result := []DaemonSet{}
	var listEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.AppsV1().DaemonSets("").List(k.cfg.Ctx, listEverything)
	if err != nil {
		panic(err.Error())
	}
	pods := k.GetPods()
	for _, daemonset := range res.Items {
		item := DaemonSet{
			Name:        daemonset.Name,
			Namespace:   daemonset.Namespace,
			Annotations: daemonset.Annotations,
		}

		// Add pods
		for _, pod := range pods {
			if pod.Owner.Name != daemonset.Name {
				continue
			}
			if pod.Owner.Kind != "DaemonSet" {
				continue
			}
			item.Pods = append(item.Pods, pod)
		}

		result = append(result, item)
	}
	return result
}
