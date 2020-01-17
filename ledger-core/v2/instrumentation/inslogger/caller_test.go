//
// Copyright 2019 Insolar Technologies GmbH
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
//

package inslogger

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/insolar/assured-ledger/ledger-core/v2/configuration"
	"github.com/insolar/assured-ledger/ledger-core/v2/log"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/global"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/logcommon"
)

// Beware, test results there depends on test file name (caller_test.go)!

const pkgRegexPrefix = "^instrumentation/inslogger/"

type loggerField struct {
	Caller string
	Func   string
}

func logFields(t *testing.T, b []byte) loggerField {
	var lf loggerField
	err := json.Unmarshal(b, &lf)
	require.NoErrorf(t, err, "failed decode: '%v'", string(b))
	return lf
}

func TestLog_ZerologCaller(t *testing.T) {
	l, err := NewLog(configuration.Log{
		Level:     "info",
		Adapter:   "zerolog",
		Formatter: "json",
	})
	require.NoError(t, err, "log creation")

	var b bytes.Buffer
	l, err = l.Copy().WithOutput(&b).WithCaller(logcommon.CallerField).Build()
	require.NoError(t, err)

	_, _, line, _ := runtime.Caller(0)
	l.Info("test")

	lf := logFields(t, b.Bytes())
	assert.Regexp(t, pkgRegexPrefix+"caller_test.go:"+strconv.Itoa(line+1), lf.Caller, "log contains call place")
	assert.NotContains(t, "github.com/insolar/assured-ledger/ledger-core/v2", lf.Caller, "log not contains package name")
	assert.Equal(t, "", lf.Func, "log not contains func name")
}

// this test result depends on test name!
func TestLog_ZerologCallerWithFunc(t *testing.T) {
	l, err := NewLog(configuration.Log{
		Level:     "info",
		Adapter:   "zerolog",
		Formatter: "json",
	})
	require.NoError(t, err, "log creation")

	var b bytes.Buffer
	l, err = l.Copy().WithOutput(&b).WithCaller(logcommon.CallerFieldWithFuncName).Build()
	require.NoError(t, err)

	_, _, line, _ := runtime.Caller(0)
	l.Info("test")

	lf := logFields(t, b.Bytes())
	assert.Regexp(t, pkgRegexPrefix+"caller_test.go:"+strconv.Itoa(line+1), lf.Caller, "log contains proper caller place")
	assert.NotContains(t, "github.com/insolar/assured-ledger/ledger-core/v2", lf.Caller, "log not contains package name")
	assert.Equal(t, "TestLog_ZerologCallerWithFunc", lf.Func, "log contains func name")
}

func TestLog_GlobalCaller(t *testing.T) {
	defer global.SaveLogger()()

	var b bytes.Buffer
	gl2, err := global.Logger().Copy().WithOutput(&b).WithCaller(logcommon.CallerField).Build()
	require.NoError(t, err)
	global.SetLogger(gl2)
	global.SetLevel(log.InfoLevel)

	_, _, line, _ := runtime.Caller(0)
	global.Info("test")
	global.Debug("test2shouldNotBeThere")

	s := b.String()
	lf := logFields(t, []byte(s))
	assert.Regexp(t, pkgRegexPrefix+"caller_test.go:"+strconv.Itoa(line+1), lf.Caller, "log contains proper call place")
	assert.Equal(t, "", lf.Func, "log not contains func name")
	assert.NotContains(t, s, "test2shouldNotBeThere")
}

func TestLog_GlobalCallerWithFunc(t *testing.T) {
	defer global.SaveLogger()()

	var b bytes.Buffer
	gl2, err := global.Logger().Copy().WithOutput(&b).WithCaller(logcommon.CallerFieldWithFuncName).Build()
	require.NoError(t, err)
	global.SetLogger(gl2)
	global.SetLevel(log.InfoLevel)

	_, _, line, _ := runtime.Caller(0)
	global.Info("test")
	global.Debug("test2shouldNotBeThere")

	s := b.String()
	lf := logFields(t, []byte(s))
	assert.Regexp(t, pkgRegexPrefix+"caller_test.go:"+strconv.Itoa(line+1), lf.Caller, "log contains proper call place")
	assert.Equal(t, "TestLog_GlobalCallerWithFunc", lf.Func, "log contains func name")
	assert.NotContains(t, s, "test2shouldNotBeThere")
}
