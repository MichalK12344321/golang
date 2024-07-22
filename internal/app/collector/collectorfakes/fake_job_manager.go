// Code generated by counterfeiter. DO NOT EDIT.
package collectorfakes

import (
	"lca/internal/app/collector"
	"lca/internal/pkg/broker"
	"sync"

	"github.com/google/uuid"
)

type FakeJobManager struct {
	GetStub        func(uuid.UUID) collector.Job
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 uuid.UUID
	}
	getReturns struct {
		result1 collector.Job
	}
	getReturnsOnCall map[int]struct {
		result1 collector.Job
	}
	StartStub        func(broker.EventPayload) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 broker.EventPayload
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	TerminateStub        func(uuid.UUID) error
	terminateMutex       sync.RWMutex
	terminateArgsForCall []struct {
		arg1 uuid.UUID
	}
	terminateReturns struct {
		result1 error
	}
	terminateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJobManager) Get(arg1 uuid.UUID) collector.Job {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 uuid.UUID
	}{arg1})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeJobManager) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeJobManager) GetCalls(stub func(uuid.UUID) collector.Job) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeJobManager) GetArgsForCall(i int) uuid.UUID {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeJobManager) GetReturns(result1 collector.Job) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 collector.Job
	}{result1}
}

func (fake *FakeJobManager) GetReturnsOnCall(i int, result1 collector.Job) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 collector.Job
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 collector.Job
	}{result1}
}

func (fake *FakeJobManager) Start(arg1 broker.EventPayload) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 broker.EventPayload
	}{arg1})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeJobManager) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeJobManager) StartCalls(stub func(broker.EventPayload) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeJobManager) StartArgsForCall(i int) broker.EventPayload {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeJobManager) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobManager) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobManager) Terminate(arg1 uuid.UUID) error {
	fake.terminateMutex.Lock()
	ret, specificReturn := fake.terminateReturnsOnCall[len(fake.terminateArgsForCall)]
	fake.terminateArgsForCall = append(fake.terminateArgsForCall, struct {
		arg1 uuid.UUID
	}{arg1})
	stub := fake.TerminateStub
	fakeReturns := fake.terminateReturns
	fake.recordInvocation("Terminate", []interface{}{arg1})
	fake.terminateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeJobManager) TerminateCallCount() int {
	fake.terminateMutex.RLock()
	defer fake.terminateMutex.RUnlock()
	return len(fake.terminateArgsForCall)
}

func (fake *FakeJobManager) TerminateCalls(stub func(uuid.UUID) error) {
	fake.terminateMutex.Lock()
	defer fake.terminateMutex.Unlock()
	fake.TerminateStub = stub
}

func (fake *FakeJobManager) TerminateArgsForCall(i int) uuid.UUID {
	fake.terminateMutex.RLock()
	defer fake.terminateMutex.RUnlock()
	argsForCall := fake.terminateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeJobManager) TerminateReturns(result1 error) {
	fake.terminateMutex.Lock()
	defer fake.terminateMutex.Unlock()
	fake.TerminateStub = nil
	fake.terminateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobManager) TerminateReturnsOnCall(i int, result1 error) {
	fake.terminateMutex.Lock()
	defer fake.terminateMutex.Unlock()
	fake.TerminateStub = nil
	if fake.terminateReturnsOnCall == nil {
		fake.terminateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.terminateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.terminateMutex.RLock()
	defer fake.terminateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJobManager) recordInvocation(key string, args []interface{}) {
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

var _ collector.JobManager = new(FakeJobManager)