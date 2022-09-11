package kubernetes

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func (k *Kubernetes) GetReplicasets() []ReplicaSet {
	result := []ReplicaSet{}
	var listEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.AppsV1().ReplicaSets("").List(k.cfg.Ctx, listEverything)
	if err != nil {
		panic(err.Error())
	}
	pods := k.GetPods()
	for _, replicaset := range res.Items {
		item := ReplicaSet{
			Name:        replicaset.Name,
			Namespace:   replicaset.Namespace,
			Annotations: replicaset.Annotations,
		}

		// Filter non-active replicasets
		if *replicaset.Spec.Replicas == 0 {
			continue
		}

		// Add pods
		for _, pod := range pods {
			if pod.Owner.Name != replicaset.Name {
				continue
			}
			if pod.Owner.Kind != "ReplicaSet" {
				continue
			}
			item.Pods = append(item.Pods, pod)
		}

		// Add owner
		if len(replicaset.OwnerReferences) > 0 {
			item.Owner = Owner{
				Name: replicaset.OwnerReferences[0].Name,
				Kind: replicaset.OwnerReferences[0].Kind,
			}
		}

		result = append(result, item)
	}
	return result
}
