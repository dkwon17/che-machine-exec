//
// Copyright (c) 2012-2018 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package exec_info

import (
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

// Component to creation new info execs on the kubernetes infrastructure.
type KubernetesInfoExecCreator struct {
	InfoExecCreator

	namespace string
	core      v1.CoreV1Interface
	config    *rest.Config
}

// Return new instance of the kubernetes info exec creator.
func NewKubernetesInfoExecCreator(
	namespace string,
	core v1.CoreV1Interface,
	config *rest.Config) *KubernetesInfoExecCreator {
	return &KubernetesInfoExecCreator{
		namespace: namespace,
		core:      core,
		config:    config,
	}
}

// Create new kubernetes info exec.
func (creator *KubernetesInfoExecCreator) CreateInfoExec(command []string, containerInfo map[string]string) InfoExec {
	containerName := containerInfo[ContainerName]
	podName := containerInfo[PodName]
	return NewKubernetesInfoExec(
		command,
		containerName,
		podName,
		creator.namespace,
		creator.core,
		creator.config,
	)
}
