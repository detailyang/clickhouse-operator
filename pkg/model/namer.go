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

package model

import (
	"fmt"
	"strconv"
	"strings"

	apps "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"

	chop "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/util"
)

const (
	// Names context length
	namePartChiMaxLenNamesCtx     = 60
	namePartClusterMaxLenNamesCtx = 15
	namePartShardMaxLenNamesCtx   = 15
	namePartReplicaMaxLenNamesCtx = 15

	// Labels context length
	namePartChiMaxLenLabelsCtx     = 63
	namePartClusterMaxLenLabelsCtx = 63
	namePartShardMaxLenLabelsCtx   = 63
	namePartReplicaMaxLenLabelsCtx = 63
)

const (
	// macrosNamespace is a sanitized namespace name where ClickHouseInstallation runs
	macrosNamespace = "{namespace}"

	// macrosChiName is a sanitized ClickHouseInstallation name
	macrosChiName = "{chi}"
	// macrosChiID is a sanitized ID made of original ClickHouseInstallation name
	macrosChiID = "{chiID}"

	// macrosClusterName is a sanitized cluster name
	macrosClusterName = "{cluster}"
	// macrosClusterID is a sanitized ID made of original cluster name
	macrosClusterID = "{clusterID}"
	// macrosClusterIndex is an index of the cluster in the CHI - integer number, converted into string
	macrosClusterIndex = "{clusterIndex}"

	// macrosShardName is a sanitized shard name
	macrosShardName = "{shard}"
	// macrosShardID is a sanitized ID made of original shard name
	macrosShardID = "{shardID}"
	// macrosShardIndex is an index of the shard in the cluster - integer number, converted into string
	macrosShardIndex = "{shardIndex}"

	// macrosReplicaName is a sanitized replica name
	macrosReplicaName = "{replica}"
	// macrosReplicaID is a sanitized ID made of original replica name
	macrosReplicaID = "{replicaID}"
	// macrosReplicaIndex is an index of the replica in the cluster - integer number, converted into string
	macrosReplicaIndex = "{replicaIndex}"

	// macrosHostName is a sanitized host name
	macrosHostName = "{host}"
	// macrosHostID is a sanitized ID made of original host name
	macrosHostID = "{hostID}"
	// macrosChiScopeIndex is an index of the host on the CHI-scope
	macrosChiScopeIndex = "{chiScopeIndex}"
	// macrosChiScopeCycleIndex is an index of the host in the CHI-scope cycle - integer number, converted into string
	macrosChiScopeCycleIndex = "{chiScopeCycleIndex}"
	// macrosChiScopeCycleOffset is an offset of the host in the CHI-scope cycle - integer number, converted into string
	macrosChiScopeCycleOffset = "{chiScopeCycleOffset}"
	// macrosClusterScopeIndex is an index of the host on the cluster-scope
	macrosClusterScopeIndex = "{clusterScopeIndex}"
	// macrosClusterScopeCycleIndex is an index of the host in the Cluster-scope cycle - integer number, converted into string
	macrosClusterScopeCycleIndex = "{clusterScopeCycleIndex}"
	// macrosClusterScopeCycleOffset is an offset of the host in the Cluster-scope cycle - integer number, converted into string
	macrosClusterScopeCycleOffset = "{clusterScopeCycleOffset}"
	// macrosShardScopeIndex is an index of the host on the shard-scope
	macrosShardScopeIndex = "{shardScopeIndex}"
	// macrosReplicaScopeIndex is an index of the host on the replica-scope
	macrosReplicaScopeIndex = "{replicaScopeIndex}"
	// macrosClusterScopeCycleHeadPointsToPreviousCycleTail is {clusterScopeIndex} of previous Cycle Tail
	macrosClusterScopeCycleHeadPointsToPreviousCycleTail = "{clusterScopeCycleHeadPointsToPreviousCycleTail}"
)

