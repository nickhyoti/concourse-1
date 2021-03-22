// Code generated by counterfeiter. DO NOT EDIT.
package runtimefakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/worker/runtime"
	"github.com/containerd/containerd"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeNetwork struct {
	AddStub        func(context.Context, containerd.Task) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 context.Context
		arg2 containerd.Task
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveStub        func(context.Context, containerd.Task) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 context.Context
		arg2 containerd.Task
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	SetupMountsStub        func(string) ([]specs.Mount, error)
	setupMountsMutex       sync.RWMutex
	setupMountsArgsForCall []struct {
		arg1 string
	}
	setupMountsReturns struct {
		result1 []specs.Mount
		result2 error
	}
	setupMountsReturnsOnCall map[int]struct {
		result1 []specs.Mount
		result2 error
	}
	SetupRestrictedNetworksStub        func() error
	setupRestrictedNetworksMutex       sync.RWMutex
	setupRestrictedNetworksArgsForCall []struct {
	}
	setupRestrictedNetworksReturns struct {
		result1 error
	}
	setupRestrictedNetworksReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetwork) Add(arg1 context.Context, arg2 containerd.Task) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 context.Context
		arg2 containerd.Task
	}{arg1, arg2})
	stub := fake.AddStub
	fakeReturns := fake.addReturns
	fake.recordInvocation("Add", []interface{}{arg1, arg2})
	fake.addMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNetwork) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeNetwork) AddCalls(stub func(context.Context, containerd.Task) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeNetwork) AddArgsForCall(i int) (context.Context, containerd.Task) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetwork) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetwork) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetwork) Remove(arg1 context.Context, arg2 containerd.Task) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 context.Context
		arg2 containerd.Task
	}{arg1, arg2})
	stub := fake.RemoveStub
	fakeReturns := fake.removeReturns
	fake.recordInvocation("Remove", []interface{}{arg1, arg2})
	fake.removeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNetwork) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeNetwork) RemoveCalls(stub func(context.Context, containerd.Task) error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeNetwork) RemoveArgsForCall(i int) (context.Context, containerd.Task) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetwork) RemoveReturns(result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetwork) RemoveReturnsOnCall(i int, result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetwork) SetupMounts(arg1 string) ([]specs.Mount, error) {
	fake.setupMountsMutex.Lock()
	ret, specificReturn := fake.setupMountsReturnsOnCall[len(fake.setupMountsArgsForCall)]
	fake.setupMountsArgsForCall = append(fake.setupMountsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetupMountsStub
	fakeReturns := fake.setupMountsReturns
	fake.recordInvocation("SetupMounts", []interface{}{arg1})
	fake.setupMountsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNetwork) SetupMountsCallCount() int {
	fake.setupMountsMutex.RLock()
	defer fake.setupMountsMutex.RUnlock()
	return len(fake.setupMountsArgsForCall)
}

func (fake *FakeNetwork) SetupMountsCalls(stub func(string) ([]specs.Mount, error)) {
	fake.setupMountsMutex.Lock()
	defer fake.setupMountsMutex.Unlock()
	fake.SetupMountsStub = stub
}

func (fake *FakeNetwork) SetupMountsArgsForCall(i int) string {
	fake.setupMountsMutex.RLock()
	defer fake.setupMountsMutex.RUnlock()
	argsForCall := fake.setupMountsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNetwork) SetupMountsReturns(result1 []specs.Mount, result2 error) {
	fake.setupMountsMutex.Lock()
	defer fake.setupMountsMutex.Unlock()
	fake.SetupMountsStub = nil
	fake.setupMountsReturns = struct {
		result1 []specs.Mount
		result2 error
	}{result1, result2}
}

func (fake *FakeNetwork) SetupMountsReturnsOnCall(i int, result1 []specs.Mount, result2 error) {
	fake.setupMountsMutex.Lock()
	defer fake.setupMountsMutex.Unlock()
	fake.SetupMountsStub = nil
	if fake.setupMountsReturnsOnCall == nil {
		fake.setupMountsReturnsOnCall = make(map[int]struct {
			result1 []specs.Mount
			result2 error
		})
	}
	fake.setupMountsReturnsOnCall[i] = struct {
		result1 []specs.Mount
		result2 error
	}{result1, result2}
}

func (fake *FakeNetwork) SetupRestrictedNetworks() error {
	fake.setupRestrictedNetworksMutex.Lock()
	ret, specificReturn := fake.setupRestrictedNetworksReturnsOnCall[len(fake.setupRestrictedNetworksArgsForCall)]
	fake.setupRestrictedNetworksArgsForCall = append(fake.setupRestrictedNetworksArgsForCall, struct {
	}{})
	stub := fake.SetupRestrictedNetworksStub
	fakeReturns := fake.setupRestrictedNetworksReturns
	fake.recordInvocation("SetupRestrictedNetworks", []interface{}{})
	fake.setupRestrictedNetworksMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNetwork) SetupRestrictedNetworksCallCount() int {
	fake.setupRestrictedNetworksMutex.RLock()
	defer fake.setupRestrictedNetworksMutex.RUnlock()
	return len(fake.setupRestrictedNetworksArgsForCall)
}

func (fake *FakeNetwork) SetupRestrictedNetworksCalls(stub func() error) {
	fake.setupRestrictedNetworksMutex.Lock()
	defer fake.setupRestrictedNetworksMutex.Unlock()
	fake.SetupRestrictedNetworksStub = stub
}

func (fake *FakeNetwork) SetupRestrictedNetworksReturns(result1 error) {
	fake.setupRestrictedNetworksMutex.Lock()
	defer fake.setupRestrictedNetworksMutex.Unlock()
	fake.SetupRestrictedNetworksStub = nil
	fake.setupRestrictedNetworksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetwork) SetupRestrictedNetworksReturnsOnCall(i int, result1 error) {
	fake.setupRestrictedNetworksMutex.Lock()
	defer fake.setupRestrictedNetworksMutex.Unlock()
	fake.SetupRestrictedNetworksStub = nil
	if fake.setupRestrictedNetworksReturnsOnCall == nil {
		fake.setupRestrictedNetworksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupRestrictedNetworksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetwork) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.setupMountsMutex.RLock()
	defer fake.setupMountsMutex.RUnlock()
	fake.setupRestrictedNetworksMutex.RLock()
	defer fake.setupRestrictedNetworksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetwork) recordInvocation(key string, args []interface{}) {
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

var _ runtime.Network = new(FakeNetwork)
