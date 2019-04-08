// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/worker"
)

type FakeResource struct {
	CheckStub        func(context.Context, atc.Source, atc.Version) ([]atc.Version, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 context.Context
		arg2 atc.Source
		arg3 atc.Version
	}
	checkReturns struct {
		result1 []atc.Version
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 []atc.Version
		result2 error
	}
	GetStub        func(context.Context, worker.Volume, resource.IOConfig, atc.Source, atc.Params, atc.Version) (resource.VersionedSource, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 worker.Volume
		arg3 resource.IOConfig
		arg4 atc.Source
		arg5 atc.Params
		arg6 atc.Version
	}
	getReturns struct {
		result1 resource.VersionedSource
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 resource.VersionedSource
		result2 error
	}
	PutStub        func(context.Context, resource.IOConfig, atc.Source, atc.Params) (resource.VersionedSource, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 context.Context
		arg2 resource.IOConfig
		arg3 atc.Source
		arg4 atc.Params
	}
	putReturns struct {
		result1 resource.VersionedSource
		result2 error
	}
	putReturnsOnCall map[int]struct {
		result1 resource.VersionedSource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResource) Check(arg1 context.Context, arg2 atc.Source, arg3 atc.Version) ([]atc.Version, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 context.Context
		arg2 atc.Source
		arg3 atc.Version
	}{arg1, arg2, arg3})
	fake.recordInvocation("Check", []interface{}{arg1, arg2, arg3})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeResource) CheckCalls(stub func(context.Context, atc.Source, atc.Version) ([]atc.Version, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeResource) CheckArgsForCall(i int) (context.Context, atc.Source, atc.Version) {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeResource) CheckReturns(result1 []atc.Version, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 []atc.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) CheckReturnsOnCall(i int, result1 []atc.Version, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 []atc.Version
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 []atc.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Get(arg1 context.Context, arg2 worker.Volume, arg3 resource.IOConfig, arg4 atc.Source, arg5 atc.Params, arg6 atc.Version) (resource.VersionedSource, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 worker.Volume
		arg3 resource.IOConfig
		arg4 atc.Source
		arg5 atc.Params
		arg6 atc.Version
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeResource) GetCalls(stub func(context.Context, worker.Volume, resource.IOConfig, atc.Source, atc.Params, atc.Version) (resource.VersionedSource, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeResource) GetArgsForCall(i int) (context.Context, worker.Volume, resource.IOConfig, atc.Source, atc.Params, atc.Version) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeResource) GetReturns(result1 resource.VersionedSource, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) GetReturnsOnCall(i int, result1 resource.VersionedSource, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 resource.VersionedSource
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Put(arg1 context.Context, arg2 resource.IOConfig, arg3 atc.Source, arg4 atc.Params) (resource.VersionedSource, error) {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 context.Context
		arg2 resource.IOConfig
		arg3 atc.Source
		arg4 atc.Params
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3, arg4})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResource) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeResource) PutCalls(stub func(context.Context, resource.IOConfig, atc.Source, atc.Params) (resource.VersionedSource, error)) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeResource) PutArgsForCall(i int) (context.Context, resource.IOConfig, atc.Source, atc.Params) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeResource) PutReturns(result1 resource.VersionedSource, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) PutReturnsOnCall(i int, result1 resource.VersionedSource, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 resource.VersionedSource
			result2 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 resource.VersionedSource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResource) recordInvocation(key string, args []interface{}) {
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

var _ resource.Resource = new(FakeResource)
