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

package statefulset

type ReconcileOptions struct {
	forceRecreate bool
	doNotWait     bool
}

func NewReconcileStatefulSetOptions() *ReconcileOptions {
	return &ReconcileOptions{}
}

func (o *ReconcileOptions) Ensure() *ReconcileOptions {
	if o == nil {
		o = NewReconcileStatefulSetOptions()
	}
	return o
}

func (o *ReconcileOptions) SetForceRecreate() *ReconcileOptions {
	o = o.Ensure()
	o.forceRecreate = true
	return o
}

func (o *ReconcileOptions) IsForceRecreate() bool {
	if o == nil {
		return false
	}
	return o.forceRecreate
}

func (o *ReconcileOptions) SetDoNotWait() *ReconcileOptions {
	o = o.Ensure()
	o.doNotWait = true
	return o
}

func (o *ReconcileOptions) IsDoNotWait() bool {
	if o == nil {
		return false
	}
	return o.doNotWait
}

type ReconcileOptionsSet []*ReconcileOptions

// NewReconcileOptionsSet creates new reconcileHostStatefulSetOptions array
func NewReconcileOptionsSet(opts ...*ReconcileOptions) (res ReconcileOptionsSet) {
	return append(res, opts...)
}

// First gets first option
func (a ReconcileOptionsSet) First() *ReconcileOptions {
	if len(a) > 0 {
		return a[0]
	}
	return nil
}
