# Model viam:everything-errors:error-gripper

`error-gripper` is a mock/simulator implementation of the Viam [gripper](https://docs.viam.com/components/gripper/) API. It initializes successfully and behaves like a normal gripper resource on the machine, but **every gripper API method returns an error**.

Use it to test how your application, control code, and observability tooling respond to a gripper that is present and configured but never actually performs any motion or reports any real state.

## Behavior

- **Initialization succeeds.** The component constructs without error and comes up healthy.
- **Every functional API call fails.** `Open`, `Grab`, `IsHoldingSomething`, `Stop`, `DoCommand`, `IsMoving`, `Geometries`, `Kinematics`, `CurrentInputs`, and `GoToInputs` all return a descriptive error.
- **`Close` succeeds.** The resource cleans up normally so it can be reconfigured or removed.

## Configuration

This model does not require any configuration attributes. An empty attribute object is sufficient:

```json
{}
```

### Attributes

This model has no configurable attributes.

### Example Configuration

```json
{
  "name": "gripper-1",
  "model": "viam:everything-errors:error-gripper",
  "type": "gripper",
  "namespace": "rdk",
  "attributes": {}
}
```

## DoCommand

`DoCommand` is implemented, but like every other method on this model it always returns an error. Any command payload you send will be rejected.

### Example DoCommand

```json
{
  "any_command": {
    "arg1": "foo",
    "arg2": 1
  }
}
```

The call above will return an error rather than a result.
