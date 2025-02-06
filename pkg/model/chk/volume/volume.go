// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package volume

import (
	apps "k8s.io/api/apps/v1"

	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
	"github.com/altinity/clickhouse-operator/pkg/model/chk/config"
	"github.com/altinity/clickhouse-operator/pkg/model/chk/namer"
	"github.com/altinity/clickhouse-operator/pkg/model/k8s"
)

type Manager struct {
	cr    api.ICustomResource
	namer *namer.Namer
}

func NewManager() *Manager {
	return &Manager{
		namer: namer.New(),
	}
}

func (m *Manager) SetupVolumes(what interfaces.VolumeType, statefulSet *apps.StatefulSet, host *api.Host) {
	switch what {
	case interfaces.VolumesForConfigMaps:
		m.stsSetupVolumesForConfigMaps(statefulSet, host)
		return
	case interfaces.VolumesUserDataWithFixedPaths:
		m.stsSetupVolumesUserDataWithFixedPaths(statefulSet, host)
		return
	}
	panic("unknown volume type")
}

func (m *Manager) SetCR(cr api.ICustomResource) {
	m.cr = cr
}

// stsSetupVolumesForConfigMaps adds to each container in the Pod VolumeMount objects
func (m *Manager) stsSetupVolumesForConfigMaps(statefulSet *apps.StatefulSet, host *api.Host) {
	configMapCommonName := m.namer.Name(interfaces.NameConfigMapCommon, m.cr)
	configMapCommonUsersName := m.namer.Name(interfaces.NameConfigMapCommonUsers, m.cr)
	configMapHostName := m.namer.Name(interfaces.NameConfigMapHost, host)

	// Add all ConfigMap objects as Volume objects of type ConfigMap
	k8s.StatefulSetAppendVolumes(
		statefulSet,
		k8s.CreateVolumeForConfigMap(configMapCommonName),
		k8s.CreateVolumeForConfigMap(configMapCommonUsersName),
		k8s.CreateVolumeForConfigMap(configMapHostName),
	)

	// And reference these Volumes in each Container via VolumeMount
	// So Pod will have ConfigMaps mounted as Volumes in each Container
	k8s.StatefulSetAppendVolumeMountsInAllContainers(
		statefulSet,
		k8s.CreateVolumeMount(configMapCommonName, config.DirPathConfigCommon),
		k8s.CreateVolumeMount(configMapCommonUsersName, config.DirPathConfigUsers),
		k8s.CreateVolumeMount(configMapHostName, config.DirPathConfigHost),
	)
}

// stsSetupVolumesUserDataWithFixedPaths
// appends VolumeMounts for Data and Log VolumeClaimTemplates on all containers.
// Creates VolumeMounts for Data and Log volumes in case these volume templates are specified in `templates`.
func (m *Manager) stsSetupVolumesUserDataWithFixedPaths(statefulSet *apps.StatefulSet, host *api.Host) {
	// Mount all named (data and log so far) VolumeClaimTemplates into all containers
	k8s.StatefulSetAppendVolumeMountsInAllContainers(
		statefulSet,
		k8s.CreateVolumeMount(host.GetTemplates().GetDataVolumeClaimTemplate(), config.DirPathDataStorage),
		k8s.CreateVolumeMount(host.GetTemplates().GetLogVolumeClaimTemplate(), config.DirPathLogStorage),
	)
}
