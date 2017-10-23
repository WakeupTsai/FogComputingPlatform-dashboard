/*
Copyright 2017 The Kubernetes Authors.

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

package kubernetes

import (
	glog "github.com/golang/glog"
	discovery "k8s.io/client-go/discovery"
	admissionregistrationv1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
<<<<<<< HEAD
=======
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
>>>>>>> upstream/master
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
<<<<<<< HEAD
	autoscalingv2alpha1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2alpha1"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
=======
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
>>>>>>> upstream/master
	batchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
<<<<<<< HEAD
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
=======
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
>>>>>>> upstream/master
	settingsv1alpha1 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AdmissionregistrationV1alpha1() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Admissionregistration() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface
<<<<<<< HEAD
	CoreV1() corev1.CoreV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Core() corev1.CoreV1Interface
	AppsV1beta1() appsv1beta1.AppsV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Apps() appsv1beta1.AppsV1beta1Interface
=======
	AppsV1beta1() appsv1beta1.AppsV1beta1Interface
	AppsV1beta2() appsv1beta2.AppsV1beta2Interface
	// Deprecated: please explicitly pick a version if possible.
	Apps() appsv1beta2.AppsV1beta2Interface
>>>>>>> upstream/master
	AuthenticationV1() authenticationv1.AuthenticationV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Authentication() authenticationv1.AuthenticationV1Interface
	AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface
	AuthorizationV1() authorizationv1.AuthorizationV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Authorization() authorizationv1.AuthorizationV1Interface
	AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface
	AutoscalingV1() autoscalingv1.AutoscalingV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Autoscaling() autoscalingv1.AutoscalingV1Interface
<<<<<<< HEAD
	AutoscalingV2alpha1() autoscalingv2alpha1.AutoscalingV2alpha1Interface
	BatchV1() batchv1.BatchV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Batch() batchv1.BatchV1Interface
=======
	AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface
	BatchV1() batchv1.BatchV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Batch() batchv1.BatchV1Interface
	BatchV1beta1() batchv1beta1.BatchV1beta1Interface
>>>>>>> upstream/master
	BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface
	CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Certificates() certificatesv1beta1.CertificatesV1beta1Interface
<<<<<<< HEAD
=======
	CoreV1() corev1.CoreV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Core() corev1.CoreV1Interface
>>>>>>> upstream/master
	ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Extensions() extensionsv1beta1.ExtensionsV1beta1Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Networking() networkingv1.NetworkingV1Interface
	PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Policy() policyv1beta1.PolicyV1beta1Interface
<<<<<<< HEAD
	RbacV1beta1() rbacv1beta1.RbacV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Rbac() rbacv1beta1.RbacV1beta1Interface
	RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface
=======
	RbacV1() rbacv1.RbacV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Rbac() rbacv1.RbacV1Interface
	RbacV1beta1() rbacv1beta1.RbacV1beta1Interface
	RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface
	SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface
>>>>>>> upstream/master
	SettingsV1alpha1() settingsv1alpha1.SettingsV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Settings() settingsv1alpha1.SettingsV1alpha1Interface
	StorageV1beta1() storagev1beta1.StorageV1beta1Interface
	StorageV1() storagev1.StorageV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Storage() storagev1.StorageV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
<<<<<<< HEAD
	*admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Client
	*corev1.CoreV1Client
	*appsv1beta1.AppsV1beta1Client
	*authenticationv1.AuthenticationV1Client
	*authenticationv1beta1.AuthenticationV1beta1Client
	*authorizationv1.AuthorizationV1Client
	*authorizationv1beta1.AuthorizationV1beta1Client
	*autoscalingv1.AutoscalingV1Client
	*autoscalingv2alpha1.AutoscalingV2alpha1Client
	*batchv1.BatchV1Client
	*batchv2alpha1.BatchV2alpha1Client
	*certificatesv1beta1.CertificatesV1beta1Client
	*extensionsv1beta1.ExtensionsV1beta1Client
	*networkingv1.NetworkingV1Client
	*policyv1beta1.PolicyV1beta1Client
	*rbacv1beta1.RbacV1beta1Client
	*rbacv1alpha1.RbacV1alpha1Client
	*settingsv1alpha1.SettingsV1alpha1Client
	*storagev1beta1.StorageV1beta1Client
	*storagev1.StorageV1Client
=======
	admissionregistrationV1alpha1 *admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Client
	appsV1beta1                   *appsv1beta1.AppsV1beta1Client
	appsV1beta2                   *appsv1beta2.AppsV1beta2Client
	authenticationV1              *authenticationv1.AuthenticationV1Client
	authenticationV1beta1         *authenticationv1beta1.AuthenticationV1beta1Client
	authorizationV1               *authorizationv1.AuthorizationV1Client
	authorizationV1beta1          *authorizationv1beta1.AuthorizationV1beta1Client
	autoscalingV1                 *autoscalingv1.AutoscalingV1Client
	autoscalingV2beta1            *autoscalingv2beta1.AutoscalingV2beta1Client
	batchV1                       *batchv1.BatchV1Client
	batchV1beta1                  *batchv1beta1.BatchV1beta1Client
	batchV2alpha1                 *batchv2alpha1.BatchV2alpha1Client
	certificatesV1beta1           *certificatesv1beta1.CertificatesV1beta1Client
	coreV1                        *corev1.CoreV1Client
	extensionsV1beta1             *extensionsv1beta1.ExtensionsV1beta1Client
	networkingV1                  *networkingv1.NetworkingV1Client
	policyV1beta1                 *policyv1beta1.PolicyV1beta1Client
	rbacV1                        *rbacv1.RbacV1Client
	rbacV1beta1                   *rbacv1beta1.RbacV1beta1Client
	rbacV1alpha1                  *rbacv1alpha1.RbacV1alpha1Client
	schedulingV1alpha1            *schedulingv1alpha1.SchedulingV1alpha1Client
	settingsV1alpha1              *settingsv1alpha1.SettingsV1alpha1Client
	storageV1beta1                *storagev1beta1.StorageV1beta1Client
	storageV1                     *storagev1.StorageV1Client
>>>>>>> upstream/master
}

// AdmissionregistrationV1alpha1 retrieves the AdmissionregistrationV1alpha1Client
func (c *Clientset) AdmissionregistrationV1alpha1() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AdmissionregistrationV1alpha1Client
=======
	return c.admissionregistrationV1alpha1
>>>>>>> upstream/master
}

// Deprecated: Admissionregistration retrieves the default version of AdmissionregistrationClient.
// Please explicitly pick a version.
func (c *Clientset) Admissionregistration() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AdmissionregistrationV1alpha1Client
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	if c == nil {
		return nil
	}
	return c.CoreV1Client
}

// Deprecated: Core retrieves the default version of CoreClient.
// Please explicitly pick a version.
func (c *Clientset) Core() corev1.CoreV1Interface {
	if c == nil {
		return nil
	}
	return c.CoreV1Client
=======
	return c.admissionregistrationV1alpha1
>>>>>>> upstream/master
}

// AppsV1beta1 retrieves the AppsV1beta1Client
func (c *Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AppsV1beta1Client
=======
	return c.appsV1beta1
}

// AppsV1beta2 retrieves the AppsV1beta2Client
func (c *Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return c.appsV1beta2
>>>>>>> upstream/master
}

// Deprecated: Apps retrieves the default version of AppsClient.
// Please explicitly pick a version.
<<<<<<< HEAD
func (c *Clientset) Apps() appsv1beta1.AppsV1beta1Interface {
	if c == nil {
		return nil
	}
	return c.AppsV1beta1Client
=======
func (c *Clientset) Apps() appsv1beta2.AppsV1beta2Interface {
	return c.appsV1beta2
>>>>>>> upstream/master
}

// AuthenticationV1 retrieves the AuthenticationV1Client
func (c *Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AuthenticationV1Client
=======
	return c.authenticationV1
>>>>>>> upstream/master
}

// Deprecated: Authentication retrieves the default version of AuthenticationClient.
// Please explicitly pick a version.
func (c *Clientset) Authentication() authenticationv1.AuthenticationV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AuthenticationV1Client
=======
	return c.authenticationV1
>>>>>>> upstream/master
}

// AuthenticationV1beta1 retrieves the AuthenticationV1beta1Client
func (c *Clientset) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AuthenticationV1beta1Client
=======
	return c.authenticationV1beta1
>>>>>>> upstream/master
}

// AuthorizationV1 retrieves the AuthorizationV1Client
func (c *Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AuthorizationV1Client
=======
	return c.authorizationV1
>>>>>>> upstream/master
}

// Deprecated: Authorization retrieves the default version of AuthorizationClient.
// Please explicitly pick a version.
func (c *Clientset) Authorization() authorizationv1.AuthorizationV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AuthorizationV1Client
=======
	return c.authorizationV1
>>>>>>> upstream/master
}

// AuthorizationV1beta1 retrieves the AuthorizationV1beta1Client
func (c *Clientset) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AuthorizationV1beta1Client
=======
	return c.authorizationV1beta1
>>>>>>> upstream/master
}

// AutoscalingV1 retrieves the AutoscalingV1Client
func (c *Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AutoscalingV1Client
=======
	return c.autoscalingV1
>>>>>>> upstream/master
}

// Deprecated: Autoscaling retrieves the default version of AutoscalingClient.
// Please explicitly pick a version.
func (c *Clientset) Autoscaling() autoscalingv1.AutoscalingV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.AutoscalingV1Client
}

// AutoscalingV2alpha1 retrieves the AutoscalingV2alpha1Client
func (c *Clientset) AutoscalingV2alpha1() autoscalingv2alpha1.AutoscalingV2alpha1Interface {
	if c == nil {
		return nil
	}
	return c.AutoscalingV2alpha1Client
=======
	return c.autoscalingV1
}

// AutoscalingV2beta1 retrieves the AutoscalingV2beta1Client
func (c *Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return c.autoscalingV2beta1
>>>>>>> upstream/master
}

// BatchV1 retrieves the BatchV1Client
func (c *Clientset) BatchV1() batchv1.BatchV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.BatchV1Client
=======
	return c.batchV1
>>>>>>> upstream/master
}

// Deprecated: Batch retrieves the default version of BatchClient.
// Please explicitly pick a version.
func (c *Clientset) Batch() batchv1.BatchV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.BatchV1Client
=======
	return c.batchV1
}

// BatchV1beta1 retrieves the BatchV1beta1Client
func (c *Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return c.batchV1beta1
>>>>>>> upstream/master
}

// BatchV2alpha1 retrieves the BatchV2alpha1Client
func (c *Clientset) BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.BatchV2alpha1Client
=======
	return c.batchV2alpha1
>>>>>>> upstream/master
}

// CertificatesV1beta1 retrieves the CertificatesV1beta1Client
func (c *Clientset) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.CertificatesV1beta1Client
=======
	return c.certificatesV1beta1
>>>>>>> upstream/master
}

// Deprecated: Certificates retrieves the default version of CertificatesClient.
// Please explicitly pick a version.
func (c *Clientset) Certificates() certificatesv1beta1.CertificatesV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.CertificatesV1beta1Client
=======
	return c.certificatesV1beta1
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}

// Deprecated: Core retrieves the default version of CoreClient.
// Please explicitly pick a version.
func (c *Clientset) Core() corev1.CoreV1Interface {
	return c.coreV1
>>>>>>> upstream/master
}

// ExtensionsV1beta1 retrieves the ExtensionsV1beta1Client
func (c *Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.ExtensionsV1beta1Client
=======
	return c.extensionsV1beta1
>>>>>>> upstream/master
}

// Deprecated: Extensions retrieves the default version of ExtensionsClient.
// Please explicitly pick a version.
func (c *Clientset) Extensions() extensionsv1beta1.ExtensionsV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.ExtensionsV1beta1Client
=======
	return c.extensionsV1beta1
>>>>>>> upstream/master
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.NetworkingV1Client
=======
	return c.networkingV1
>>>>>>> upstream/master
}

// Deprecated: Networking retrieves the default version of NetworkingClient.
// Please explicitly pick a version.
func (c *Clientset) Networking() networkingv1.NetworkingV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.NetworkingV1Client
=======
	return c.networkingV1
>>>>>>> upstream/master
}

// PolicyV1beta1 retrieves the PolicyV1beta1Client
func (c *Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.PolicyV1beta1Client
=======
	return c.policyV1beta1
>>>>>>> upstream/master
}

// Deprecated: Policy retrieves the default version of PolicyClient.
// Please explicitly pick a version.
func (c *Clientset) Policy() policyv1beta1.PolicyV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.PolicyV1beta1Client
}

// RbacV1beta1 retrieves the RbacV1beta1Client
func (c *Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	if c == nil {
		return nil
	}
	return c.RbacV1beta1Client
=======
	return c.policyV1beta1
}

// RbacV1 retrieves the RbacV1Client
func (c *Clientset) RbacV1() rbacv1.RbacV1Interface {
	return c.rbacV1
>>>>>>> upstream/master
}

// Deprecated: Rbac retrieves the default version of RbacClient.
// Please explicitly pick a version.
<<<<<<< HEAD
func (c *Clientset) Rbac() rbacv1beta1.RbacV1beta1Interface {
	if c == nil {
		return nil
	}
	return c.RbacV1beta1Client
=======
func (c *Clientset) Rbac() rbacv1.RbacV1Interface {
	return c.rbacV1
}

// RbacV1beta1 retrieves the RbacV1beta1Client
func (c *Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return c.rbacV1beta1
>>>>>>> upstream/master
}

// RbacV1alpha1 retrieves the RbacV1alpha1Client
func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.RbacV1alpha1Client
=======
	return c.rbacV1alpha1
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
}

// Deprecated: Scheduling retrieves the default version of SchedulingClient.
// Please explicitly pick a version.
func (c *Clientset) Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return c.schedulingV1alpha1
>>>>>>> upstream/master
}

// SettingsV1alpha1 retrieves the SettingsV1alpha1Client
func (c *Clientset) SettingsV1alpha1() settingsv1alpha1.SettingsV1alpha1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.SettingsV1alpha1Client
=======
	return c.settingsV1alpha1
>>>>>>> upstream/master
}

// Deprecated: Settings retrieves the default version of SettingsClient.
// Please explicitly pick a version.
func (c *Clientset) Settings() settingsv1alpha1.SettingsV1alpha1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.SettingsV1alpha1Client
=======
	return c.settingsV1alpha1
>>>>>>> upstream/master
}

// StorageV1beta1 retrieves the StorageV1beta1Client
func (c *Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.StorageV1beta1Client
=======
	return c.storageV1beta1
>>>>>>> upstream/master
}

// StorageV1 retrieves the StorageV1Client
func (c *Clientset) StorageV1() storagev1.StorageV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.StorageV1Client
=======
	return c.storageV1
>>>>>>> upstream/master
}

// Deprecated: Storage retrieves the default version of StorageClient.
// Please explicitly pick a version.
func (c *Clientset) Storage() storagev1.StorageV1Interface {
<<<<<<< HEAD
	if c == nil {
		return nil
	}
	return c.StorageV1Client
=======
	return c.storageV1
>>>>>>> upstream/master
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
<<<<<<< HEAD
	cs.AdmissionregistrationV1alpha1Client, err = admissionregistrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.CoreV1Client, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AppsV1beta1Client, err = appsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AuthenticationV1Client, err = authenticationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AuthenticationV1beta1Client, err = authenticationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AuthorizationV1Client, err = authorizationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AuthorizationV1beta1Client, err = authorizationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AutoscalingV1Client, err = autoscalingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AutoscalingV2alpha1Client, err = autoscalingv2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.BatchV1Client, err = batchv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.BatchV2alpha1Client, err = batchv2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.CertificatesV1beta1Client, err = certificatesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.ExtensionsV1beta1Client, err = extensionsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.NetworkingV1Client, err = networkingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.PolicyV1beta1Client, err = policyv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.RbacV1beta1Client, err = rbacv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.RbacV1alpha1Client, err = rbacv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.SettingsV1alpha1Client, err = settingsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.StorageV1beta1Client, err = storagev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.StorageV1Client, err = storagev1.NewForConfig(&configShallowCopy)
=======
	cs.admissionregistrationV1alpha1, err = admissionregistrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta1, err = appsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsV1beta2, err = appsv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1, err = authenticationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authenticationV1beta1, err = authenticationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1, err = authorizationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.authorizationV1beta1, err = authorizationv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1, err = autoscalingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV2beta1, err = autoscalingv2beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV1, err = batchv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV1beta1, err = batchv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.batchV2alpha1, err = batchv2alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certificatesV1beta1, err = certificatesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.extensionsV1beta1, err = extensionsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.policyV1beta1, err = policyv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1, err = rbacv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1beta1, err = rbacv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1alpha1, err = rbacv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schedulingV1alpha1, err = schedulingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.settingsV1alpha1, err = settingsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1beta1, err = storagev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1, err = storagev1.NewForConfig(&configShallowCopy)
>>>>>>> upstream/master
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
<<<<<<< HEAD
	cs.AdmissionregistrationV1alpha1Client = admissionregistrationv1alpha1.NewForConfigOrDie(c)
	cs.CoreV1Client = corev1.NewForConfigOrDie(c)
	cs.AppsV1beta1Client = appsv1beta1.NewForConfigOrDie(c)
	cs.AuthenticationV1Client = authenticationv1.NewForConfigOrDie(c)
	cs.AuthenticationV1beta1Client = authenticationv1beta1.NewForConfigOrDie(c)
	cs.AuthorizationV1Client = authorizationv1.NewForConfigOrDie(c)
	cs.AuthorizationV1beta1Client = authorizationv1beta1.NewForConfigOrDie(c)
	cs.AutoscalingV1Client = autoscalingv1.NewForConfigOrDie(c)
	cs.AutoscalingV2alpha1Client = autoscalingv2alpha1.NewForConfigOrDie(c)
	cs.BatchV1Client = batchv1.NewForConfigOrDie(c)
	cs.BatchV2alpha1Client = batchv2alpha1.NewForConfigOrDie(c)
	cs.CertificatesV1beta1Client = certificatesv1beta1.NewForConfigOrDie(c)
	cs.ExtensionsV1beta1Client = extensionsv1beta1.NewForConfigOrDie(c)
	cs.NetworkingV1Client = networkingv1.NewForConfigOrDie(c)
	cs.PolicyV1beta1Client = policyv1beta1.NewForConfigOrDie(c)
	cs.RbacV1beta1Client = rbacv1beta1.NewForConfigOrDie(c)
	cs.RbacV1alpha1Client = rbacv1alpha1.NewForConfigOrDie(c)
	cs.SettingsV1alpha1Client = settingsv1alpha1.NewForConfigOrDie(c)
	cs.StorageV1beta1Client = storagev1beta1.NewForConfigOrDie(c)
	cs.StorageV1Client = storagev1.NewForConfigOrDie(c)
=======
	cs.admissionregistrationV1alpha1 = admissionregistrationv1alpha1.NewForConfigOrDie(c)
	cs.appsV1beta1 = appsv1beta1.NewForConfigOrDie(c)
	cs.appsV1beta2 = appsv1beta2.NewForConfigOrDie(c)
	cs.authenticationV1 = authenticationv1.NewForConfigOrDie(c)
	cs.authenticationV1beta1 = authenticationv1beta1.NewForConfigOrDie(c)
	cs.authorizationV1 = authorizationv1.NewForConfigOrDie(c)
	cs.authorizationV1beta1 = authorizationv1beta1.NewForConfigOrDie(c)
	cs.autoscalingV1 = autoscalingv1.NewForConfigOrDie(c)
	cs.autoscalingV2beta1 = autoscalingv2beta1.NewForConfigOrDie(c)
	cs.batchV1 = batchv1.NewForConfigOrDie(c)
	cs.batchV1beta1 = batchv1beta1.NewForConfigOrDie(c)
	cs.batchV2alpha1 = batchv2alpha1.NewForConfigOrDie(c)
	cs.certificatesV1beta1 = certificatesv1beta1.NewForConfigOrDie(c)
	cs.coreV1 = corev1.NewForConfigOrDie(c)
	cs.extensionsV1beta1 = extensionsv1beta1.NewForConfigOrDie(c)
	cs.networkingV1 = networkingv1.NewForConfigOrDie(c)
	cs.policyV1beta1 = policyv1beta1.NewForConfigOrDie(c)
	cs.rbacV1 = rbacv1.NewForConfigOrDie(c)
	cs.rbacV1beta1 = rbacv1beta1.NewForConfigOrDie(c)
	cs.rbacV1alpha1 = rbacv1alpha1.NewForConfigOrDie(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.NewForConfigOrDie(c)
	cs.settingsV1alpha1 = settingsv1alpha1.NewForConfigOrDie(c)
	cs.storageV1beta1 = storagev1beta1.NewForConfigOrDie(c)
	cs.storageV1 = storagev1.NewForConfigOrDie(c)
>>>>>>> upstream/master

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
<<<<<<< HEAD
	cs.AdmissionregistrationV1alpha1Client = admissionregistrationv1alpha1.New(c)
	cs.CoreV1Client = corev1.New(c)
	cs.AppsV1beta1Client = appsv1beta1.New(c)
	cs.AuthenticationV1Client = authenticationv1.New(c)
	cs.AuthenticationV1beta1Client = authenticationv1beta1.New(c)
	cs.AuthorizationV1Client = authorizationv1.New(c)
	cs.AuthorizationV1beta1Client = authorizationv1beta1.New(c)
	cs.AutoscalingV1Client = autoscalingv1.New(c)
	cs.AutoscalingV2alpha1Client = autoscalingv2alpha1.New(c)
	cs.BatchV1Client = batchv1.New(c)
	cs.BatchV2alpha1Client = batchv2alpha1.New(c)
	cs.CertificatesV1beta1Client = certificatesv1beta1.New(c)
	cs.ExtensionsV1beta1Client = extensionsv1beta1.New(c)
	cs.NetworkingV1Client = networkingv1.New(c)
	cs.PolicyV1beta1Client = policyv1beta1.New(c)
	cs.RbacV1beta1Client = rbacv1beta1.New(c)
	cs.RbacV1alpha1Client = rbacv1alpha1.New(c)
	cs.SettingsV1alpha1Client = settingsv1alpha1.New(c)
	cs.StorageV1beta1Client = storagev1beta1.New(c)
	cs.StorageV1Client = storagev1.New(c)
=======
	cs.admissionregistrationV1alpha1 = admissionregistrationv1alpha1.New(c)
	cs.appsV1beta1 = appsv1beta1.New(c)
	cs.appsV1beta2 = appsv1beta2.New(c)
	cs.authenticationV1 = authenticationv1.New(c)
	cs.authenticationV1beta1 = authenticationv1beta1.New(c)
	cs.authorizationV1 = authorizationv1.New(c)
	cs.authorizationV1beta1 = authorizationv1beta1.New(c)
	cs.autoscalingV1 = autoscalingv1.New(c)
	cs.autoscalingV2beta1 = autoscalingv2beta1.New(c)
	cs.batchV1 = batchv1.New(c)
	cs.batchV1beta1 = batchv1beta1.New(c)
	cs.batchV2alpha1 = batchv2alpha1.New(c)
	cs.certificatesV1beta1 = certificatesv1beta1.New(c)
	cs.coreV1 = corev1.New(c)
	cs.extensionsV1beta1 = extensionsv1beta1.New(c)
	cs.networkingV1 = networkingv1.New(c)
	cs.policyV1beta1 = policyv1beta1.New(c)
	cs.rbacV1 = rbacv1.New(c)
	cs.rbacV1beta1 = rbacv1beta1.New(c)
	cs.rbacV1alpha1 = rbacv1alpha1.New(c)
	cs.schedulingV1alpha1 = schedulingv1alpha1.New(c)
	cs.settingsV1alpha1 = settingsv1alpha1.New(c)
	cs.storageV1beta1 = storagev1beta1.New(c)
	cs.storageV1 = storagev1.New(c)
>>>>>>> upstream/master

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
