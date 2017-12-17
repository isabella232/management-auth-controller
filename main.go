package main

import (
	"context"
	"os"

	"github.com/rancher/management-auth-controller/controller"
	"github.com/rancher/types/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"k8s.io/client-go/tools/clientcmd"
)

// main is just for testing. The Register function should be called by something else in some other part of rancher
func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		return run()
	}

	app.ExitErrHandler = func(c *cli.Context, err error) {
		logrus.Fatal(err)
	}

	app.Run(os.Args)
}

func run() error {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}

	management, err := config.NewManagementContext(*kubeConfig)
	if err != nil {
		return err
	}

	ctx := context.Background()
	controller.Register(ctx, management)

	return management.StartAndWait()
}
