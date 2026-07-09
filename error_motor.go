package everythingerrors

import (
	"context"

	"fmt"

	motor "go.viam.com/rdk/components/motor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var (
	ErrorMotor = resource.NewModel("viam", "everything-errors", "error-motor")
)

func init() {
	resource.RegisterComponent(motor.API, ErrorMotor,
		resource.Registration[motor.Motor, *MotorConfig]{
			Constructor: newEverythingErrorsErrorMotor,
		},
	)
}

type MotorConfig struct {
	/*
		Put config attributes here. There should be public/exported fields
		with a `json` parameter at the end of each attribute.

		If your model does not need a config, replace *MotorConfig in the init
		function with resource.NoNativeConfig
	*/
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns three values:
//  1. Required dependencies: other resources that must exist for this resource to work.
//  2. Optional dependencies: other resources that may exist but are not required.
//  3. An error if any Config fields are missing or invalid.
func (cfg *MotorConfig) Validate(path string) ([]string, []string, error) {
	// Add config validation code here
	return nil, nil, nil
}

type everythingErrorsErrorMotor struct {
	resource.AlwaysRebuild
	resource.Named

	name resource.Name

	logger logging.Logger
	cfg    *MotorConfig

	cancelCtx  context.Context
	cancelFunc func()
}

func newEverythingErrorsErrorMotor(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (motor.Motor, error) {
	conf, err := resource.NativeConfig[*MotorConfig](rawConf)
	if err != nil {
		return nil, err
	}

	return NewErrorMotor(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewErrorMotor(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *MotorConfig, logger logging.Logger) (motor.Motor, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &everythingErrorsErrorMotor{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *everythingErrorsErrorMotor) Name() resource.Name {
	return s.name
}

// SetPower sets the percentage of power the motor should employ between -1 and 1.
func (s *everythingErrorsErrorMotor) SetPower(ctx context.Context, powerPct float64, extra map[string]interface{}) error {
	return fmt.Errorf("the SetPower operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the requested power percentage was not applied, the motor did not begin or change rotation, and no electrical output was produced, so please implement the SetPower method before attempting to drive this motor")
}

// GoFor instructs the motor to go a specific number of revolutions at a given speed in RPM.
func (s *everythingErrorsErrorMotor) GoFor(ctx context.Context, rpm, revolutions float64, extra map[string]interface{}) error {
	return fmt.Errorf("the GoFor operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the motor did not rotate for the requested number of revolutions at the requested speed and did not perform any blocking motion, so please implement the GoFor method before attempting to command relative motion on this motor")
}

// GoTo instructs the motor to go to a specific position (in revolutions from home) at a specific speed.
func (s *everythingErrorsErrorMotor) GoTo(ctx context.Context, rpm, positionRevolutions float64, extra map[string]interface{}) error {
	return fmt.Errorf("the GoTo operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the motor did not move toward the requested absolute position at the requested speed and never reached the target, so please implement the GoTo method before attempting to command absolute positioning on this motor")
}

// SetRPM instructs the motor to move at the specified RPM indefinitely.
func (s *everythingErrorsErrorMotor) SetRPM(ctx context.Context, rpm float64, extra map[string]interface{}) error {
	return fmt.Errorf("the SetRPM operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the motor did not begin rotating at the requested revolutions-per-minute and no continuous motion was started, so please implement the SetRPM method before attempting to command a target speed on this motor")
}

// ResetZeroPosition sets the current position (+/- offset) to be the new zero (home) position.
func (s *everythingErrorsErrorMotor) ResetZeroPosition(ctx context.Context, offset float64, extra map[string]interface{}) error {
	return fmt.Errorf("the ResetZeroPosition operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the encoder's home/zero reference was not updated with the requested offset and the reported position remains unchanged, so please implement the ResetZeroPosition method before attempting to redefine this motor's home position")
}

// Position reports the position of an encoded motor based on its encoder, in revolutions.
func (s *everythingErrorsErrorMotor) Position(ctx context.Context, extra map[string]interface{}) (float64, error) {
	return 0, fmt.Errorf("the Position operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the current encoder position in revolutions could not be measured or reported, so the returned value of zero is not meaningful and should not be trusted until the Position method is implemented")
}

// Properties returns whether or not the motor supports certain optional properties.
func (s *everythingErrorsErrorMotor) Properties(ctx context.Context, extra map[string]interface{}) (motor.Properties, error) {
	return motor.Properties{}, fmt.Errorf("the Properties operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the set of optional motor capabilities such as position reporting could not be determined, so the returned properties are empty and unreliable until the Properties method is implemented")
}

// IsPowered returns whether or not the motor is currently on, and the percent power between 0 and 1.
func (s *everythingErrorsErrorMotor) IsPowered(ctx context.Context, extra map[string]interface{}) (bool, float64, error) {
	return false, 0, fmt.Errorf("the IsPowered operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the motor could not determine whether it is currently energized or report its current power percentage, so the returned values should not be trusted until the IsPowered method is implemented")
}

func (s *everythingErrorsErrorMotor) Stop(ctx context.Context, extra map[string]interface{}) error {
	return fmt.Errorf("the Stop operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the motor could not be commanded to halt all movement, and any ongoing rotation was not actively stopped, so please implement the Stop method before relying on this motor to cease motion")
}

func (s *everythingErrorsErrorMotor) IsMoving(ctx context.Context) (bool, error) {
	return false, fmt.Errorf("the IsMoving operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the motor could not determine or report whether it is currently in motion, so the returned value should not be trusted until the IsMoving method is implemented")
}

func (s *everythingErrorsErrorMotor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("the DoCommand operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the supplied command map was not processed, no custom behavior was executed, and no result payload could be produced, so please implement the DoCommand method before sending arbitrary commands to this motor")
}

func (s *everythingErrorsErrorMotor) Status(ctx context.Context) (map[string]interface{}, error) {
	return nil, fmt.Errorf("the Status operation on the everything-errors error-motor component could not be completed because this method has not yet been implemented; the current status of the motor could not be assembled or reported, so no status map is available until the Status method is implemented")
}

func (s *everythingErrorsErrorMotor) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