const (
	// chiServiceNamePattern is a template of CHI Service name. "clickhouse-{chi}"
	chiServiceNamePattern = "clickhouse-" + macrosChiName

	// clusterServiceNamePattern is a template of cluster Service name. "cluster-{chi}-{cluster}"
	clusterServiceNamePattern = "cluster-" + macrosChiName + "-" + macrosClusterName

	// shardServiceNamePattern is a template of shard Service name. "shard-{chi}-{cluster}-{shard}"
	shardServiceNamePattern = "shard-" + macrosChiName + "-" + macrosClusterName + "-" + macrosShardName

	// replicaServiceNamePattern is a template of replica Service name. "shard-{chi}-{cluster}-{replica}"
	replicaServiceNamePattern = "shard-" + macrosChiName + "-" + macrosClusterName + "-" + macrosReplicaName

	// statefulSetNamePattern is a template of hosts's StatefulSet's name. "chi-{chi}-{cluster}-{shard}-{host}"
	statefulSetNamePattern = "chi-" + macrosChiName + "-" + macrosClusterName + "-" + macrosHostName

	// statefulSetServiceNamePattern is a template of hosts's StatefulSet's Service name. "chi-{chi}-{cluster}-{shard}-{host}"
	statefulSetServiceNamePattern = "chi-" + macrosChiName + "-" + macrosClusterName + "-" + macrosHostName

	// configMapCommonNamePattern is a template of common settings for the CHI ConfigMap. "chi-{chi}-common-configd"
	configMapCommonNamePattern = "chi-" + macrosChiName + "-common-configd"

	// configMapCommonUsersNamePattern is a template of common users settings for the CHI ConfigMap. "chi-{chi}-common-usersd"
	configMapCommonUsersNamePattern = "chi-" + macrosChiName + "-common-usersd"

	// configMapDeploymentNamePattern is a template of macros ConfigMap. "chi-{chi}-deploy-confd-{cluster}-{shard}-{host}"
	configMapDeploymentNamePattern = "chi-" + macrosChiName + "-deploy-confd-" + macrosClusterName + "-" + macrosHostName

	// namespaceDomainPattern presents Domain Name pattern of a namespace
	// In this pattern "%s" is substituted namespace name's value
	// Ex.: my-dev-namespace.svc.cluster.local
	namespaceDomainPattern = "%s.svc.cluster.local"

	// ServiceName.domain.name
	serviceFQDNPattern = "%s" + "." + namespaceDomainPattern

	// podFQDNPattern consists of 3 parts:
	// 1. nameless service of of stateful set
	// 2. namespace name
	// Hostname.domain.name
	podFQDNPattern = "%s" + "." + namespaceDomainPattern

	// podNamePattern is a name of a Pod within StatefulSet. In our setup each StatefulSet has only 1 pod,
	// so all pods would have '-0' suffix after StatefulSet name
	// Ex.: StatefulSetName-0
	podNamePattern = "%s-0"
)

// sanitize makes string fulfil kubernetes naming restrictions
// String can't end with '-', '_' and '.'
func sanitize(s string) string {
	return strings.Trim(s, "-_.")
}

const (
	namerContextLabels = "labels"
	namerContextNames  = "names"
)

type namerContext string
type namer struct {
	ctx namerContext
}

// newNamer
func newNamer(ctx namerContext) *namer {
	return &namer{
		ctx: ctx,
	}
}

// namePartNamespace
func (n *namer) namePartNamespace(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartChiMaxLenLabelsCtx
	} else {
		_len = namePartChiMaxLenNamesCtx
	}
	return sanitize(util.StringHead(name, _len))
}

// namePartChiName
func (n *namer) namePartChiName(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartChiMaxLenLabelsCtx
	} else {
		_len = namePartChiMaxLenNamesCtx
	}
	return sanitize(util.StringHead(name, _len))
}

// namePartChiNameID
func (n *namer) namePartChiNameID(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartChiMaxLenLabelsCtx
	} else {
		_len = namePartChiMaxLenNamesCtx
	}
	return util.CreateStringID(name, _len)
}

// namePartClusterName
func (n *namer) namePartClusterName(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartClusterMaxLenLabelsCtx
	} else {
		_len = namePartClusterMaxLenNamesCtx
	}
	return sanitize(util.StringHead(name, _len))
}

