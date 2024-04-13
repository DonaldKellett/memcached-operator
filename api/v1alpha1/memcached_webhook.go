/*
Copyright 2024.

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

package v1alpha1

import (
	"errors"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var memcachedlog = logf.Log.WithName("memcached-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Memcached) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-cache-donaldsebleung-com-v1alpha1-memcached,mutating=true,failurePolicy=fail,sideEffects=None,groups=cache.donaldsebleung.com,resources=memcacheds,verbs=create;update,versions=v1alpha1,name=mmemcached.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Memcached{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Memcached) Default() {
	memcachedlog.Info("default", "name", r.Name)

	if r.Spec.Size == 0 {
		r.Spec.Size = 1
	}
	if r.Spec.ContainerPort == 0 {
		r.Spec.ContainerPort = 11211
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-cache-donaldsebleung-com-v1alpha1-memcached,mutating=false,failurePolicy=fail,sideEffects=None,groups=cache.donaldsebleung.com,resources=memcacheds,verbs=create;update,versions=v1alpha1,name=vmemcached.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Memcached{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateCreate() (admission.Warnings, error) {
	memcachedlog.Info("validate create", "name", r.Name)

	return nil, r.validateMemcached()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	memcachedlog.Info("validate update", "name", r.Name)

	if oldMemcached, ok := old.(*Memcached); !ok {
		return nil, errors.New("failed asserting that old runtime.Object is of type *Memcached")
	} else if oldMemcached.Spec.ContainerPort != r.Spec.ContainerPort {
		return nil, field.Invalid(field.NewPath("spec").Child("containerPort"), r.Name, "containerPort field cannot be modified after creation")
	} else {
		return nil, r.validateMemcached()
	}
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateDelete() (admission.Warnings, error) {
	memcachedlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}

// Common validation logic for Memcached
func (r *Memcached) validateMemcached() error {
	var allErrs field.ErrorList
	if err := r.validateMemcachedName(); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := r.validateMemcachedSpec(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "cache.donaldsebleung.com", Kind: "Memcached"},
		r.Name, allErrs)
}

// Validate Memcached resource name
func (r *Memcached) validateMemcachedName() *field.Error {
	if len(r.ObjectMeta.Name) > validationutils.DNS1035LabelMaxLength-16 {
		// The pod name length is 63 characters like all Kubernetes objects
		// (which must fit in a DNS subdomain). The Memcached operator owns
		// a Deployment which appends a 16-character suffix to each Pod it
		// creates. The Pod name length limit is 63 characters. Therefore
		// Memcached names must have length <=63-16=47. If we don't validate
		// this here, then Pod creation will fail later.
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 47 characters")
	}
	return nil
}

// Validate Memcached resource spec
func (r *Memcached) validateMemcachedSpec() *field.Error {
	if err := validateSize(r.Spec.Size, field.NewPath("spec").Child("size")); err != nil {
		return err
	}
	if err := validateContainerPort(r.Spec.ContainerPort, field.NewPath("spec").Child("containerPort")); err != nil {
		return err
	}
	return nil
}

// Validate Memcached size
func validateSize(size int32, fldPath *field.Path) *field.Error {
	if size < 1 || size > 5 {
		return field.Invalid(fldPath, size, "must be between 1 and 5, both inclusive")
	}
	return nil
}

// Validate Memcached container port
func validateContainerPort(containerPort int32, fldPath *field.Path) *field.Error {
	if containerPort <= 0 || containerPort >= 65536 {
		return field.Invalid(fldPath, containerPort, "must be between 0 and 65536, both exclusive")
	}
	return nil
}
