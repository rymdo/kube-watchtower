package kubernetes

func (k *Kubernetes) GetResource() []Resource {
	resources := []Resource{}
	resources = append(resources, k.GetDeployments()...)
	resources = append(resources, k.GetStatefulsets()...)
	resources = append(resources, k.GetDaemonsets()...)
	return resources
}