// namePartClusterNameID
func (n *namer) namePartClusterNameID(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartClusterMaxLenLabelsCtx
	} else {
		_len = namePartClusterMaxLenNamesCtx
	}
	return util.CreateStringID(name, _len)
}

// namePartShardName
func (n *namer) namePartShardName(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartShardMaxLenLabelsCtx
	} else {
		_len = namePartShardMaxLenNamesCtx
	}
	return sanitize(util.StringHead(name, _len))
}

// namePartShardNameID
func (n *namer) namePartShardNameID(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartShardMaxLenLabelsCtx
	} else {
		_len = namePartShardMaxLenNamesCtx
	}
	return util.CreateStringID(name, _len)
}

// namePartReplicaName
func (n *namer) namePartReplicaName(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartReplicaMaxLenLabelsCtx
	} else {
		_len = namePartReplicaMaxLenNamesCtx
	}
	return sanitize(util.StringHead(name, _len))
}

// namePartReplicaNameID
func (n *namer) namePartReplicaNameID(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartReplicaMaxLenLabelsCtx
	} else {
		_len = namePartReplicaMaxLenNamesCtx
	}
	return util.CreateStringID(name, _len)
}

// namePartHostName
func (n *namer) namePartHostName(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartReplicaMaxLenLabelsCtx
	} else {
		_len = namePartReplicaMaxLenNamesCtx
	}
	return sanitize(util.StringHead(name, _len))
}

// namePartHostNameID
func (n *namer) namePartHostNameID(name string) string {
	var _len int
	if n.ctx == namerContextLabels {
		_len = namePartReplicaMaxLenLabelsCtx
	} else {
		_len = namePartReplicaMaxLenNamesCtx
	}
	return util.CreateStringID(name, _len)
}

// getNamePartNamespace
func (n *namer) getNamePartNamespace(obj interface{}) string {
	switch obj.(type) {
	case *chop.ClickHouseInstallation:
		chi := obj.(*chop.ClickHouseInstallation)
		return n.namePartChiName(chi.Namespace)
	case *chop.ChiCluster:
		cluster := obj.(*chop.ChiCluster)
		return n.namePartChiName(cluster.Address.Namespace)
	case *chop.ChiShard:
		shard := obj.(*chop.ChiShard)
		return n.namePartChiName(shard.Address.Namespace)
	case *chop.ChiHost:
		host := obj.(*chop.ChiHost)
		return n.namePartChiName(host.Address.Namespace)
	}

	return "ERROR"
}

// getNamePartCHIName
func (n *namer) getNamePartCHIName(obj interface{}) string {
	switch obj.(type) {
	case *chop.ClickHouseInstallation:
		chi := obj.(*chop.ClickHouseInstallation)
		return n.namePartChiName(chi.Name)
	case *chop.ChiCluster:
		cluster := obj.(*chop.ChiCluster)
		return n.namePartChiName(cluster.Address.CHIName)
	case *chop.ChiShard:
		shard := obj.(*chop.ChiShard)
		return n.namePartChiName(shard.Address.CHIName)
	case *chop.ChiHost:
		host := obj.(*chop.ChiHost)
		return n.namePartChiName(host.Address.CHIName)
	}

	return "ERROR"
}

// getNamePartClusterName
func (n *namer) getNamePartClusterName(obj interface{}) string {
	switch obj.(type) {
	case *chop.ChiCluster:
		cluster := obj.(*chop.ChiCluster)
		return n.namePartClusterName(cluster.Address.ClusterName)
	case *chop.ChiShard:
		shard := obj.(*chop.ChiShard)
		return n.namePartClusterName(shard.Address.ClusterName)
	case *chop.ChiHost:
		host := obj.(*chop.ChiHost)
		return n.namePartClusterName(host.Address.ClusterName)
	}

	return "ERROR"
}

