// Copyright The Karbour Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package audit

import "github.com/KusionStack/karbour/pkg/core/handler"

// Ensure that ClusterPayload implements the handler.Payload interface.
var _ handler.Payload = &AuditPayload{}

// Payload represents the structure for audit request data. It includes the
// manifest which is typically a string containing declarative configuration
// data that needs to be audited.
type AuditPayload struct {
	Manifest string `json:"manifest"` // Manifest is the content to be audited.
}
