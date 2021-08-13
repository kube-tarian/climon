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
	monitoringv1alpha1 "github.com/monitoring/api/v1alpha1"
	"github.com/monitoring/mcontext"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	requeueDelay = time.Second * 10
	//Name of our custom finalizer
	myFinalizerName = "finalizer.monitoringconfigurations.soi.dev"
)

// MonitoringConfigurationReconciler reconciles a MonitoringConfiguration object
type MonitoringConfigurationReconciler struct {
	Client         client.Client
	Log            logr.Logger
	Scheme         *runtime.Scheme
	monConfContext *mcontext.MonConfContext
}

//+kubebuilder:rbac:groups=monitoring.soi.dev,resources=monitoringconfigurations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=monitoring.soi.dev,resources=monitoringconfigurations/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MonitoringConfiguration object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.6.4/pkg/reconcile
func (r *MonitoringConfigurationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("monitoringconfiguration", req.NamespacedName)

	instance := &monitoringv1alpha1.MonitoringConfiguration{}
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

	//Examine DeletionTimestamp to determine if the object is under deletion
	if instance.GetDeletionTimestamp().IsZero() {
		if !containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Client.Update(context.Background(), instance); err != nil {
				return reconcile.Result{RequeueAfter: requeueDelay}, err
			}
		}
	} else {
		//The object is being deleted
		if containsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			// Our finalizer is present, so lets handled any external dependency
			r.Log.Info("Handle delete event for %v, MonitoringConfiguration Name: %v", instance.Name, instance.Spec.Name)

			if err := r.monConfContext.HandleEvent("delete", instance); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return reconcile.Result{RequeueAfter: requeueDelay}, err
			}

			// remove our finalizer from the list and update it.
			instance.ObjectMeta.Finalizers = removeString(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Client.Update(context.Background(), instance); err != nil {
				return reconcile.Result{RequeueAfter: requeueDelay}, err
			}
		}
		// Stop reconciliation as the item is being deleted
		return reconcile.Result{}, nil
	}

	//Do this in case of create or an update event
	r.Log.Info("Update or CreateOrUpdate event occured for %+v, MonitoringConfiguration Name: %v", instance, instance.Spec.Name)

	err = r.monConfContext.HandleEvent("createOrUpdate", instance)
	if err != nil {
		r.Log.Error(err, "Error reconciling crd during creation or updation")
		return reconcile.Result{RequeueAfter: requeueDelay}, err
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

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

// SetupWithManager sets up the controller with the Manager.
func (r *MonitoringConfigurationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitoringv1alpha1.MonitoringConfiguration{}).
		Complete(r)
}
