package discovery

import (
	"fmt"

	"github.com/wavefronthq/wavefront-kubernetes-collector/internal/metrics"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceType int

const (
	PodType     ResourceType = 1
	ServiceType ResourceType = 2
	IngressType ResourceType = 3
)

func (resType ResourceType) String() string {
	switch resType {
	case PodType:
		return "pod"
	case ServiceType:
		return "service"
	case IngressType:
		return "ingress"
	default:
		return fmt.Sprintf("%d", int(resType))
	}
}

type Manager interface {
	ListPods(ns string, labels map[string]string) ([]*v1.Pod, error)
	ListServices(ns string, labels map[string]string) ([]*v1.Service, error)
	RegisterProvider(provider metrics.MetricsSourceProvider)
	UnregisterProvider(providerName string)
}

type Discoverer interface {
	Discover(ip, kind string, obj metav1.ObjectMeta)
	Delete(kind string, obj metav1.ObjectMeta)
	Process(config Config)
}