// getNamePartShardName
func (n *namer) getNamePartShardName(obj interface{}) string {
	switch obj.(type) {
	case *chop.ChiShard:
		shard := obj.(*chop.ChiShard)
		return n.namePartShardName(shard.Address.ShardName)
	case *chop.ChiHost:
		host := obj.(*chop.ChiHost)
		return n.namePartShardName(host.Address.ShardName)
	}

	return "ERROR"
}

// getNamePartReplicaName
func (n *namer) getNamePartReplicaName(host *chop.ChiHost) string {
	return n.namePartReplicaName(host.Address.ReplicaName)
}

// getNamePartHostName
func (n *namer) getNamePartHostName(host *chop.ChiHost) string {
	return n.namePartHostName(host.Address.HostName)
}

// getNamePartCHIScopeCycleSize
func getNamePartCHIScopeCycleSize(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.CHIScopeCycleSize)
}

// getNamePartCHIScopeCycleIndex
func getNamePartCHIScopeCycleIndex(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.CHIScopeCycleIndex)
}

// getNamePartCHIScopeCycleOffset
func getNamePartCHIScopeCycleOffset(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.CHIScopeCycleOffset)
}

// getNamePartClusterScopeCycleSize
func getNamePartClusterScopeCycleSize(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.ClusterScopeCycleSize)
}

// getNamePartClusterScopeCycleIndex
func getNamePartClusterScopeCycleIndex(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.ClusterScopeCycleIndex)
}

// getNamePartClusterScopeCycleOffset
func getNamePartClusterScopeCycleOffset(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.ClusterScopeCycleOffset)
}

// getNamePartCHIScopeIndex
func getNamePartCHIScopeIndex(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.CHIScopeIndex)
}

// getNamePartClusterScopeIndex
func getNamePartClusterScopeIndex(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.ClusterScopeIndex)
}

// getNamePartShardScopeIndex
func getNamePartShardScopeIndex(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.ShardScopeIndex)
}

// getNamePartReplicaScopeIndex
func getNamePartReplicaScopeIndex(host *chop.ChiHost) string {
	return strconv.Itoa(host.Address.ReplicaScopeIndex)
}

var namesNamer = newNamer(namerContextNames)

// newNameMacroReplacerChi
func newNameMacroReplacerChi(chi *chop.ClickHouseInstallation) *strings.Replacer {
	n := namesNamer
	return strings.NewReplacer(
		macrosNamespace, n.namePartNamespace(chi.Namespace),
		macrosChiName, n.namePartChiName(chi.Name),
		macrosChiID, n.namePartChiNameID(chi.Name),
	)
}

// newNameMacroReplacerCluster
func newNameMacroReplacerCluster(cluster *chop.ChiCluster) *strings.Replacer {
	n := namesNamer
	return strings.NewReplacer(
		macrosNamespace, n.namePartNamespace(cluster.Address.Namespace),
		macrosChiName, n.namePartChiName(cluster.Address.CHIName),
		macrosChiID, n.namePartChiNameID(cluster.Address.CHIName),
		macrosClusterName, n.namePartClusterName(cluster.Address.ClusterName),
		macrosClusterID, n.namePartClusterNameID(cluster.Address.ClusterName),
		macrosClusterIndex, strconv.Itoa(cluster.Address.ClusterIndex),
	)
}

// newNameMacroReplacerShard
func newNameMacroReplacerShard(shard *chop.ChiShard) *strings.Replacer {
	n := namesNamer
	return strings.NewReplacer(
		macrosNamespace, n.namePartNamespace(shard.Address.Namespace),
		macrosChiName, n.namePartChiName(shard.Address.CHIName),
		macrosChiID, n.namePartChiNameID(shard.Address.CHIName),
		macrosClusterName, n.namePartClusterName(shard.Address.ClusterName),
		macrosClusterID, n.namePartClusterNameID(shard.Address.ClusterName),
		macrosClusterIndex, strconv.Itoa(shard.Address.ClusterIndex),
		macrosShardName, n.namePartShardName(shard.Address.ShardName),
		macrosShardID, n.namePartShardNameID(shard.Address.ShardName),
		macrosShardIndex, strconv.Itoa(shard.Address.ShardIndex),
	)
}

