package kubernetes

func (k *Kubernetes) GetResources() []Resource {
	resources := []Resource{}
	resources = append(resources, k.GetDeployments()...)
	resources = append(resources, k.GetStatefulsets()...)
	resources = append(resources, k.GetDaemonsets()...)
	return resources
}
