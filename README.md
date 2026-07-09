# Module everything-errors

`everything-errors` is a collection of Viam mock/simulator components whose sole purpose is to fail. Every model in this module initializes successfully and can be added to a machine like any real component, but **every API method past initialization returns an error**.

This is deliberate. The module is a testing and development aid for exercising the *unhappy path* of your system: it lets you verify how your code, control logic, data pipelines, alerting, and UI behave when a component is present and configured but consistently refuses to do any real work.

## Why use this module?

- **Test error handling** — Confirm that callers, remotes, and higher-level services surface and handle component errors gracefully instead of crashing or hanging.
- **Validate observability** — Make sure failures show up where you expect them: logs, alerts, dashboards, and the Viam app.
- **Exercise retry/fallback logic** — Point your fallback or redundancy logic at a component that is guaranteed to fail on every call.
- **Reproduce failure states** — Get a deterministic, always-failing component without having to physically break real hardware.

## Behavior

Each model in this module:

1. Constructs and initializes successfully — the resource comes up healthy and is registered on the machine.
2. Returns a descriptive error from every functional API method (e.g. `Open`, `Grab`, `DoCommand`, etc.).

The `Close` method is the only exception: it cleans up the resource normally so the component can be reconfigured or removed without leaking resources.

## Models

This module provides the following model(s):

- [`viam:everything-errors:error-gripper`](viam_everything-errors_error-gripper.md) - A mock [gripper](https://docs.viam.com/components/gripper/) that initializes successfully and then returns an error from every gripper API call.

## Configuration

None of the models in this module require any configuration attributes. Add a model to your machine, and it will initialize successfully with no attributes required. See each model's documentation page for details.