// clusterScopeIndexOfPreviousCycleTail gets cluster-scope index of previous cycle tail
func clusterScopeIndexOfPreviousCycleTail(host *chop.ChiHost) int {
	if host.Address.ClusterScopeCycleOffset == 0 {
		// This is the cycle head - the first host of the cycle
		// We need to point to previous host in this cluster - which would be previous cycle tail

		if host.Address.ClusterScopeIndex == 0 {
			// This is the very first host in the cluster - head of the first cycle
			// No previous host available, so just point to the same host, mainly because label must be an empty string
			// or consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character
			// So we can't set it to "-1"
			return host.Address.ClusterScopeIndex
		}

		// This is head of non-first cycle, point to previous host in the cluster - which would be previous cycle tail
		return host.Address.ClusterScopeIndex - 1
	}

	// This is not cycle head - just point to the same host
	return host.Address.ClusterScopeIndex
}

// newNameMacroReplacerHost
func newNameMacroReplacerHost(host *chop.ChiHost) *strings.Replacer {
	n := namesNamer
	return strings.NewReplacer(
		macrosNamespace, n.namePartNamespace(host.Address.Namespace),
		macrosChiName, n.namePartChiName(host.Address.CHIName),
		macrosChiID, n.namePartChiNameID(host.Address.CHIName),
		macrosClusterName, n.namePartClusterName(host.Address.ClusterName),
		macrosClusterID, n.namePartClusterNameID(host.Address.ClusterName),
		macrosClusterIndex, strconv.Itoa(host.Address.ClusterIndex),
		macrosShardName, n.namePartShardName(host.Address.ShardName),
		macrosShardID, n.namePartShardNameID(host.Address.ShardName),
		macrosShardIndex, strconv.Itoa(host.Address.ShardIndex),
		macrosShardScopeIndex, strconv.Itoa(host.Address.ShardScopeIndex), // TODO use appropriate namePart function
		macrosReplicaName, n.namePartReplicaName(host.Address.ReplicaName),
		macrosReplicaID, n.namePartReplicaNameID(host.Address.ReplicaName),
		macrosReplicaIndex, strconv.Itoa(host.Address.ReplicaIndex),
		macrosReplicaScopeIndex, strconv.Itoa(host.Address.ReplicaScopeIndex), // TODO use appropriate namePart function
		macrosHostName, n.namePartHostName(host.Address.HostName),
		macrosHostID, n.namePartHostNameID(host.Address.HostName),
		macrosChiScopeIndex, strconv.Itoa(host.Address.CHIScopeIndex), // TODO use appropriate namePart function
		macrosChiScopeCycleIndex, strconv.Itoa(host.Address.CHIScopeCycleIndex), // TODO use appropriate namePart function
		macrosChiScopeCycleOffset, strconv.Itoa(host.Address.CHIScopeCycleOffset), // TODO use appropriate namePart function
		macrosClusterScopeIndex, strconv.Itoa(host.Address.ClusterScopeIndex), // TODO use appropriate namePart function
		macrosClusterScopeCycleIndex, strconv.Itoa(host.Address.ClusterScopeCycleIndex), // TODO use appropriate namePart function
		macrosClusterScopeCycleOffset, strconv.Itoa(host.Address.ClusterScopeCycleOffset), // TODO use appropriate namePart function
		macrosClusterScopeCycleHeadPointsToPreviousCycleTail, strconv.Itoa(clusterScopeIndexOfPreviousCycleTail(host)),
	)
}

// CreateConfigMapPersonalName returns a name for a ConfigMap for replica's personal config
func CreateConfigMapPersonalName(host *chop.ChiHost) string {
	return newNameMacroReplacerHost(host).Replace(configMapDeploymentNamePattern)
}

// CreateConfigMapCommonName returns a name for a ConfigMap for replica's common config
func CreateConfigMapCommonName(chi *chop.ClickHouseInstallation) string {
	return newNameMacroReplacerChi(chi).Replace(configMapCommonNamePattern)
}

