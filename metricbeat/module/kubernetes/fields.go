// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of module/kubernetes.
func AssetKubernetes() string {
	return "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"
}