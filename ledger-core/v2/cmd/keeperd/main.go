// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"

	"github.com/insolar/assured-ledger/ledger-core/v2/instrumentation/inslogger"
	"github.com/insolar/assured-ledger/ledger-core/v2/log/global"
	"github.com/insolar/assured-ledger/ledger-core/v2/metrics"
)

var configPath string

func main() {
	var rootCmd = &cobra.Command{
		Use: "keeperd --config=<path to config>",
		Run: rootCommand,
	}
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "path to config file")
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Wrong input params:", err.Error())
	}
}

func rootCommand(cmd *cobra.Command, args []string) {
	jww.SetStdoutThreshold(jww.LevelInfo)
	var err error

	vp := viper.New()
	cfg := NewConfig()
	if len(configPath) != 0 {
		vp.SetConfigFile(configPath)
		err = vp.ReadInConfig()
		if err != nil {
			global.Fatal("failed to load configuration from file: ", err.Error())
		}
		err = vp.Unmarshal(&cfg)
		if err != nil {
			global.Fatal("failed to load configuration from file: ", err.Error())
		}
	}

	ctx := context.Background()
	ctx, logger := inslogger.InitNodeLogger(ctx, cfg.Log, "", "keeperd")

	m := metrics.NewMetrics(cfg.Metrics, GetRegistry(), "keeper")
	err = m.Init(ctx)
	if err != nil {
		global.Fatal("Couldn't init metrics:", err)
		os.Exit(1)
	}
	err = m.Start(ctx)
	if err != nil {
		global.Fatal("Couldn't start metrics:", err)
		os.Exit(1)
	}

	keeper := NewKeeper(cfg.Keeper)
	keeper.Run(ctx)

	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Reloading config file")
		cfg := NewConfig()
		err := vp.Unmarshal(&cfg)
		if err != nil {
			logger.Errorf("Failed to reload config: %s", err.Error())
			return
		}
		keeper.SetConfig(cfg.Keeper)
	})

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop
}
