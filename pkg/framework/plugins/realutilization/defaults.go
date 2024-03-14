/*
Copyright 2022 The Kubernetes Authors.
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

package realutilization

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_LowNodeRealUtilizationArgs
// TODO: the final default values would be discussed in community
func SetDefaults_LowNodeRealUtilizationArgs(obj runtime.Object) {
	args := obj.(*LowNodeRealUtilizationArgs)
	if args.Thresholds == nil {
		args.Thresholds = nil
	}
	if args.TargetThresholds == nil {
		args.TargetThresholds = nil
	}
	if args.NumberOfNodes == 0 {
		args.NumberOfNodes = 0
	}
}
