/*
Copyright 2024 chiragkyal.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// MypodSpec defines the desired state of Mypod
type MypodSpec struct {

	MinReplicas int `json:"minReplicas"`

	Replicas int `json:"replicas"`
	
	MaxReplicas int `json:"maxReplicas"`
}

// MypodStatus defines the observed state of Mypod
type MypodStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Mypod is the Schema for the mypods API
type Mypod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MypodSpec   `json:"spec,omitempty"`
	Status MypodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MypodList contains a list of Mypod
type MypodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mypod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mypod{}, &MypodList{})
}
