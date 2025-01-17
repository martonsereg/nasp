// Copyright (c) 2022 Cisco and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2etest

import _ "embed"

var (
	// contains one echo endpoint in the demo namespace as the 'echo'
	//go:embed config_v0.json
	configV0 []byte

	// contains three echo endpoints in the demo namespace as the 'echo' deployment has been scaled up to three replicas
	//go:embed config_v1.json
	configV1 []byte
)