// CreateConfigMapCommonUsersName returns a name for a ConfigMap for replica's common users config
func CreateConfigMapCommonUsersName(chi *chop.ClickHouseInstallation) string {
	return newNameMacroReplacerChi(chi).Replace(configMapCommonUsersNamePattern)
}

// CreateCHIServiceName creates a name of a root ClickHouseInstallation Service resource
func CreateCHIServiceName(chi *chop.ClickHouseInstallation) string {
	// Name can be generated either from default name pattern,
	// or from personal name pattern provided in ServiceTemplate

	// Start with default name pattern
	pattern := chiServiceNamePattern

	// ServiceTemplate may have personal name pattern specified
	if template, ok := chi.GetCHIServiceTemplate(); ok {
		// ServiceTemplate available
		if template.GenerateName != "" {
			// ServiceTemplate has explicitly specified name pattern
			pattern = template.GenerateName
		}
	}

	// Create Service name based on name pattern available
	return newNameMacroReplacerChi(chi).Replace(pattern)
}

// CreateCHIServiceFQDN creates a FQD name of a root ClickHouseInstallation Service resource
func CreateCHIServiceFQDN(chi *chop.ClickHouseInstallation) string {
	// FQDN can be generated either from default pattern,
	// or from personal pattern provided

	// Start with default pattern
	pattern := serviceFQDNPattern

	if chi.Spec.NamespaceDomainPattern != "" {
		// NamespaceDomainPattern has been explicitly specified
		pattern = "%s." + chi.Spec.NamespaceDomainPattern
	}

	// Create FQDN based on pattern available
	return fmt.Sprintf(
		pattern,
		CreateCHIServiceName(chi),
		chi.Namespace,
	)
}

// CreateClusterServiceName returns a name of a cluster's Service
func CreateClusterServiceName(cluster *chop.ChiCluster) string {
	// Name can be generated either from default name pattern,
	// or from personal name pattern provided in ServiceTemplate

	// Start with default name pattern
	pattern := clusterServiceNamePattern

	// ServiceTemplate may have personal name pattern specified
	if template, ok := cluster.GetServiceTemplate(); ok {
		// ServiceTemplate available
		if template.GenerateName != "" {
			// ServiceTemplate has explicitly specified name pattern
			pattern = template.GenerateName
		}
	}

	// Create Service name based on name pattern available
	return newNameMacroReplacerCluster(cluster).Replace(pattern)
}

// CreateShardServiceName returns a name of a shard's Service
func CreateShardServiceName(shard *chop.ChiShard) string {
	// Name can be generated either from default name pattern,
	// or from personal name pattern provided in ServiceTemplate

	// Start with default name pattern
	pattern := shardServiceNamePattern

	// ServiceTemplate may have personal name pattern specified
	if template, ok := shard.GetServiceTemplate(); ok {
		// ServiceTemplate available
		if template.GenerateName != "" {
			// ServiceTemplate has explicitly specified name pattern
			pattern = template.GenerateName
		}
	}

	// Create Service name based on name pattern available
	return newNameMacroReplacerShard(shard).Replace(pattern)
}

// CreateShardName return a name of a shard
func CreateShardName(shard *chop.ChiShard, index int) string {
	return strconv.Itoa(index)
}

// IsAutoGeneratedShardName checks whether provided name is auto-generated
func IsAutoGeneratedShardName(name string, shard *chop.ChiShard, index int) bool {
	return name == CreateShardName(shard, index)
}

// CreateReplicaName return a name of a replica.
// Here replica is a CHOp-internal replica - i.e. a vertical slice of hosts field.
// In case you are looking for replica name in terms of a hostname to address particular host as in remote_servers.xml
// you need to take a look on CreateReplicaHostname function
func CreateReplicaName(replica *chop.ChiReplica, index int) string {
	return strconv.Itoa(index)
}

// IsAutoGeneratedReplicaName checks whether provided name is auto-generated
func IsAutoGeneratedReplicaName(name string, replica *chop.ChiReplica, index int) bool {
	return name == CreateReplicaName(replica, index)
}

