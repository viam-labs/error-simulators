package main

import (
	"everythingerrors"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
	gripper "go.viam.com/rdk/components/gripper"
)

func main() {
	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(resource.APIModel{ gripper.API, everythingerrors.ErrorGripper})
}
