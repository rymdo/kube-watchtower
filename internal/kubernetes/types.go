package kubernetes

type Owner struct {
	Name string
	Kind string
}

type Container struct {
	Name    string
	Image   string
	ImageID string
}

type Pod struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	Containers  []Container
	Owner       Owner
}

type ReplicaSet struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	Pods        []Pod
	Owner       Owner
}

type Deployment struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	ReplicaSets []ReplicaSet
}

type StatefulSet struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	Pods        []Pod
}

type DaemonSet struct {
	Name        string
	Namespace   string
	Annotations map[string]string
	Pods        []Pod
}
