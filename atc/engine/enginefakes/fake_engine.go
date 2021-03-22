// Code generated by counterfeiter. DO NOT EDIT.
package enginefakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/engine"
)

type FakeEngine struct {
	DrainStub        func(context.Context)
	drainMutex       sync.RWMutex
	drainArgsForCall []struct {
		arg1 context.Context
	}
	NewBuildStub        func(db.Build) engine.Runnable
	newBuildMutex       sync.RWMutex
	newBuildArgsForCall []struct {
		arg1 db.Build
	}
	newBuildReturns struct {
		result1 engine.Runnable
	}
	newBuildReturnsOnCall map[int]struct {
		result1 engine.Runnable
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEngine) Drain(arg1 context.Context) {
	fake.drainMutex.Lock()
	fake.drainArgsForCall = append(fake.drainArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.DrainStub
	fake.recordInvocation("Drain", []interface{}{arg1})
	fake.drainMutex.Unlock()
	if stub != nil {
		fake.DrainStub(arg1)
	}
}

func (fake *FakeEngine) DrainCallCount() int {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	return len(fake.drainArgsForCall)
}

func (fake *FakeEngine) DrainCalls(stub func(context.Context)) {
	fake.drainMutex.Lock()
	defer fake.drainMutex.Unlock()
	fake.DrainStub = stub
}

func (fake *FakeEngine) DrainArgsForCall(i int) context.Context {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	argsForCall := fake.drainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEngine) NewBuild(arg1 db.Build) engine.Runnable {
	fake.newBuildMutex.Lock()
	ret, specificReturn := fake.newBuildReturnsOnCall[len(fake.newBuildArgsForCall)]
	fake.newBuildArgsForCall = append(fake.newBuildArgsForCall, struct {
		arg1 db.Build
	}{arg1})
	stub := fake.NewBuildStub
	fakeReturns := fake.newBuildReturns
	fake.recordInvocation("NewBuild", []interface{}{arg1})
	fake.newBuildMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEngine) NewBuildCallCount() int {
	fake.newBuildMutex.RLock()
	defer fake.newBuildMutex.RUnlock()
	return len(fake.newBuildArgsForCall)
}

func (fake *FakeEngine) NewBuildCalls(stub func(db.Build) engine.Runnable) {
	fake.newBuildMutex.Lock()
	defer fake.newBuildMutex.Unlock()
	fake.NewBuildStub = stub
}

func (fake *FakeEngine) NewBuildArgsForCall(i int) db.Build {
	fake.newBuildMutex.RLock()
	defer fake.newBuildMutex.RUnlock()
	argsForCall := fake.newBuildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEngine) NewBuildReturns(result1 engine.Runnable) {
	fake.newBuildMutex.Lock()
	defer fake.newBuildMutex.Unlock()
	fake.NewBuildStub = nil
	fake.newBuildReturns = struct {
		result1 engine.Runnable
	}{result1}
}

func (fake *FakeEngine) NewBuildReturnsOnCall(i int, result1 engine.Runnable) {
	fake.newBuildMutex.Lock()
	defer fake.newBuildMutex.Unlock()
	fake.NewBuildStub = nil
	if fake.newBuildReturnsOnCall == nil {
		fake.newBuildReturnsOnCall = make(map[int]struct {
			result1 engine.Runnable
		})
	}
	fake.newBuildReturnsOnCall[i] = struct {
		result1 engine.Runnable
	}{result1}
}

func (fake *FakeEngine) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	fake.newBuildMutex.RLock()
	defer fake.newBuildMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEngine) recordInvocation(key string, args []interface{}) {
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

var _ engine.Engine = new(FakeEngine)
