// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"io"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/worker"
	"github.com/concourse/concourse/tracing"
	"go.opentelemetry.io/otel/api/trace"
)

type FakeTaskDelegate struct {
	ErroredStub        func(lager.Logger, string)
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	FetchImageStub        func(context.Context, atc.ImageResource, atc.VersionedResourceTypes, bool) (worker.ImageSpec, error)
	fetchImageMutex       sync.RWMutex
	fetchImageArgsForCall []struct {
		arg1 context.Context
		arg2 atc.ImageResource
		arg3 atc.VersionedResourceTypes
		arg4 bool
	}
	fetchImageReturns struct {
		result1 worker.ImageSpec
		result2 error
	}
	fetchImageReturnsOnCall map[int]struct {
		result1 worker.ImageSpec
		result2 error
	}
	FinishedStub        func(lager.Logger, exec.ExitStatus, worker.ContainerPlacementStrategy, worker.Client)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 worker.ContainerPlacementStrategy
		arg4 worker.Client
	}
	InitializingStub        func(lager.Logger)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 lager.Logger
	}
	SelectWorkerStub        func(context.Context, worker.Pool, db.ContainerOwner, worker.ContainerSpec, worker.WorkerSpec, worker.ContainerPlacementStrategy, time.Duration, time.Duration) (worker.Client, error)
	selectWorkerMutex       sync.RWMutex
	selectWorkerArgsForCall []struct {
		arg1 context.Context
		arg2 worker.Pool
		arg3 db.ContainerOwner
		arg4 worker.ContainerSpec
		arg5 worker.WorkerSpec
		arg6 worker.ContainerPlacementStrategy
		arg7 time.Duration
		arg8 time.Duration
	}
	selectWorkerReturns struct {
		result1 worker.Client
		result2 error
	}
	selectWorkerReturnsOnCall map[int]struct {
		result1 worker.Client
		result2 error
	}
	SelectedWorkerStub        func(lager.Logger, string)
	selectedWorkerMutex       sync.RWMutex
	selectedWorkerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	SetTaskConfigStub        func(atc.TaskConfig)
	setTaskConfigMutex       sync.RWMutex
	setTaskConfigArgsForCall []struct {
		arg1 atc.TaskConfig
	}
	StartSpanStub        func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)
	startSpanMutex       sync.RWMutex
	startSpanArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}
	startSpanReturns struct {
		result1 context.Context
		result2 trace.Span
	}
	startSpanReturnsOnCall map[int]struct {
		result1 context.Context
		result2 trace.Span
	}
	StartingStub        func(lager.Logger)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		arg1 lager.Logger
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskDelegate) Errored(arg1 lager.Logger, arg2 string) {
	fake.erroredMutex.Lock()
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.ErroredStub
	fake.recordInvocation("Errored", []interface{}{arg1, arg2})
	fake.erroredMutex.Unlock()
	if stub != nil {
		fake.ErroredStub(arg1, arg2)
	}
}

func (fake *FakeTaskDelegate) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *FakeTaskDelegate) ErroredCalls(stub func(lager.Logger, string)) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *FakeTaskDelegate) ErroredArgsForCall(i int) (lager.Logger, string) {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	argsForCall := fake.erroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskDelegate) FetchImage(arg1 context.Context, arg2 atc.ImageResource, arg3 atc.VersionedResourceTypes, arg4 bool) (worker.ImageSpec, error) {
	fake.fetchImageMutex.Lock()
	ret, specificReturn := fake.fetchImageReturnsOnCall[len(fake.fetchImageArgsForCall)]
	fake.fetchImageArgsForCall = append(fake.fetchImageArgsForCall, struct {
		arg1 context.Context
		arg2 atc.ImageResource
		arg3 atc.VersionedResourceTypes
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.FetchImageStub
	fakeReturns := fake.fetchImageReturns
	fake.recordInvocation("FetchImage", []interface{}{arg1, arg2, arg3, arg4})
	fake.fetchImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskDelegate) FetchImageCallCount() int {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	return len(fake.fetchImageArgsForCall)
}

func (fake *FakeTaskDelegate) FetchImageCalls(stub func(context.Context, atc.ImageResource, atc.VersionedResourceTypes, bool) (worker.ImageSpec, error)) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = stub
}

func (fake *FakeTaskDelegate) FetchImageArgsForCall(i int) (context.Context, atc.ImageResource, atc.VersionedResourceTypes, bool) {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	argsForCall := fake.fetchImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskDelegate) FetchImageReturns(result1 worker.ImageSpec, result2 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	fake.fetchImageReturns = struct {
		result1 worker.ImageSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDelegate) FetchImageReturnsOnCall(i int, result1 worker.ImageSpec, result2 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	if fake.fetchImageReturnsOnCall == nil {
		fake.fetchImageReturnsOnCall = make(map[int]struct {
			result1 worker.ImageSpec
			result2 error
		})
	}
	fake.fetchImageReturnsOnCall[i] = struct {
		result1 worker.ImageSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDelegate) Finished(arg1 lager.Logger, arg2 exec.ExitStatus, arg3 worker.ContainerPlacementStrategy, arg4 worker.Client) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 worker.ContainerPlacementStrategy
		arg4 worker.Client
	}{arg1, arg2, arg3, arg4})
	stub := fake.FinishedStub
	fake.recordInvocation("Finished", []interface{}{arg1, arg2, arg3, arg4})
	fake.finishedMutex.Unlock()
	if stub != nil {
		fake.FinishedStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *FakeTaskDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeTaskDelegate) FinishedCalls(stub func(lager.Logger, exec.ExitStatus, worker.ContainerPlacementStrategy, worker.Client)) {
	fake.finishedMutex.Lock()
	defer fake.finishedMutex.Unlock()
	fake.FinishedStub = stub
}

func (fake *FakeTaskDelegate) FinishedArgsForCall(i int) (lager.Logger, exec.ExitStatus, worker.ContainerPlacementStrategy, worker.Client) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	argsForCall := fake.finishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskDelegate) Initializing(arg1 lager.Logger) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.InitializingStub
	fake.recordInvocation("Initializing", []interface{}{arg1})
	fake.initializingMutex.Unlock()
	if stub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakeTaskDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakeTaskDelegate) InitializingCalls(stub func(lager.Logger)) {
	fake.initializingMutex.Lock()
	defer fake.initializingMutex.Unlock()
	fake.InitializingStub = stub
}

