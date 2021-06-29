// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtobufLogsUnmarshaler_error(t *testing.T) {
	p := NewProtobufLogsUnmarshaler()
	_, err := p.UnmarshalLogs([]byte("+$%"))
	assert.Error(t, err)
}

func TestProtobufMetricsUnmarshaler_error(t *testing.T) {
	p := NewProtobufMetricsUnmarshaler()
	_, err := p.UnmarshalMetrics([]byte("+$%"))
	assert.Error(t, err)
}

func TestProtobufTracesUnmarshaler_error(t *testing.T) {
	p := NewProtobufTracesUnmarshaler()
	_, err := p.UnmarshalTraces([]byte("+$%"))
	assert.Error(t, err)
}