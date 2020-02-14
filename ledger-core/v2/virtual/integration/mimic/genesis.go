// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package mimic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/insolar/assured-ledger/ledger-core/v2/application"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/record"
	"github.com/insolar/assured-ledger/ledger-core/v2/ledger/drop"
)

const (
	LaunchnetRelativePath = "scripts/insolard/launchnet.sh"
	GenesisRelativePath   = "launchnet/configs/heavy_genesis.json"
)

func GenerateBootstrap(skipBuild bool) (func(), string, error) {
	artifactsDir, err := ioutil.TempDir("", "mimic")
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to create temporary directory")
	}

	cleanupFunc := func() {
		err := os.RemoveAll(artifactsDir)
		if err != nil {
			fmt.Printf("[ Error ] Failed to cleanup temporary dir %s: %s\n", artifactsDir, err.Error())
		}
	}

	cmd := exec.Command(LaunchnetRelativePath, "-b")
	cmd.Dir = insolar.RootModuleDir()
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "INSOLAR_ARTIFACTS_DIR="+artifactsDir)
	if skipBuild {
		cmd.Env = append(cmd.Env, "SKIP_BUILD=1")
	}

	if output, err := cmd.CombinedOutput(); err != nil {
		cleanupFunc()

		fmt.Printf("[ Error ] Failed to execute bootstrap: %s\n", err.Error())
		fmt.Printf("[ Error ] Output of bootstrap is:\n")

		outputString := string(bytes.TrimSpace(output))
		for _, line := range strings.Split(outputString, "\n") {
			fmt.Printf("[ Error ] > %s", line)
		}

		return nil, "", errors.Wrapf(err, "Failed to execute bootstrap: %s", err.Error())
	}

	return cleanupFunc, artifactsDir, nil
}

func ReadGenesisContractsConfig(dirPath string) (*application.GenesisContractsConfig, error) {
	genesisConfigPath := path.Join(dirPath, GenesisRelativePath)

	fh, err := os.Open(genesisConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open genesis config for reading")
	}

	rv := application.GenesisHeavyConfig{}
	if err := json.NewDecoder(fh).Decode(&rv); err != nil {
		return nil, errors.Wrap(err, "failed to decode genesis config")
	}

	return &rv.ContractsConfig, nil
}

type recordModifierMock struct{}

func (d dropModifierMock) Set(_ context.Context, _ drop.Drop) error {
	return nil
}

type dropModifierMock struct{}

func (r recordModifierMock) Set(_ context.Context, _ record.Material) error {
	return nil
}
func (r recordModifierMock) BatchSet(_ context.Context, _ []record.Material) error {
	return nil
}

type indexModifierMock struct{}

func (i indexModifierMock) UpdateLastKnownPulse(_ context.Context, _ insolar.PulseNumber) error {
	return nil
}
func (i indexModifierMock) SetIndex(_ context.Context, _ insolar.PulseNumber, _ record.Index) error {
	return nil
}
