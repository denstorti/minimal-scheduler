package main

import (
	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"sigs.k8s.io/scheduler-plugins/pkg/trimaran/loadvariationriskbalancing"
	"sigs.k8s.io/scheduler-plugins/pkg/trimaran/targetloadpacking"

	// Ensure scheme package is initialized.
	_ "sigs.k8s.io/scheduler-plugins/apis/config/scheme"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewSchedulerCommand(
		app.WithPlugin(loadvariationriskbalancing.Name, loadvariationriskbalancing.New),
		app.WithPlugin(targetloadpacking.Name, targetloadpacking.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
