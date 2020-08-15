/*
    Copyright (C) 2020 Accurics, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package writer

import (
	"io"

	policyengine "github.com/accurics/terrascan/pkg/policy-engine"
)

// supportedFormat data type for supported formats
type supportedFormat string

// writerMap stores mapping of supported writer formats with respective functions
var writerMap = make(map[supportedFormat](func(policyengine.EngineOutput, io.Writer) error))

// RegisterWriter registers a writer for terrascan
func RegisterWriter(format supportedFormat, writerFunc func(policyengine.EngineOutput, io.Writer) error) {
	writerMap[format] = writerFunc
}