// CreateHostName return a name of a host
func CreateHostName(host *chop.ChiHost, shard *chop.ChiShard, shardIndex int, replica *chop.ChiReplica, replicaIndex int) string {
	return fmt.Sprintf("%s-%s", shard.Name, replica.Name)
}

// CreateReplicaHostname returns hostname (pod-hostname + service or FQDN) which can be used as a replica name
// in all places where ClickHouse requires replica name. These are such places as:
// 1. "remote_servers.xml" config file
// 2. statements like SYSTEM DROP REPLICA <replica_name>
// any other places
// Function operations are based on .Spec.Defaults.ReplicasUseFQDN
func CreateReplicaHostname(host *chop.ChiHost) string {
	if util.IsStringBoolTrue(host.GetCHI().Spec.Defaults.ReplicasUseFQDN) {
		// In case .Spec.Defaults.ReplicasUseFQDN is set replicas would use FQDN pod hostname,
		// otherwise hostname+service name (unique within namespace) would be used
		// .my-dev-namespace.svc.cluster.local
		return createPodFQDN(host)
	}

	return CreatePodHostname(host)
}

// IsAutoGeneratedHostName checks whether name is auto-generated
func IsAutoGeneratedHostName(
	name string,
	host *chop.ChiHost,
	shard *chop.ChiShard,
	shardIndex int,
	replica *chop.ChiReplica,
	replicaIndex int,
) bool {
	if name == CreateHostName(host, shard, shardIndex, replica, replicaIndex) {
		// Current version of the name
		return true
	}

	if name == fmt.Sprintf("%d-%d", shardIndex, replicaIndex) {
		// old version - index-index
		return true
	}

	if name == fmt.Sprintf("%d", shardIndex) {
		// old version - index
		return true
	}

	if name == fmt.Sprintf("%d", replicaIndex) {
		// old version - index
		return true
	}

	return false
}

// CreateStatefulSetName creates a name of a StatefulSet for ClickHouse instance
func CreateStatefulSetName(host *chop.ChiHost) string {
	// Name can be generated either from default name pattern,
	// or from personal name pattern provided in PodTemplate

	// Start with default name pattern
	pattern := statefulSetNamePattern

	// PodTemplate may have personal name pattern specified
	if template, ok := host.GetPodTemplate(); ok {
		// PodTemplate available
		if template.GenerateName != "" {
			// PodTemplate has explicitly specified name pattern
			pattern = template.GenerateName
		}
	}

	// Create StatefulSet name based on name pattern available
	return newNameMacroReplacerHost(host).Replace(pattern)
}

// CreateStatefulSetServiceName returns a name of a StatefulSet-related Service for ClickHouse instance
func CreateStatefulSetServiceName(host *chop.ChiHost) string {
	// Name can be generated either from default name pattern,
	// or from personal name pattern provided in ServiceTemplate

	// Start with default name pattern
	pattern := statefulSetServiceNamePattern

	// ServiceTemplate may have personal name pattern specified
	if template, ok := host.GetServiceTemplate(); ok {
		// ServiceTemplate available
		if template.GenerateName != "" {
			// ServiceTemplate has explicitly specified name pattern
			pattern = template.GenerateName
		}
	}

	// Create Service name based on name pattern available
	return newNameMacroReplacerHost(host).Replace(pattern)
}

// CreatePodHostname returns a name of a Pod of a ClickHouse instance
func CreatePodHostname(host *chop.ChiHost) string {
	// Do not use Pod own hostname - point to appropriate StatefulSet's Service
	return CreateStatefulSetServiceName(host)
}

// createPodFQDN creates a fully qualified domain name of a pod
// ss-1eb454-2-0.my-dev-domain.svc.cluster.local
func createPodFQDN(host *chop.ChiHost) string {
	// FQDN can be generated either from default pattern,
	// or from personal pattern provided

	// Start with default pattern
	pattern := podFQDNPattern

	if host.CHI.Spec.NamespaceDomainPattern != "" {
		// NamespaceDomainPattern has been explicitly specified
		pattern = "%s." + host.CHI.Spec.NamespaceDomainPattern
	}

	// Create FQDN based on pattern available
	return fmt.Sprintf(
		pattern,
		CreatePodHostname(host),
		host.Address.Namespace,
	)
}

