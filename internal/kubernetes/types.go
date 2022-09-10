package kubernetes

type ResourceType string

const (
	Deployment  ResourceType = "deployment"
	Statefulset ResourceType = "statefulset"
	Daemonset   ResourceType = "daemonset"
)

type ResourceContainer struct {
	Name  string
	Image string
}

type Resource struct {
	Type        ResourceType
	Name        string
	Namespace   string
	Annotations map[string]string
	Containers  []ResourceContainer
}