func (fake *FakeTaskDelegate) InitializingArgsForCall(i int) lager.Logger {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	argsForCall := fake.initializingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskDelegate) SelectWorker(arg1 context.Context, arg2 worker.Pool, arg3 db.ContainerOwner, arg4 worker.ContainerSpec, arg5 worker.WorkerSpec, arg6 worker.ContainerPlacementStrategy, arg7 time.Duration, arg8 time.Duration) (worker.Client, error) {
	fake.selectWorkerMutex.Lock()
	ret, specificReturn := fake.selectWorkerReturnsOnCall[len(fake.selectWorkerArgsForCall)]
	fake.selectWorkerArgsForCall = append(fake.selectWorkerArgsForCall, struct {
		arg1 context.Context
		arg2 worker.Pool
		arg3 db.ContainerOwner
		arg4 worker.ContainerSpec
		arg5 worker.WorkerSpec
		arg6 worker.ContainerPlacementStrategy
		arg7 time.Duration
		arg8 time.Duration
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	stub := fake.SelectWorkerStub
	fakeReturns := fake.selectWorkerReturns
	fake.recordInvocation("SelectWorker", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.selectWorkerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskDelegate) SelectWorkerCallCount() int {
	fake.selectWorkerMutex.RLock()
	defer fake.selectWorkerMutex.RUnlock()
	return len(fake.selectWorkerArgsForCall)
}

func (fake *FakeTaskDelegate) SelectWorkerCalls(stub func(context.Context, worker.Pool, db.ContainerOwner, worker.ContainerSpec, worker.WorkerSpec, worker.ContainerPlacementStrategy, time.Duration, time.Duration) (worker.Client, error)) {
	fake.selectWorkerMutex.Lock()
	defer fake.selectWorkerMutex.Unlock()
	fake.SelectWorkerStub = stub
}

func (fake *FakeTaskDelegate) SelectWorkerArgsForCall(i int) (context.Context, worker.Pool, db.ContainerOwner, worker.ContainerSpec, worker.WorkerSpec, worker.ContainerPlacementStrategy, time.Duration, time.Duration) {
	fake.selectWorkerMutex.RLock()
	defer fake.selectWorkerMutex.RUnlock()
	argsForCall := fake.selectWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8
}

func (fake *FakeTaskDelegate) SelectWorkerReturns(result1 worker.Client, result2 error) {
	fake.selectWorkerMutex.Lock()
	defer fake.selectWorkerMutex.Unlock()
	fake.SelectWorkerStub = nil
	fake.selectWorkerReturns = struct {
		result1 worker.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDelegate) SelectWorkerReturnsOnCall(i int, result1 worker.Client, result2 error) {
	fake.selectWorkerMutex.Lock()
	defer fake.selectWorkerMutex.Unlock()
	fake.SelectWorkerStub = nil
	if fake.selectWorkerReturnsOnCall == nil {
		fake.selectWorkerReturnsOnCall = make(map[int]struct {
			result1 worker.Client
			result2 error
		})
	}
	fake.selectWorkerReturnsOnCall[i] = struct {
		result1 worker.Client
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDelegate) SelectedWorker(arg1 lager.Logger, arg2 string) {
	fake.selectedWorkerMutex.Lock()
	fake.selectedWorkerArgsForCall = append(fake.selectedWorkerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.SelectedWorkerStub
	fake.recordInvocation("SelectedWorker", []interface{}{arg1, arg2})
	fake.selectedWorkerMutex.Unlock()
	if stub != nil {
		fake.SelectedWorkerStub(arg1, arg2)
	}
}

func (fake *FakeTaskDelegate) SelectedWorkerCallCount() int {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	return len(fake.selectedWorkerArgsForCall)
}

func (fake *FakeTaskDelegate) SelectedWorkerCalls(stub func(lager.Logger, string)) {
	fake.selectedWorkerMutex.Lock()
	defer fake.selectedWorkerMutex.Unlock()
	fake.SelectedWorkerStub = stub
}

func (fake *FakeTaskDelegate) SelectedWorkerArgsForCall(i int) (lager.Logger, string) {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	argsForCall := fake.selectedWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskDelegate) SetTaskConfig(arg1 atc.TaskConfig) {
	fake.setTaskConfigMutex.Lock()
	fake.setTaskConfigArgsForCall = append(fake.setTaskConfigArgsForCall, struct {
		arg1 atc.TaskConfig
	}{arg1})
	stub := fake.SetTaskConfigStub
	fake.recordInvocation("SetTaskConfig", []interface{}{arg1})
	fake.setTaskConfigMutex.Unlock()
	if stub != nil {
		fake.SetTaskConfigStub(arg1)
	}
}

func (fake *FakeTaskDelegate) SetTaskConfigCallCount() int {
	fake.setTaskConfigMutex.RLock()
	defer fake.setTaskConfigMutex.RUnlock()
	return len(fake.setTaskConfigArgsForCall)
}

func (fake *FakeTaskDelegate) SetTaskConfigCalls(stub func(atc.TaskConfig)) {
	fake.setTaskConfigMutex.Lock()
	defer fake.setTaskConfigMutex.Unlock()
	fake.SetTaskConfigStub = stub
}

func (fake *FakeTaskDelegate) SetTaskConfigArgsForCall(i int) atc.TaskConfig {
	fake.setTaskConfigMutex.RLock()
	defer fake.setTaskConfigMutex.RUnlock()
	argsForCall := fake.setTaskConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskDelegate) StartSpan(arg1 context.Context, arg2 string, arg3 tracing.Attrs) (context.Context, trace.Span) {
	fake.startSpanMutex.Lock()
	ret, specificReturn := fake.startSpanReturnsOnCall[len(fake.startSpanArgsForCall)]
	fake.startSpanArgsForCall = append(fake.startSpanArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}{arg1, arg2, arg3})
	stub := fake.StartSpanStub
	fakeReturns := fake.startSpanReturns
	fake.recordInvocation("StartSpan", []interface{}{arg1, arg2, arg3})
	fake.startSpanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskDelegate) StartSpanCallCount() int {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	return len(fake.startSpanArgsForCall)
}

func (fake *FakeTaskDelegate) StartSpanCalls(stub func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = stub
}

func (fake *FakeTaskDelegate) StartSpanArgsForCall(i int) (context.Context, string, tracing.Attrs) {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	argsForCall := fake.startSpanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTaskDelegate) StartSpanReturns(result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	fake.startSpanReturns = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeTaskDelegate) StartSpanReturnsOnCall(i int, result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	if fake.startSpanReturnsOnCall == nil {
		fake.startSpanReturnsOnCall = make(map[int]struct {
			result1 context.Context
			result2 trace.Span
		})
	}
	fake.startSpanReturnsOnCall[i] = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeTaskDelegate) Starting(arg1 lager.Logger) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.StartingStub
	fake.recordInvocation("Starting", []interface{}{arg1})
	fake.startingMutex.Unlock()
	if stub != nil {
		fake.StartingStub(arg1)
	}
}

func (fake *FakeTaskDelegate) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakeTaskDelegate) StartingCalls(stub func(lager.Logger)) {
	fake.startingMutex.Lock()
	defer fake.startingMutex.Unlock()
	fake.StartingStub = stub
}

func (fake *FakeTaskDelegate) StartingArgsForCall(i int) lager.Logger {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	argsForCall := fake.startingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTaskDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	stub := fake.StderrStub
	fakeReturns := fake.stderrReturns
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeTaskDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakeTaskDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeTaskDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeTaskDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	stub := fake.StdoutStub
	fakeReturns := fake.stdoutReturns
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeTaskDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeTaskDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeTaskDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeTaskDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	fake.selectWorkerMutex.RLock()
	defer fake.selectWorkerMutex.RUnlock()
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	fake.setTaskConfigMutex.RLock()
	defer fake.setTaskConfigMutex.RUnlock()
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskDelegate) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ exec.TaskDelegate = new(FakeTaskDelegate)
