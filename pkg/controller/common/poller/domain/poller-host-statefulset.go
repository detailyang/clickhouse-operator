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

package domain

import (
	"context"

	apps "k8s.io/api/apps/v1"

	log "github.com/altinity/clickhouse-operator/pkg/announcer"
	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
	"github.com/altinity/clickhouse-operator/pkg/model/k8s"
)

type readyMarkDeleter interface {
	DeleteReadyMarkOnPodAndService(ctx context.Context, host *api.Host) error
}

// HostStatefulSetPoller enriches StatefulSet poller with host capabilities
type HostStatefulSetPoller struct {
	*StatefulSetPoller
	interfaces.IKubeSTS
	readyMarkDeleter
}

// NewHostStatefulSetPoller creates new HostStatefulSetPoller from StatefulSet poller
func NewHostStatefulSetPoller(poller *StatefulSetPoller, kube interfaces.IKube, labeler readyMarkDeleter) *HostStatefulSetPoller {
	return &HostStatefulSetPoller{
		StatefulSetPoller: poller,
		IKubeSTS:          kube.STS(),
		readyMarkDeleter:  labeler,
	}
}

// WaitHostStatefulSetReady polls host's StatefulSet until it is ready
func (p *HostStatefulSetPoller) WaitHostStatefulSetReady(ctx context.Context, host *api.Host) error {
	// Wait for StatefulSet to reach generation
	err := p.PollHostStatefulSet(
		ctx,
		host,
		func(_ctx context.Context, sts *apps.StatefulSet) bool {
			if sts == nil {
				return false
			}
			p.deleteReadyMark(_ctx, host)
			return k8s.IsStatefulSetGeneration(sts, sts.Generation)
		},
		func(_ctx context.Context) {
			p.deleteReadyMark(_ctx, host)
		},
	)
	if err != nil {
		return err
	}

	// Wait StatefulSet to reach ready status
	err = p.PollHostStatefulSet(
		ctx,
		host,
		func(_ctx context.Context, sts *apps.StatefulSet) bool {
			p.deleteReadyMark(_ctx, host)
			return k8s.IsStatefulSetReady(sts)
		},
		func(_ctx context.Context) {
			p.deleteReadyMark(_ctx, host)
		},
	)

	return err
}

//// waitHostNotReady polls host's StatefulSet for not exists or not ready
//func (c *HostStatefulSetPoller) WaitHostNotReady(ctx context.Context, host *api.Host) error {
//	err := c.PollHostStatefulSet(
//		ctx,
//		host,
//		// Since we are waiting for host to be nopt readylet's assyme that it should exist already
//		// and thus let's set GetErrorTimeout to zero, since we are not expecting getter function
//		// to return any errors
//		poller.NewPollerOptions().
//			FromConfig(chop.Config()).
//			SetGetErrorTimeout(0),
//		func(_ context.Context, sts *apps.StatefulSet) bool {
//			return k8s.IsStatefulSetNotReady(sts)
//		},
//		nil,
//	)
//	if apiErrors.IsNotFound(err) {
//		err = nil
//	}
//
//	return err
//}

//// WaitHostStatefulSetDeleted polls host's StatefulSet until it is not available
//func (p *HostStatefulSetPoller) WaitHostStatefulSetDeleted(host *api.Host) {
//	for {
//		// TODO
//		// Probably there would be better way to wait until k8s reported StatefulSet deleted
//		if _, err := p.IKubeSTS.Get(context.TODO(), host); err == nil {
//			log.V(2).Info("cache NOT yet synced")
//			time.Sleep(15 * time.Second)
//		} else {
//			log.V(1).Info("cache synced")
//			return
//		}
//	}
//}

func (p *HostStatefulSetPoller) deleteReadyMark(ctx context.Context, host *api.Host) {
	if p == nil {
		return
	}
	if p.readyMarkDeleter == nil {
		log.V(3).F().Info("no mark deleter specified")
		return
	}

	log.V(3).F().Info("Has mark deleter specified")
	_ = p.readyMarkDeleter.DeleteReadyMarkOnPodAndService(ctx, host)
}
