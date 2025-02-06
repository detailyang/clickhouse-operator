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

package normalizer

import (
	log "github.com/altinity/clickhouse-operator/pkg/announcer"
	chi "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/apis/common/types"
	"github.com/altinity/clickhouse-operator/pkg/apis/deployment"
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
	"github.com/altinity/clickhouse-operator/pkg/model/common/creator"
	"github.com/altinity/clickhouse-operator/pkg/model/common/namer"
)

func (n *Normalizer) hostApplyHostTemplateSpecifiedOrDefault(host *chi.Host) {
	hostTemplate := n.hostGetHostTemplate(host)
	hostApplyHostTemplate(host, hostTemplate)
}

// hostGetHostTemplate gets Host Template to be used to normalize Host
func (n *Normalizer) hostGetHostTemplate(host *chi.Host) *chi.HostTemplate {
	// Which host template would be used - either explicitly defined in or a default one
	if hostTemplate, ok := host.GetHostTemplate(); ok {
		// Host explicitly references known HostTemplate
		log.V(2).M(host).F().Info("host: %s uses custom hostTemplate: %s", host.Name, hostTemplate.Name)
		return hostTemplate
	}

	// Host references either no template or an UNKNOWN HostTemplate, thus will use a default host template.
	// However, with the default host template there is a nuance - hostNetwork requires different default host template.

	// Check hostNetwork case at first
	if podTemplate, ok := host.GetPodTemplate(); ok {
		if podTemplate.Spec.HostNetwork {
			// HostNetwork
			log.V(3).M(host).F().Info("host: %s uses default hostTemplate for HostNetwork", host.Name)
			return creator.CreateHostTemplate(interfaces.HostTemplateHostNetwork, n.namer.Name(interfaces.NameHostTemplate, host))
		}
	}

	// Pick default host template
	log.V(3).M(host).F().Info("host: %s uses default hostTemplate", host.Name)
	return creator.CreateHostTemplate(interfaces.HostTemplateCommon, n.namer.Name(interfaces.NameHostTemplate, host))
}

// hostApplyHostTemplate
func hostApplyHostTemplate(host *chi.Host, template *chi.HostTemplate) {
	if host.GetName() == "" {
		host.Name = template.Spec.Name
		log.V(3).M(host).F().Info("host has no name specified thus assigning name from Spec: %s", host.GetName())
	}

	host.Insecure = host.Insecure.MergeFrom(template.Spec.Insecure)
	host.Secure = host.Secure.MergeFrom(template.Spec.Secure)

	hostApplyHostTemplatePortDistribution(host, template)
	hostApplyPortsFromSettings(host)

	host.InheritTemplatesFrom(template)
}

func hostApplyHostTemplatePortDistribution(host *chi.Host, template *chi.HostTemplate) {
	for _, portDistribution := range template.PortDistribution {
		switch portDistribution.Type {
		case deployment.PortDistributionUnspecified:
			if !host.TCPPort.HasValue() {
				host.TCPPort = template.Spec.TCPPort
			}
			if !host.TLSPort.HasValue() {
				host.TLSPort = template.Spec.TLSPort
			}
			if !host.HTTPPort.HasValue() {
				host.HTTPPort = template.Spec.HTTPPort
			}
			if !host.HTTPSPort.HasValue() {
				host.HTTPSPort = template.Spec.HTTPSPort
			}
			if !host.InterserverHTTPPort.HasValue() {
				host.InterserverHTTPPort = template.Spec.InterserverHTTPPort
			}
		case deployment.PortDistributionClusterScopeIndex:
			if !host.TCPPort.HasValue() {
				base := chi.ChDefaultTCPPortNumber
				if template.Spec.TCPPort.HasValue() {
					base = template.Spec.TCPPort.Value()
				}
				host.TCPPort = types.NewInt32(base + int32(host.Runtime.Address.ClusterScopeIndex))
			}
			if !host.TLSPort.HasValue() {
				base := chi.ChDefaultTLSPortNumber
				if template.Spec.TLSPort.HasValue() {
					base = template.Spec.TLSPort.Value()
				}
				host.TLSPort = types.NewInt32(base + int32(host.Runtime.Address.ClusterScopeIndex))
			}
			if !host.HTTPPort.HasValue() {
				base := chi.ChDefaultHTTPPortNumber
				if template.Spec.HTTPPort.HasValue() {
					base = template.Spec.HTTPPort.Value()
				}
				host.HTTPPort = types.NewInt32(base + int32(host.Runtime.Address.ClusterScopeIndex))
			}
			if !host.HTTPSPort.HasValue() {
				base := chi.ChDefaultHTTPSPortNumber
				if template.Spec.HTTPSPort.HasValue() {
					base = template.Spec.HTTPSPort.Value()
				}
				host.HTTPSPort = types.NewInt32(base + int32(host.Runtime.Address.ClusterScopeIndex))
			}
			if !host.InterserverHTTPPort.HasValue() {
				base := chi.ChDefaultInterserverHTTPPortNumber
				if template.Spec.InterserverHTTPPort.HasValue() {
					base = template.Spec.InterserverHTTPPort.Value()
				}
				host.InterserverHTTPPort = types.NewInt32(base + int32(host.Runtime.Address.ClusterScopeIndex))
			}
		}
	}
}

