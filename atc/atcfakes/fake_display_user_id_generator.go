// Code generated by counterfeiter. DO NOT EDIT.
package atcfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
)

type FakeDisplayUserIdGenerator struct {
	DisplayUserIdStub        func(string, string, string, string, string) string
	displayUserIdMutex       sync.RWMutex
	displayUserIdArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	displayUserIdReturns struct {
		result1 string
	}
	displayUserIdReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDisplayUserIdGenerator) DisplayUserId(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) string {
	fake.displayUserIdMutex.Lock()
	ret, specificReturn := fake.displayUserIdReturnsOnCall[len(fake.displayUserIdArgsForCall)]
	fake.displayUserIdArgsForCall = append(fake.displayUserIdArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.DisplayUserIdStub
	fakeReturns := fake.displayUserIdReturns
	fake.recordInvocation("DisplayUserId", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.displayUserIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDisplayUserIdGenerator) DisplayUserIdCallCount() int {
	fake.displayUserIdMutex.RLock()
	defer fake.displayUserIdMutex.RUnlock()
	return len(fake.displayUserIdArgsForCall)
}

func (fake *FakeDisplayUserIdGenerator) DisplayUserIdCalls(stub func(string, string, string, string, string) string) {
	fake.displayUserIdMutex.Lock()
	defer fake.displayUserIdMutex.Unlock()
	fake.DisplayUserIdStub = stub
}

func (fake *FakeDisplayUserIdGenerator) DisplayUserIdArgsForCall(i int) (string, string, string, string, string) {
	fake.displayUserIdMutex.RLock()
	defer fake.displayUserIdMutex.RUnlock()
	argsForCall := fake.displayUserIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeDisplayUserIdGenerator) DisplayUserIdReturns(result1 string) {
	fake.displayUserIdMutex.Lock()
	defer fake.displayUserIdMutex.Unlock()
	fake.DisplayUserIdStub = nil
	fake.displayUserIdReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeDisplayUserIdGenerator) DisplayUserIdReturnsOnCall(i int, result1 string) {
	fake.displayUserIdMutex.Lock()
	defer fake.displayUserIdMutex.Unlock()
	fake.DisplayUserIdStub = nil
	if fake.displayUserIdReturnsOnCall == nil {
		fake.displayUserIdReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.displayUserIdReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeDisplayUserIdGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.displayUserIdMutex.RLock()
	defer fake.displayUserIdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDisplayUserIdGenerator) recordInvocation(key string, args []interface{}) {
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

var _ atc.DisplayUserIdGenerator = new(FakeDisplayUserIdGenerator)
