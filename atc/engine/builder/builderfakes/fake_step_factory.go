// Code generated by counterfeiter. DO NOT EDIT.
package builderfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/engine/builder"
	"github.com/concourse/concourse/atc/exec"
)

type FakeStepFactory struct {
	ArtifactInputStepStub        func(atc.Plan, db.Build, exec.BuildStepDelegate) exec.Step
	artifactInputStepMutex       sync.RWMutex
	artifactInputStepArgsForCall []struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.BuildStepDelegate
	}
	artifactInputStepReturns struct {
		result1 exec.Step
	}
	artifactInputStepReturnsOnCall map[int]struct {
		result1 exec.Step
	}
	ArtifactOutputStepStub        func(atc.Plan, db.Build, exec.BuildStepDelegate) exec.Step
	artifactOutputStepMutex       sync.RWMutex
	artifactOutputStepArgsForCall []struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.BuildStepDelegate
	}
	artifactOutputStepReturns struct {
		result1 exec.Step
	}
	artifactOutputStepReturnsOnCall map[int]struct {
		result1 exec.Step
	}
	GetStepStub        func(atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.GetDelegate) exec.Step
	getStepMutex       sync.RWMutex
	getStepArgsForCall []struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.StepMetadata
		arg4 db.ContainerMetadata
		arg5 exec.GetDelegate
	}
	getStepReturns struct {
		result1 exec.Step
	}
	getStepReturnsOnCall map[int]struct {
		result1 exec.Step
	}
	PutStepStub        func(atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.PutDelegate) exec.Step
	putStepMutex       sync.RWMutex
	putStepArgsForCall []struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.StepMetadata
		arg4 db.ContainerMetadata
		arg5 exec.PutDelegate
	}
	putStepReturns struct {
		result1 exec.Step
	}
	putStepReturnsOnCall map[int]struct {
		result1 exec.Step
	}
	TaskStepStub        func(atc.Plan, db.Build, db.ContainerMetadata, exec.TaskDelegate) exec.Step
	taskStepMutex       sync.RWMutex
	taskStepArgsForCall []struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 db.ContainerMetadata
		arg4 exec.TaskDelegate
	}
	taskStepReturns struct {
		result1 exec.Step
	}
	taskStepReturnsOnCall map[int]struct {
		result1 exec.Step
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStepFactory) ArtifactInputStep(arg1 atc.Plan, arg2 db.Build, arg3 exec.BuildStepDelegate) exec.Step {
	fake.artifactInputStepMutex.Lock()
	ret, specificReturn := fake.artifactInputStepReturnsOnCall[len(fake.artifactInputStepArgsForCall)]
	fake.artifactInputStepArgsForCall = append(fake.artifactInputStepArgsForCall, struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.BuildStepDelegate
	}{arg1, arg2, arg3})
	fake.recordInvocation("ArtifactInputStep", []interface{}{arg1, arg2, arg3})
	fake.artifactInputStepMutex.Unlock()
	if fake.ArtifactInputStepStub != nil {
		return fake.ArtifactInputStepStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.artifactInputStepReturns
	return fakeReturns.result1
}

func (fake *FakeStepFactory) ArtifactInputStepCallCount() int {
	fake.artifactInputStepMutex.RLock()
	defer fake.artifactInputStepMutex.RUnlock()
	return len(fake.artifactInputStepArgsForCall)
}

func (fake *FakeStepFactory) ArtifactInputStepCalls(stub func(atc.Plan, db.Build, exec.BuildStepDelegate) exec.Step) {
	fake.artifactInputStepMutex.Lock()
	defer fake.artifactInputStepMutex.Unlock()
	fake.ArtifactInputStepStub = stub
}

func (fake *FakeStepFactory) ArtifactInputStepArgsForCall(i int) (atc.Plan, db.Build, exec.BuildStepDelegate) {
	fake.artifactInputStepMutex.RLock()
	defer fake.artifactInputStepMutex.RUnlock()
	argsForCall := fake.artifactInputStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStepFactory) ArtifactInputStepReturns(result1 exec.Step) {
	fake.artifactInputStepMutex.Lock()
	defer fake.artifactInputStepMutex.Unlock()
	fake.ArtifactInputStepStub = nil
	fake.artifactInputStepReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) ArtifactInputStepReturnsOnCall(i int, result1 exec.Step) {
	fake.artifactInputStepMutex.Lock()
	defer fake.artifactInputStepMutex.Unlock()
	fake.ArtifactInputStepStub = nil
	if fake.artifactInputStepReturnsOnCall == nil {
		fake.artifactInputStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
		})
	}
	fake.artifactInputStepReturnsOnCall[i] = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) ArtifactOutputStep(arg1 atc.Plan, arg2 db.Build, arg3 exec.BuildStepDelegate) exec.Step {
	fake.artifactOutputStepMutex.Lock()
	ret, specificReturn := fake.artifactOutputStepReturnsOnCall[len(fake.artifactOutputStepArgsForCall)]
	fake.artifactOutputStepArgsForCall = append(fake.artifactOutputStepArgsForCall, struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.BuildStepDelegate
	}{arg1, arg2, arg3})
	fake.recordInvocation("ArtifactOutputStep", []interface{}{arg1, arg2, arg3})
	fake.artifactOutputStepMutex.Unlock()
	if fake.ArtifactOutputStepStub != nil {
		return fake.ArtifactOutputStepStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.artifactOutputStepReturns
	return fakeReturns.result1
}

