package kubernetes

import (
	"regexp"

	v1 "k8s.io/api/core/v1"
)

func cleanupImageIdSha(id string) string {
	regexsha := `(sha256:.*)`
	re := regexp.MustCompile(regexsha)
	matches := re.FindStringSubmatch(id)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func getContainerImageID(p v1.Pod, c v1.Container) string {
	var imageID string
	for _, cs := range p.Status.ContainerStatuses {
		if c.Image == cs.Image {
			imageID = cleanupImageIdSha(cs.ImageID)
		}
	}
	return imageID
}
