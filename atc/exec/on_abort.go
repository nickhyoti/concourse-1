package exec

import (
	"context"
	"errors"
)

// OnAbortStep will run one step, and then a second step if the first step
// aborts (but not errors).
type OnAbortStep struct {
	step Step
	hook Step
}

// OnAbort constructs an OnAbortStep factory.
func OnAbort(step Step, hook Step) OnAbortStep {
	return OnAbortStep{
		step: step,
		hook: hook,
	}
}

// Run will call Run on the first step and wait for it to complete. If the
// first step errors, Run returns the error. OnAbortStep is ready as soon as
// the first step is ready.
//
// If the first step aborts (that is, it gets interrupted), the second
// step is executed. If the second step errors, its error is returned.
func (o OnAbortStep) Run(ctx context.Context, state RunState) (bool, error) {
	stepRunOk, stepRunErr := o.step.Run(ctx, state)
	if stepRunErr == nil {
		return stepRunOk, nil
	}

	if errors.Is(stepRunErr, context.Canceled) {
		// run only on abort, not timeout
		o.hook.Run(context.Background(), state)
	}

	return stepRunOk, stepRunErr
}
