package main

import (
	"everythingerrors"
	gripper "go.viam.com/rdk/components/gripper"
	motor "go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

func main() {
	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(
		resource.APIModel{API: gripper.API, Model: everythingerrors.ErrorGripper},
		resource.APIModel{API: motor.API, Model: everythingerrors.ErrorMotor},
	)
}
