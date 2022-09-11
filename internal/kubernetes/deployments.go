package kubernetes

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

func (k *Kubernetes) GetDeployments() []Deployment {
	result := []Deployment{}
	var listEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.AppsV1().Deployments("").List(k.cfg.Ctx, listEverything)
	if err != nil {
		panic(err.Error())
	}
	replicasets := k.GetReplicasets()
	for _, deployment := range res.Items {
		item := Deployment{
			Name:        deployment.Name,
			Namespace:   deployment.Namespace,
			Annotations: deployment.Annotations,
		}

		// Filter non-active replicasets
		if *deployment.Spec.Replicas == 0 {
			continue
		}

		// Add pods
		for _, replicaset := range replicasets {
			if replicaset.Owner.Name != deployment.Name {
				continue
			}
			if replicaset.Owner.Kind != "Deployment" {
				continue
			}
			item.ReplicaSets = append(item.ReplicaSets, replicaset)
		}

		// Filter replicasets without replicasets
		if len(item.ReplicaSets) == 0 {
			continue
		}

		result = append(result, item)
	}
	return result
}