// createPodFQDNsOfCluster creates fully qualified domain names of all pods in a cluster
func createPodFQDNsOfCluster(cluster *chop.ChiCluster) []string {
	fqdns := make([]string, 0)
	cluster.WalkHosts(func(host *chop.ChiHost) error {
		fqdns = append(fqdns, createPodFQDN(host))
		return nil
	})
	return fqdns
}

// createPodFQDNsOfShard creates fully qualified domain names of all pods in a shard
func createPodFQDNsOfShard(shard *chop.ChiShard) []string {
	fqdns := make([]string, 0)
	shard.WalkHosts(func(host *chop.ChiHost) error {
		fqdns = append(fqdns, createPodFQDN(host))
		return nil
	})
	return fqdns
}

// createPodFQDNsOfCHI creates fully qualified domain names of all pods in a CHI
func createPodFQDNsOfCHI(chi *chop.ClickHouseInstallation) []string {
	fqdns := make([]string, 0)
	chi.WalkHosts(func(host *chop.ChiHost) error {
		fqdns = append(fqdns, createPodFQDN(host))
		return nil
	})
	return fqdns
}

// CreateFQDN is a wrapper over pod FQDN function
func CreateFQDN(host *chop.ChiHost) string {
	return createPodFQDN(host)
}

// CreateFQDNs is a wrapper over set of create FQDN functions
// obj specifies source object to create FQDNs from
// scope specifies target scope - what entity to create FQDNs for - be it CHI, cluster, shard or a host
// excludeSelf specifies whether to exclude the host itself from the result. Applicable only in case obj is a host
func CreateFQDNs(obj interface{}, scope interface{}, excludeSelf bool) []string {
	switch typed := obj.(type) {
	case *chop.ClickHouseInstallation:
		return createPodFQDNsOfCHI(typed)
	case *chop.ChiCluster:
		return createPodFQDNsOfCluster(typed)
	case *chop.ChiShard:
		return createPodFQDNsOfShard(typed)
	case *chop.ChiHost:
		self := ""
		if excludeSelf {
			self = createPodFQDN(typed)
		}
		switch scope.(type) {
		case chop.ChiHost:
			return util.RemoveFromArray(self, []string{createPodFQDN(typed)})
		case chop.ChiShard:
			return util.RemoveFromArray(self, createPodFQDNsOfShard(typed.GetShard()))
		case chop.ChiCluster:
			return util.RemoveFromArray(self, createPodFQDNsOfCluster(typed.GetCluster()))
		case chop.ClickHouseInstallation:
			return util.RemoveFromArray(self, createPodFQDNsOfCHI(typed.GetCHI()))
		}
	}
	return nil
}

// CreatePodRegexp creates pod regexp
// template is defined in operator config:
// CHConfigNetworksHostRegexpTemplate: chi-{chi}-[^.]+\\d+-\\d+\\.{namespace}.svc.cluster.local$"
func CreatePodRegexp(chi *chop.ClickHouseInstallation, template string) string {
	return newNameMacroReplacerChi(chi).Replace(template)
}

// CreatePodName create Pod name based on specified StatefulSet or Host
func CreatePodName(obj interface{}) string {
	switch obj.(type) {
	case *apps.StatefulSet:
		statefulSet := obj.(*apps.StatefulSet)
		return fmt.Sprintf(podNamePattern, statefulSet.Name)
	case *chop.ChiHost:
		host := obj.(*chop.ChiHost)
		return fmt.Sprintf(podNamePattern, CreateStatefulSetName(host))
	}
	return "unknown-type"
}

// CreatePVCName create PVC name from components, to which PVC belongs to
func CreatePVCName(host *chop.ChiHost, _ *v1.VolumeMount, template *chop.ChiVolumeClaimTemplate) string {
	return template.Name + "-" + CreatePodName(host)
}
