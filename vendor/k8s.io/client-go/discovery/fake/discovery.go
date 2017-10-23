/*
Copyright 2016 The Kubernetes Authors.

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

package fake

import (
	"fmt"

	"github.com/emicklei/go-restful-swagger12"
<<<<<<< HEAD

	"github.com/go-openapi/spec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/pkg/api/v1"
=======
	"github.com/googleapis/gnostic/OpenAPIv2"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
>>>>>>> upstream/master
	kubeversion "k8s.io/client-go/pkg/version"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/testing"
)

<<<<<<< HEAD
type FakeDiscovery struct {
	*testing.Fake
=======
// FakeDiscovery implements discovery.DiscoveryInterface and sometimes calls testing.Fake.Invoke with an action,
// but doesn't respect the return value if any. There is a way to fake static values like ServerVersion by using the Faked... fields on the struct.
type FakeDiscovery struct {
	*testing.Fake
	FakedServerVersion *version.Info
>>>>>>> upstream/master
}

func (c *FakeDiscovery) ServerResourcesForGroupVersion(groupVersion string) (*metav1.APIResourceList, error) {
	action := testing.ActionImpl{
		Verb:     "get",
		Resource: schema.GroupVersionResource{Resource: "resource"},
	}
	c.Invokes(action, nil)
	for _, resourceList := range c.Resources {
		if resourceList.GroupVersion == groupVersion {
			return resourceList, nil
		}
	}
	return nil, fmt.Errorf("GroupVersion %q not found", groupVersion)
}

func (c *FakeDiscovery) ServerResources() ([]*metav1.APIResourceList, error) {
	action := testing.ActionImpl{
		Verb:     "get",
		Resource: schema.GroupVersionResource{Resource: "resource"},
	}
	c.Invokes(action, nil)
	return c.Resources, nil
}

func (c *FakeDiscovery) ServerPreferredResources() ([]*metav1.APIResourceList, error) {
	return nil, nil
}

func (c *FakeDiscovery) ServerPreferredNamespacedResources() ([]*metav1.APIResourceList, error) {
	return nil, nil
}

func (c *FakeDiscovery) ServerGroups() (*metav1.APIGroupList, error) {
	return nil, nil
}

func (c *FakeDiscovery) ServerVersion() (*version.Info, error) {
	action := testing.ActionImpl{}
	action.Verb = "get"
	action.Resource = schema.GroupVersionResource{Resource: "version"}
<<<<<<< HEAD

	c.Invokes(action, nil)
=======
	c.Invokes(action, nil)

	if c.FakedServerVersion != nil {
		return c.FakedServerVersion, nil
	}

>>>>>>> upstream/master
	versionInfo := kubeversion.Get()
	return &versionInfo, nil
}

func (c *FakeDiscovery) SwaggerSchema(version schema.GroupVersion) (*swagger.ApiDeclaration, error) {
	action := testing.ActionImpl{}
	action.Verb = "get"
	if version == v1.SchemeGroupVersion {
		action.Resource = schema.GroupVersionResource{Resource: "/swaggerapi/api/" + version.Version}
	} else {
		action.Resource = schema.GroupVersionResource{Resource: "/swaggerapi/apis/" + version.Group + "/" + version.Version}
	}

	c.Invokes(action, nil)
	return &swagger.ApiDeclaration{}, nil
}

<<<<<<< HEAD
func (c *FakeDiscovery) OpenAPISchema() (*spec.Swagger, error) { return &spec.Swagger{}, nil }
=======
func (c *FakeDiscovery) OpenAPISchema() (*openapi_v2.Document, error) {
	return &openapi_v2.Document{}, nil
}
>>>>>>> upstream/master

func (c *FakeDiscovery) RESTClient() restclient.Interface {
	return nil
}
