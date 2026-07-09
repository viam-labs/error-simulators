package everythingerrors

import (
	"context"
	"errors"
	"fmt"

	gripper "go.viam.com/rdk/components/gripper"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var (
	ErrorGripper     = resource.NewModel("viam", "everything-errors", "error-gripper")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterComponent(gripper.API, ErrorGripper,
		resource.Registration[gripper.Gripper, *Config]{
			Constructor: newEverythingErrorsErrorGripper,
		},
	)
}

type Config struct {
	/*
		Put config attributes here. There should be public/exported fields
		with a `json` parameter at the end of each attribute.

		Example config struct:
			type Config struct {
				Pin   string `json:"pin"`
				Board string `json:"board"`
				MinDeg *float64 `json:"min_angle_deg,omitempty"`
			}

		If your model does not need a config, replace *Config in the init
		function with resource.NoNativeConfig
	*/
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns three values:
//  1. Required dependencies: other resources that must exist for this resource to work.
//  2. Optional dependencies: other resources that may exist but are not required.
//  3. An error if any Config fields are missing or invalid.
//
// The `path` parameter indicates
// where this resource appears in the machine's JSON configuration
// (for example, "components.0"). You can use it in error messages
// to indicate which resource has a problem.
func (cfg *Config) Validate(path string) ([]string, []string, error) {
	// Add config validation code here
	return nil, nil, nil
}

type everythingErrorsErrorGripper struct {
	resource.AlwaysRebuild
	resource.Named

	name resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
}

func newEverythingErrorsErrorGripper(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (gripper.Gripper, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	return NewErrorGripper(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewErrorGripper(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *Config, logger logging.Logger) (gripper.Gripper, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &everythingErrorsErrorGripper{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *everythingErrorsErrorGripper) Name() resource.Name {
	return s.name
}

// Open opens the gripper.
// This will block until done or a new operation cancels this one.
func (s *everythingErrorsErrorGripper) Open(ctx context.Context, extra map[string]interface{}) error {
	return fmt.Errorf("the Open operation on the everything-errors error-gripper component could not be completed because this method has not yet been implemented; the gripper was unable to transition to its fully open position and no motion was performed, so please implement the Open method before attempting to command this gripper to open")
}

// Grab makes the gripper grab.
// returns true if we grabbed something.
// This will block until done or a new operation cancels this one.
func (s *everythingErrorsErrorGripper) Grab(ctx context.Context, extra map[string]interface{}) (bool, error) {
	return false, fmt.Errorf("the Grab operation on the everything-errors error-gripper component could not be completed because this method has not yet been implemented; the gripper did not attempt to close on or secure any object and therefore cannot report whether something was successfully grabbed, so please implement the Grab method before attempting to command this gripper to grab")
}

// IsHoldingSomething returns whether the gripper is currently holding onto an object.
func (s *everythingErrorsErrorGripper) IsHoldingSomething(ctx context.Context, extra map[string]interface{}) (gripper.HoldingStatus, error) {
	var holdingStatusRetVal gripper.HoldingStatus

	return holdingStatusRetVal, fmt.Errorf("the IsHoldingSomething operation on the everything-errors error-gripper component could not be completed because this method has not yet been implemented; the gripper was unable to determine or report whether it is currently holding onto an object, so the returned holding status should not be trusted, and please implement the IsHoldingSomething method before relying on this status")
}

func (s *everythingErrorsErrorGripper) Stop(ctx context.Context, extra map[string]interface{}) error {
	return fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("the DoCommand operation on the everything-errors error-gripper component could not be completed because this method has not yet been implemented; the supplied command map was not processed, no custom behavior was executed, and no result payload could be produced, so please implement the DoCommand method before sending arbitrary commands to this gripper")
}

func (s *everythingErrorsErrorGripper) Status(ctx context.Context) (map[string]interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) IsMoving(ctx context.Context) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) Geometries(ctx context.Context, extra map[string]interface{}) ([]spatialmath.Geometry, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) Kinematics(ctx context.Context) (referenceframe.Model, error) {
	var modelRetVal referenceframe.Model

	return modelRetVal, fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) GoToInputs(ctx context.Context, inputSteps ...[]referenceframe.Input) error {
	return fmt.Errorf("not implemented")
}

func (s *everythingErrorsErrorGripper) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
