/*
Copyright (c) 2021, 2022 Oracle and/or its affiliates.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var logger = ctrl.Log.WithName("ocimachinepool-resource")

var (
	_ webhook.Defaulter = &OCIManagedMachinePool{}
	_ webhook.Validator = &OCIManagedMachinePool{}
)

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1beta1-ocimanagedmachinepool,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=ocimanagedmachinepools,versions=v1beta1,name=validation.ocimanagedmachinepool.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1beta1-ocimanagedmachinepool,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=ocimanagedmachinepools,versions=v1beta1,name=default.ocimanagedmachinepool.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1;v1beta1

func (m *OCIManagedMachinePool) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(m).
		Complete()
}

func (m *OCIManagedMachinePool) Default() {
	if m.Spec.NodePoolNodeConfig == nil {
		m.Spec.NodePoolNodeConfig = &NodePoolNodeConfig{}
	}
	if m.Spec.NodePoolNodeConfig.NodePoolPodNetworkOptionDetails == nil {
		m.Spec.NodePoolNodeConfig.NodePoolPodNetworkOptionDetails = &NodePoolPodNetworkOptionDetails{
			CniType: VCNNativeCNI,
			VcnIpNativePodNetworkOptions: VcnIpNativePodNetworkOptions{
				SubnetNames: []string{PodDefaultName},
				NSGNames:    []string{PodDefaultName},
			},
		}
	}
}

func (m *OCIManagedMachinePool) ValidateCreate() error {
	return nil
}

func (m *OCIManagedMachinePool) ValidateUpdate(old runtime.Object) error {
	return nil
}

func (m *OCIManagedMachinePool) ValidateDelete() error {
	return nil
}
