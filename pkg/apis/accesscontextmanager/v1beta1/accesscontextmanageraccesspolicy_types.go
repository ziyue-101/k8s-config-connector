// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AccessContextManagerAccessPolicySpec struct {
	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	ResourceID string `json:"resourceID,omitempty"`
	/* Human readable title. Does not affect behavior. */
	Title string `json:"title,omitempty"`
}

type AccessContextManagerAccessPolicyStatus struct {
	/* Conditions represents the latest available observations of the
	   AccessContextManagerAccessPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time the AccessPolicy was created in UTC. */
	CreateTime string `json:"createTime,omitempty"`
	/* Resource name of the AccessPolicy. Format: {policy_id} */
	Name string `json:"name,omitempty"`
	/* Time the AccessPolicy was updated in UTC. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerAccessPolicy is the Schema for the accesscontextmanager API
// +k8s:openapi-gen=true
type AccessContextManagerAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerAccessPolicySpec   `json:"spec,omitempty"`
	Status AccessContextManagerAccessPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerAccessPolicyList contains a list of AccessContextManagerAccessPolicy
type AccessContextManagerAccessPolicyList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []AccessContextManagerAccessPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerAccessPolicy{}, &AccessContextManagerAccessPolicyList{})
}