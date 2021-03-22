// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/exec"
)

type FakeBuildStepDelegateFactory struct {
	BuildStepDelegateStub        func(exec.RunState) exec.BuildStepDelegate
	buildStepDelegateMutex       sync.RWMutex
	buildStepDelegateArgsForCall []struct {
		arg1 exec.RunState
	}
	buildStepDelegateReturns struct {
		result1 exec.BuildStepDelegate
	}
	buildStepDelegateReturnsOnCall map[int]struct {
		result1 exec.BuildStepDelegate
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildStepDelegateFactory) BuildStepDelegate(arg1 exec.RunState) exec.BuildStepDelegate {
	fake.buildStepDelegateMutex.Lock()
	ret, specificReturn := fake.buildStepDelegateReturnsOnCall[len(fake.buildStepDelegateArgsForCall)]
	fake.buildStepDelegateArgsForCall = append(fake.buildStepDelegateArgsForCall, struct {
		arg1 exec.RunState
	}{arg1})
	stub := fake.BuildStepDelegateStub
	fakeReturns := fake.buildStepDelegateReturns
	fake.recordInvocation("BuildStepDelegate", []interface{}{arg1})
	fake.buildStepDelegateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildStepDelegateFactory) BuildStepDelegateCallCount() int {
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	return len(fake.buildStepDelegateArgsForCall)
}

func (fake *FakeBuildStepDelegateFactory) BuildStepDelegateCalls(stub func(exec.RunState) exec.BuildStepDelegate) {
	fake.buildStepDelegateMutex.Lock()
	defer fake.buildStepDelegateMutex.Unlock()
	fake.BuildStepDelegateStub = stub
}

func (fake *FakeBuildStepDelegateFactory) BuildStepDelegateArgsForCall(i int) exec.RunState {
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	argsForCall := fake.buildStepDelegateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildStepDelegateFactory) BuildStepDelegateReturns(result1 exec.BuildStepDelegate) {
	fake.buildStepDelegateMutex.Lock()
	defer fake.buildStepDelegateMutex.Unlock()
	fake.BuildStepDelegateStub = nil
	fake.buildStepDelegateReturns = struct {
		result1 exec.BuildStepDelegate
	}{result1}
}

func (fake *FakeBuildStepDelegateFactory) BuildStepDelegateReturnsOnCall(i int, result1 exec.BuildStepDelegate) {
	fake.buildStepDelegateMutex.Lock()
	defer fake.buildStepDelegateMutex.Unlock()
	fake.BuildStepDelegateStub = nil
	if fake.buildStepDelegateReturnsOnCall == nil {
		fake.buildStepDelegateReturnsOnCall = make(map[int]struct {
			result1 exec.BuildStepDelegate
		})
	}
	fake.buildStepDelegateReturnsOnCall[i] = struct {
		result1 exec.BuildStepDelegate
	}{result1}
}

func (fake *FakeBuildStepDelegateFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildStepDelegateMutex.RLock()
	defer fake.buildStepDelegateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildStepDelegateFactory) recordInvocation(key string, args []interface{}) {
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

var _ exec.BuildStepDelegateFactory = new(FakeBuildStepDelegateFactory)
