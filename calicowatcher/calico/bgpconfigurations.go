/*
Copyright 2021. Netris, Inc.

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

package calico

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

// BGPConfiguration .
type BGPConfiguration struct {
	Name            string `json:"name"`
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	Metadata metav1.ObjectMeta `json:"metadata"`

	// Specification of the BGPConfiguration.
	Spec BGPConfigurationSpec `json:"spec,omitempty"`
}

// BGPConfigurationSpec contains the values of the BGP configuration.
type BGPConfigurationSpec struct {
	// LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]
	// LogSeverityScreen string `json:"logSeverityScreen,omitempty" validate:"omitempty,logLevel" confignamev1:"loglevel"`

	// NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]
	NodeToNodeMeshEnabled *bool `json:"nodeToNodeMeshEnabled,omitempty" validate:"omitempty" confignamev1:"node_mesh"`

	// ASNumber is the default AS number used by a node. [Default: 64512]
	ASNumber int `json:"asNumber,omitempty" validate:"omitempty" confignamev1:"as_num"`

	// ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs.
	// Kubernetes Service ExternalIPs will only be advertised if they are within one of these blocks.
	ServiceExternalIPs []ServiceExternalIPBlock `json:"serviceExternalIPs,omitempty" validate:"omitempty,dive" confignamev1:"svc_external_ips"`

	// ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated.
	// If specified, Calico will advertise these blocks, as well as any cluster IPs within them.
	ServiceClusterIPs []ServiceClusterIPBlock `json:"serviceClusterIPs,omitempty" validate:"omitempty,dive" confignamev1:"svc_cluster_ips"`
}

// ServiceExternalIPBlock represents a single whitelisted CIDR External IP block.
type ServiceExternalIPBlock struct {
	CIDR string `json:"cidr,omitempty" validate:"omitempty,net"`
}

// ServiceClusterIPBlock represents a single whitelisted CIDR block for ClusterIPs.
type ServiceClusterIPBlock struct {
	CIDR string `json:"cidr,omitempty" validate:"omitempty,net"`
}

// IsMissingResource error message parser for missing calico case.
func IsMissingResource(err error) bool {
	return err.Error() == "the server could not find the requested resource"
}

// GetBGPConfiguration .
func (c *Calico) GetBGPConfiguration(config *rest.Config) ([]*BGPConfiguration, error) {
	ctx, cancel := context.WithTimeout(cntxt, contextTimeout)
	defer cancel()
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("{getBGPConfiguration} %s", err)
	}

	bgpConfigurationResource := schema.GroupVersionResource{
		Group:    "crd.projectcalico.org",
		Version:  "v1",
		Resource: "bgpconfigurations",
	}

	list, err := dynClient.Resource(bgpConfigurationResource).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var bgpConfigurations []*BGPConfiguration
	for _, peer := range list.Items {
		js, err := peer.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("{getBGPConfiguration} %s", err)
		}

		var bgpConfiguration *BGPConfiguration
		err = json.Unmarshal(js, &bgpConfiguration)
		if err != nil {
			return nil, fmt.Errorf("{getBGPConfiguration} %s", err)
		}
		bgpConfiguration.Name = bgpConfiguration.Metadata.Name
		bgpConfigurations = append(bgpConfigurations, bgpConfiguration)
	}
	return bgpConfigurations, nil
}

// UpdateBGPConfiguration .
func (c *Calico) UpdateBGPConfiguration(bgpConf *BGPConfiguration, config *rest.Config) error {
	ctx, cancel := context.WithTimeout(cntxt, contextTimeout)
	defer cancel()
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("{UpdateBGPConfiguration} %s", err)
	}

	bgpPeerResource := schema.GroupVersionResource{
		Group:    "crd.projectcalico.org",
		Version:  "v1",
		Resource: "bgpconfigurations",
	}

	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(bgpConf)
	if err != nil {
		return fmt.Errorf("{UpdateBGPConfiguration} %s", err)
	}

	obj := &unstructured.Unstructured{
		Object: m,
	}

	_, err = dynClient.Resource(bgpPeerResource).Update(ctx, obj, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("{UpdateBGPConfiguration} %s", err)
	}
	return nil
}
