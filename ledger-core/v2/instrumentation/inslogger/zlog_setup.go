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
	"errors"

	"github.com/rs/zerolog"

	"github.com/insolar/assured-ledger/ledger-core/v2/log"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/adapters/zlog"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/logcommon"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/logmsgfmt"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/logoutput"
)

func initZlog() {
	zerolog.CallerMarshalFunc = fileLineMarshaller
	zerolog.TimeFieldFormat = TimestampFormat
	zerolog.CallerFieldName = logoutput.CallerFieldName
	zerolog.LevelFieldName = logoutput.LevelFieldName
	zerolog.TimestampFieldName = logoutput.TimestampFieldName
	zerolog.MessageFieldName = logoutput.MessageFieldName
}

func newZerologAdapter(pCfg ParsedLogConfig, msgFmt logmsgfmt.MsgFormatConfig) (log.LoggerBuilder, error) {
	zc := logcommon.Config{}

	var err error
	zc.BareOutput, err = logoutput.OpenLogBareOutput(pCfg.OutputType, pCfg.OutputParam)
	if err != nil {
		return log.LoggerBuilder{}, err
	}
	if zc.BareOutput.Writer == nil {
		return log.LoggerBuilder{}, errors.New("output is nil")
	}

	sfb := zlog.ZerologSkipFrameCount + pCfg.SkipFrameBaselineAdjustment
	if sfb < 0 {
		sfb = 0
	}

	zc.Output = pCfg.Output
	zc.Instruments = pCfg.Instruments
	zc.MsgFormat = msgFmt
	zc.Instruments.SkipFrameCountBaseline = uint8(sfb)

	return log.NewBuilder(zlog.NewFactory(), zc, pCfg.LogLevel), nil
}
