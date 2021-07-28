/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	monitoringv1alpha1 "github.com/climon-operator/api/v1alpha1"
)

const (
	requeueDelay = time.Second * 10
	//Name of our custom finalizer
	myFinalizerName = "finalizer.prometheusdeployments.climon.com"
)

// PrometheusDeploymentReconciler reconciles a PrometheusDeployment object
type PrometheusDeploymentReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=monitoring.climon.com,resources=prometheusdeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.climon.com,resources=prometheusdeployments/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PrometheusDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.6.4/pkg/reconcile
func (r *PrometheusDeploymentReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("prometheusdeployment", req.NamespacedName)

	// your logic here

	instance := &monitoringv1alpha1.PrometheusDeployment{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)

	if err != nil {
		if k8sErrors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		//Error reading the object -- requeue the request
		return ctrl.Result{}, err
	}

	// examine DeletionTimestamp to determine if object is under deletion
	if instance.GetDeletionTimestamp().IsZero() {
		if !containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Client.Update(context.Background(), instance); err != nil {
				return reconcile.Result{RequeueAfter: requeueDelay}, err
			}
		}
	} else {
		// The object is being deleted
		if containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			// Our finalizer is present, so lets handled any external dependency
			r.Log.Info("Handle delete event for %v, Rule Name: %v", instance.Name, instance.Spec.Name)

		}
	}

	return ctrl.Result{}, nil
}

// Helper functions to check and remove string from a slice of strings.
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// SetupWithManager sets up the controller with the Manager.
func (r *PrometheusDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringv1alpha1.PrometheusDeployment{}).
		Complete(r)
}
