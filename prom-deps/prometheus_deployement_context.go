package promdeps

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/climon-operator/api/v1alpha1"
	//"helm.sh/helm/v3/pkg/action"
)

var (
	promdepsContext *PrometheusDeploymentContext
)

// PrometheusDeploymentContext
type PrometheusDeploymentContext struct {
	client *http.Client
}

// PrometheusDeployment ...
type PrometheusDeployment struct {
	Prometheus   json.RawMessage `json:"prometheus"`
	LastModifier string          `json:"last_modifier"`
	Metadata     json.RawMessage `json:"metadata"`
	Name         string          `json:"name"`
}

// GetPrometheusDeploymentContext
func GetPrometheusDeploymentContext() *PrometheusDeploymentContext {
	return promdepsContext
}

// NewPrometheusDeploymentContext
func NewPrometheusDeploymentContext() *PrometheusDeploymentContext {
	if promdepsContext == nil {
		promdepsContext = &PrometheusDeploymentContext{client: &http.Client{Timeout: 20 * time.Second}}
	}

	return promdepsContext
}

// HandleEvent
func (p *PrometheusDeploymentContext) HandleEvent(event string, instance *v1alpha1.PrometheusDeployment) error {
	promDep := copytoLocalPromDep(instance)

	switch event {
	case "delete":
		return p.DeletePrometheusDeployment()
	case "createOrUpdate":
		// Check if the instance exists, if not create.
		if p.checkIfPrometheusDepExists(instance.Spec.Name) {
			//Invoke Helm upgrade
			return p.UpdatePrometheusDeployment(promDep)
		}
		return p.CreatePrometheusDeployment(promDep)
	}
	return nil
}

func (p *PrometheusDeploymentContext) CreatePrometheusDeployment(promDep *PrometheusDeployment) error {
	return nil
}

func (p *PrometheusDeploymentContext) UpdatePrometheusDeployment(promDep *PrometheusDeployment) error {
	return nil
}

func (p *PrometheusDeploymentContext) checkIfPrometheusDepExists(name string) bool {
	return false
}

func (p *PrometheusDeploymentContext) DeletePrometheusDeployment() error {
	return nil
}

func copytoLocalPromDep(instance *v1alpha1.PrometheusDeployment) *PrometheusDeployment {
	return &PrometheusDeployment{
		Name:         instance.Spec.Name,
		LastModifier: instance.Spec.LastModifier,
		Prometheus:   json.RawMessage(instance.Spec.Prometheus),
		Metadata:     json.RawMessage(instance.Spec.Metadata),
	}
}
