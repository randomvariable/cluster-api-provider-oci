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

package scope

import (
	"os"
	"testing"

	infrastructurev1beta1 "github.com/oracle/cluster-api-provider-oci/api/v1beta1"
	infrav1exp "github.com/oracle/cluster-api-provider-oci/exp/api/v1beta1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func TestMain(m *testing.M) {
	code := 0
	defer func() { os.Exit(code) }()
	setup()
	code = m.Run()
}

func setup() {
	utilruntime.Must(infrastructurev1beta1.AddToScheme(scheme.Scheme))
	utilruntime.Must(clusterv1.AddToScheme(scheme.Scheme))
	utilruntime.Must(infrav1exp.AddToScheme(scheme.Scheme))
}
