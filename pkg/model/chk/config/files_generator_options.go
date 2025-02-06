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

package config

import (
	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/model/common/config"
)

// FilesGeneratorOptions specifies options for configuration files generator
type FilesGeneratorOptions struct {
	host        *api.Host
	RaftOptions *config.HostSelector
}

// defaultFilesGeneratorOptions creates new default options for files generator
func defaultFilesGeneratorOptions() *FilesGeneratorOptions {
	return NewFilesGeneratorOptions()
}

// NewFilesGeneratorOptions creates new options for configuration files generator
func NewFilesGeneratorOptions() *FilesGeneratorOptions {
	return &FilesGeneratorOptions{}
}

func (o *FilesGeneratorOptions) GetHost() *api.Host {
	if o == nil {
		return nil
	}
	return o.host
}

func (o *FilesGeneratorOptions) SetHost(host *api.Host) *FilesGeneratorOptions {
	if o == nil {
		return nil
	}
	o.host = host

	return o
}

// GetRaftOptions gets RAFT generator options
func (o *FilesGeneratorOptions) GetRaftOptions() *config.HostSelector {
	if o == nil {
		return nil
	}
	return o.RaftOptions
}

// SetRaftOptions sets RAFT generator options
func (o *FilesGeneratorOptions) SetRaftOptions(opts *config.HostSelector) *FilesGeneratorOptions {
	if o == nil {
		return nil
	}
	o.RaftOptions = opts

	return o
}