// hostApplyPortsFromSettings
func hostApplyPortsFromSettings(host *chi.Host) {
	// Use host personal settings at first
	hostEnsurePortValuesFromSettings(host, host.GetSettings(), false)
	// Fallback to common settings
	hostEnsurePortValuesFromSettings(host, host.GetCR().GetSpec().GetConfiguration().GetSettings(), true)
}

// hostEnsurePortValuesFromSettings fetches port spec from settings, if any provided
func hostEnsurePortValuesFromSettings(host *chi.Host, settings *chi.Settings, final bool) {
	//
	// 1. Setup fallback/default ports
	//
	// For intermittent (non-final) setup fallback values should be from "MustBeAssignedLater" family,
	// because this is not final setup (just intermittent) and all these ports may be overwritten later
	var (
		fallbackTCPPort             *types.Int32
		fallbackTLSPort             *types.Int32
		fallbackHTTPPort            *types.Int32
		fallbackHTTPSPort           *types.Int32
		fallbackInterserverHTTPPort *types.Int32
	)

	// On the other hand, for final setup we need to assign real numbers to ports
	if final {
		if host.IsInsecure() {
			fallbackTCPPort = types.NewInt32(chi.ChDefaultTCPPortNumber)
			fallbackHTTPPort = types.NewInt32(chi.ChDefaultHTTPPortNumber)
		}
		if host.IsSecure() {
			fallbackTLSPort = types.NewInt32(chi.ChDefaultTLSPortNumber)
			fallbackHTTPSPort = types.NewInt32(chi.ChDefaultHTTPSPortNumber)
		}
		fallbackInterserverHTTPPort = types.NewInt32(chi.ChDefaultInterserverHTTPPortNumber)
	}

	//
	// 2. Setup ports
	//
	host.TCPPort = types.EnsurePortValue(host.TCPPort, settings.GetTCPPort(), fallbackTCPPort)
	host.TLSPort = types.EnsurePortValue(host.TLSPort, settings.GetTCPPortSecure(), fallbackTLSPort)
	host.HTTPPort = types.EnsurePortValue(host.HTTPPort, settings.GetHTTPPort(), fallbackHTTPPort)
	host.HTTPSPort = types.EnsurePortValue(host.HTTPSPort, settings.GetHTTPSPort(), fallbackHTTPSPort)
	host.InterserverHTTPPort = types.EnsurePortValue(host.InterserverHTTPPort, settings.GetInterserverHTTPPort(), fallbackInterserverHTTPPort)
}

// createHostsField
func createHostsField(cluster *chi.Cluster) {
	// Create HostsField of required size
	cluster.Layout.HostsField = chi.NewHostsField(cluster.Layout.ShardsCount, cluster.Layout.ReplicasCount)

	//
	// Migrate hosts from Shards and Replicas into HostsField.
	// Hosts which are explicitly specified in Shards and Replicas are migrated into HostsField for further use
	//
	hostMigrationFunc := func(shard, replica int, host *chi.Host) error {
		if curHost := cluster.Layout.HostsField.Get(shard, replica); curHost == nil {
			cluster.Layout.HostsField.Set(shard, replica, host)
		} else {
			curHost.MergeFrom(host)
		}
		return nil
	}

	// Run host migration func on all hosts specified in shards and replicas - migrate specified hosts into hosts field
	cluster.WalkHostsByShards(hostMigrationFunc)
	cluster.WalkHostsByReplicas(hostMigrationFunc)
}

// normalizeHost normalizes a host
func (n *Normalizer) normalizeHost(
	host *chi.Host,
	shard chi.IShard,
	replica chi.IReplica,
	cluster chi.ICluster,
	shardIndex int,
	replicaIndex int,
) {

	n.normalizeHostName(host, shard, shardIndex, replica, replicaIndex)
	// Inherit from either Shard or Replica
	var s chi.IShard
	var r chi.IReplica
	if cluster.IsShardSpecified() {
		s = shard
	} else {
		r = replica
	}
	host.InheritSettingsFrom(s, r)
	host.Settings = n.normalizeConfigurationSettings(host.Settings)
	host.InheritFilesFrom(s, r)
	host.Files = n.normalizeConfigurationFiles(host.Files)
	host.InheritTemplatesFrom(s, r)

	n.normalizeHostEnvVars()
}

func (n *Normalizer) normalizeHostEnvVars() {
}

// normalizeHostName normalizes host's name
func (n *Normalizer) normalizeHostName(
	host *chi.Host,
	shard chi.IShard,
	shardIndex int,
	replica chi.IReplica,
	replicaIndex int,
) {
	hasHostName := len(host.GetName()) > 0
	explicitlySpecifiedHostName := !namer.IsAutoGeneratedHostName(host.GetName(), host, shard, shardIndex, replica, replicaIndex)
	if hasHostName && explicitlySpecifiedHostName {
		// Has explicitly specified name already, normalization is not required
		return
	}

	// Create host name
	host.Name = n.namer.Name(interfaces.NameHost, host, shard, shardIndex, replica, replicaIndex)
}