func (fake *FakeStepFactory) ArtifactOutputStepCallCount() int {
	fake.artifactOutputStepMutex.RLock()
	defer fake.artifactOutputStepMutex.RUnlock()
	return len(fake.artifactOutputStepArgsForCall)
}

func (fake *FakeStepFactory) ArtifactOutputStepCalls(stub func(atc.Plan, db.Build, exec.BuildStepDelegate) exec.Step) {
	fake.artifactOutputStepMutex.Lock()
	defer fake.artifactOutputStepMutex.Unlock()
	fake.ArtifactOutputStepStub = stub
}

func (fake *FakeStepFactory) ArtifactOutputStepArgsForCall(i int) (atc.Plan, db.Build, exec.BuildStepDelegate) {
	fake.artifactOutputStepMutex.RLock()
	defer fake.artifactOutputStepMutex.RUnlock()
	argsForCall := fake.artifactOutputStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStepFactory) ArtifactOutputStepReturns(result1 exec.Step) {
	fake.artifactOutputStepMutex.Lock()
	defer fake.artifactOutputStepMutex.Unlock()
	fake.ArtifactOutputStepStub = nil
	fake.artifactOutputStepReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) ArtifactOutputStepReturnsOnCall(i int, result1 exec.Step) {
	fake.artifactOutputStepMutex.Lock()
	defer fake.artifactOutputStepMutex.Unlock()
	fake.ArtifactOutputStepStub = nil
	if fake.artifactOutputStepReturnsOnCall == nil {
		fake.artifactOutputStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
		})
	}
	fake.artifactOutputStepReturnsOnCall[i] = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) GetStep(arg1 atc.Plan, arg2 db.Build, arg3 exec.StepMetadata, arg4 db.ContainerMetadata, arg5 exec.GetDelegate) exec.Step {
	fake.getStepMutex.Lock()
	ret, specificReturn := fake.getStepReturnsOnCall[len(fake.getStepArgsForCall)]
	fake.getStepArgsForCall = append(fake.getStepArgsForCall, struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.StepMetadata
		arg4 db.ContainerMetadata
		arg5 exec.GetDelegate
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("GetStep", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.getStepMutex.Unlock()
	if fake.GetStepStub != nil {
		return fake.GetStepStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getStepReturns
	return fakeReturns.result1
}

func (fake *FakeStepFactory) GetStepCallCount() int {
	fake.getStepMutex.RLock()
	defer fake.getStepMutex.RUnlock()
	return len(fake.getStepArgsForCall)
}

func (fake *FakeStepFactory) GetStepCalls(stub func(atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.GetDelegate) exec.Step) {
	fake.getStepMutex.Lock()
	defer fake.getStepMutex.Unlock()
	fake.GetStepStub = stub
}

func (fake *FakeStepFactory) GetStepArgsForCall(i int) (atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.GetDelegate) {
	fake.getStepMutex.RLock()
	defer fake.getStepMutex.RUnlock()
	argsForCall := fake.getStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeStepFactory) GetStepReturns(result1 exec.Step) {
	fake.getStepMutex.Lock()
	defer fake.getStepMutex.Unlock()
	fake.GetStepStub = nil
	fake.getStepReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) GetStepReturnsOnCall(i int, result1 exec.Step) {
	fake.getStepMutex.Lock()
	defer fake.getStepMutex.Unlock()
	fake.GetStepStub = nil
	if fake.getStepReturnsOnCall == nil {
		fake.getStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
		})
	}
	fake.getStepReturnsOnCall[i] = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) PutStep(arg1 atc.Plan, arg2 db.Build, arg3 exec.StepMetadata, arg4 db.ContainerMetadata, arg5 exec.PutDelegate) exec.Step {
	fake.putStepMutex.Lock()
	ret, specificReturn := fake.putStepReturnsOnCall[len(fake.putStepArgsForCall)]
	fake.putStepArgsForCall = append(fake.putStepArgsForCall, struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 exec.StepMetadata
		arg4 db.ContainerMetadata
		arg5 exec.PutDelegate
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("PutStep", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.putStepMutex.Unlock()
	if fake.PutStepStub != nil {
		return fake.PutStepStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putStepReturns
	return fakeReturns.result1
}

func (fake *FakeStepFactory) PutStepCallCount() int {
	fake.putStepMutex.RLock()
	defer fake.putStepMutex.RUnlock()
	return len(fake.putStepArgsForCall)
}

func (fake *FakeStepFactory) PutStepCalls(stub func(atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.PutDelegate) exec.Step) {
	fake.putStepMutex.Lock()
	defer fake.putStepMutex.Unlock()
	fake.PutStepStub = stub
}

func (fake *FakeStepFactory) PutStepArgsForCall(i int) (atc.Plan, db.Build, exec.StepMetadata, db.ContainerMetadata, exec.PutDelegate) {
	fake.putStepMutex.RLock()
	defer fake.putStepMutex.RUnlock()
	argsForCall := fake.putStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeStepFactory) PutStepReturns(result1 exec.Step) {
	fake.putStepMutex.Lock()
	defer fake.putStepMutex.Unlock()
	fake.PutStepStub = nil
	fake.putStepReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) PutStepReturnsOnCall(i int, result1 exec.Step) {
	fake.putStepMutex.Lock()
	defer fake.putStepMutex.Unlock()
	fake.PutStepStub = nil
	if fake.putStepReturnsOnCall == nil {
		fake.putStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
		})
	}
	fake.putStepReturnsOnCall[i] = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) TaskStep(arg1 atc.Plan, arg2 db.Build, arg3 db.ContainerMetadata, arg4 exec.TaskDelegate) exec.Step {
	fake.taskStepMutex.Lock()
	ret, specificReturn := fake.taskStepReturnsOnCall[len(fake.taskStepArgsForCall)]
	fake.taskStepArgsForCall = append(fake.taskStepArgsForCall, struct {
		arg1 atc.Plan
		arg2 db.Build
		arg3 db.ContainerMetadata
		arg4 exec.TaskDelegate
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("TaskStep", []interface{}{arg1, arg2, arg3, arg4})
	fake.taskStepMutex.Unlock()
	if fake.TaskStepStub != nil {
		return fake.TaskStepStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.taskStepReturns
	return fakeReturns.result1
}

func (fake *FakeStepFactory) TaskStepCallCount() int {
	fake.taskStepMutex.RLock()
	defer fake.taskStepMutex.RUnlock()
	return len(fake.taskStepArgsForCall)
}

func (fake *FakeStepFactory) TaskStepCalls(stub func(atc.Plan, db.Build, db.ContainerMetadata, exec.TaskDelegate) exec.Step) {
	fake.taskStepMutex.Lock()
	defer fake.taskStepMutex.Unlock()
	fake.TaskStepStub = stub
}

func (fake *FakeStepFactory) TaskStepArgsForCall(i int) (atc.Plan, db.Build, db.ContainerMetadata, exec.TaskDelegate) {
	fake.taskStepMutex.RLock()
	defer fake.taskStepMutex.RUnlock()
	argsForCall := fake.taskStepArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeStepFactory) TaskStepReturns(result1 exec.Step) {
	fake.taskStepMutex.Lock()
	defer fake.taskStepMutex.Unlock()
	fake.TaskStepStub = nil
	fake.taskStepReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) TaskStepReturnsOnCall(i int, result1 exec.Step) {
	fake.taskStepMutex.Lock()
	defer fake.taskStepMutex.Unlock()
	fake.TaskStepStub = nil
	if fake.taskStepReturnsOnCall == nil {
		fake.taskStepReturnsOnCall = make(map[int]struct {
			result1 exec.Step
		})
	}
	fake.taskStepReturnsOnCall[i] = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeStepFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.artifactInputStepMutex.RLock()
	defer fake.artifactInputStepMutex.RUnlock()
	fake.artifactOutputStepMutex.RLock()
	defer fake.artifactOutputStepMutex.RUnlock()
	fake.getStepMutex.RLock()
	defer fake.getStepMutex.RUnlock()
	fake.putStepMutex.RLock()
	defer fake.putStepMutex.RUnlock()
	fake.taskStepMutex.RLock()
	defer fake.taskStepMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStepFactory) recordInvocation(key string, args []interface{}) {
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

var _ builder.StepFactory = new(FakeStepFactory)
