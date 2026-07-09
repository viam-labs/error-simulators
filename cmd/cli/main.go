package main

import (
	"context"
	"everythingerrors"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	gripper "go.viam.com/rdk/components/gripper"
)

func main() {
	err := realMain()
	if err != nil {
		panic(err)
	}
}

func realMain() error {
	ctx := context.Background()
	logger := logging.NewLogger("cli")

	deps := resource.Dependencies{}
	// can load these from a remote machine if you need

	cfg := everythingerrors.Config{}

	thing, err := everythingerrors.NewErrorGripper(ctx, deps, gripper.Named("foo"), &cfg, logger)
	if err != nil {
		return err
	}
	defer thing.Close(ctx)

	return nil
}
