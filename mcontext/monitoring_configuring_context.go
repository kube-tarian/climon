package mcontext

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/monitoring/api/v1alpha1"
	//"helm.sh/helm/v3/pkg/action"
)

var (
	monConfContext *MonConfContext
)

// MonConfContext
type MonConfContext struct {
	client *http.Client
}

// MonConf ...
type MonConf struct {
	Apps         json.RawMessage `json:"apps"`
	LastModifier string          `json:"last_modifier"`
	Metadata     json.RawMessage `json:"metadata"`
	Name         string          `json:"name"`
}

// GetPrometheusDeploymentContext
func GetPrometheusDeploymentContext() *MonConfContext {
	return monConfContext
}

// NewPrometheusDeploymentContext
func NewPrometheusDeploymentContext() *MonConfContext {
	if monConfContext == nil {
		monConfContext = &MonConfContext{client: &http.Client{Timeout: 20 * time.Second}}
	}

	return monConfContext
}

// HandleEvent
func (p *MonConfContext) HandleEvent(event string, instance *v1alpha1.MonitoringConfiguration) error {
	monConf := copytoLocalMonConf(instance)

	switch event {
	case "delete":
		return p.DeleteMonitoringConfiguration()
	case "createOrUpdate":
		// Check if the instance exists, if not create.
		if p.checkIfMonitoringConfExists(instance.Spec.Name) {
			//Invoke Helm upgrade
			return p.UpdateMonitoringConfiguration(monConf)
		}
		return p.CreateMonitoringConfiguration(monConf)
	}
	return nil
}

func (p *MonConfContext) CreateMonitoringConfiguration(promDep *MonConf) error {
	return nil
}

func (p *MonConfContext) UpdateMonitoringConfiguration(promDep *MonConf) error {
	return nil
}

func (p *MonConfContext) checkIfMonitoringConfExists(name string) bool {
	return false
}

func (p *MonConfContext) DeleteMonitoringConfiguration() error {
	return nil
}

func copytoLocalMonConf(instance *v1alpha1.MonitoringConfiguration) *MonConf {
	return &MonConf{
		Name:         instance.Spec.Name,
		LastModifier: instance.Spec.LastModifier,
		Metadata:     json.RawMessage(instance.Spec.Metadata),
	}
}
