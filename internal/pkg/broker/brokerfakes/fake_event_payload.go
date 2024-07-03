// Code generated by counterfeiter. DO NOT EDIT.
package brokerfakes

import (
	"lca/internal/pkg/broker"
	"sync"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type FakeEventPayload struct {
	AcknowledgeStub        func(...error) error
	acknowledgeMutex       sync.RWMutex
	acknowledgeArgsForCall []struct {
		arg1 []error
	}
	acknowledgeReturns struct {
		result1 error
	}
	acknowledgeReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func() *rabbitmq.Delivery
	getMutex       sync.RWMutex
	getArgsForCall []struct {
	}
	getReturns struct {
		result1 *rabbitmq.Delivery
	}
	getReturnsOnCall map[int]struct {
		result1 *rabbitmq.Delivery
	}
	ParseBodyStub        func(any, broker.ParseBodyFn) error
	parseBodyMutex       sync.RWMutex
	parseBodyArgsForCall []struct {
		arg1 any
		arg2 broker.ParseBodyFn
	}
	parseBodyReturns struct {
		result1 error
	}
	parseBodyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventPayload) Acknowledge(arg1 ...error) error {
	fake.acknowledgeMutex.Lock()
	ret, specificReturn := fake.acknowledgeReturnsOnCall[len(fake.acknowledgeArgsForCall)]
	fake.acknowledgeArgsForCall = append(fake.acknowledgeArgsForCall, struct {
		arg1 []error
	}{arg1})
	stub := fake.AcknowledgeStub
	fakeReturns := fake.acknowledgeReturns
	fake.recordInvocation("Acknowledge", []interface{}{arg1})
	fake.acknowledgeMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEventPayload) AcknowledgeCallCount() int {
	fake.acknowledgeMutex.RLock()
	defer fake.acknowledgeMutex.RUnlock()
	return len(fake.acknowledgeArgsForCall)
}

func (fake *FakeEventPayload) AcknowledgeCalls(stub func(...error) error) {
	fake.acknowledgeMutex.Lock()
	defer fake.acknowledgeMutex.Unlock()
	fake.AcknowledgeStub = stub
}

func (fake *FakeEventPayload) AcknowledgeArgsForCall(i int) []error {
	fake.acknowledgeMutex.RLock()
	defer fake.acknowledgeMutex.RUnlock()
	argsForCall := fake.acknowledgeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventPayload) AcknowledgeReturns(result1 error) {
	fake.acknowledgeMutex.Lock()
	defer fake.acknowledgeMutex.Unlock()
	fake.AcknowledgeStub = nil
	fake.acknowledgeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventPayload) AcknowledgeReturnsOnCall(i int, result1 error) {
	fake.acknowledgeMutex.Lock()
	defer fake.acknowledgeMutex.Unlock()
	fake.AcknowledgeStub = nil
	if fake.acknowledgeReturnsOnCall == nil {
		fake.acknowledgeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.acknowledgeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventPayload) Get() *rabbitmq.Delivery {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
	}{})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEventPayload) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeEventPayload) GetCalls(stub func() *rabbitmq.Delivery) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeEventPayload) GetReturns(result1 *rabbitmq.Delivery) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *rabbitmq.Delivery
	}{result1}
}

func (fake *FakeEventPayload) GetReturnsOnCall(i int, result1 *rabbitmq.Delivery) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *rabbitmq.Delivery
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *rabbitmq.Delivery
	}{result1}
}

func (fake *FakeEventPayload) ParseBody(arg1 any, arg2 broker.ParseBodyFn) error {
	fake.parseBodyMutex.Lock()
	ret, specificReturn := fake.parseBodyReturnsOnCall[len(fake.parseBodyArgsForCall)]
	fake.parseBodyArgsForCall = append(fake.parseBodyArgsForCall, struct {
		arg1 any
		arg2 broker.ParseBodyFn
	}{arg1, arg2})
	stub := fake.ParseBodyStub
	fakeReturns := fake.parseBodyReturns
	fake.recordInvocation("ParseBody", []interface{}{arg1, arg2})
	fake.parseBodyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEventPayload) ParseBodyCallCount() int {
	fake.parseBodyMutex.RLock()
	defer fake.parseBodyMutex.RUnlock()
	return len(fake.parseBodyArgsForCall)
}

func (fake *FakeEventPayload) ParseBodyCalls(stub func(any, broker.ParseBodyFn) error) {
	fake.parseBodyMutex.Lock()
	defer fake.parseBodyMutex.Unlock()
	fake.ParseBodyStub = stub
}

func (fake *FakeEventPayload) ParseBodyArgsForCall(i int) (any, broker.ParseBodyFn) {
	fake.parseBodyMutex.RLock()
	defer fake.parseBodyMutex.RUnlock()
	argsForCall := fake.parseBodyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventPayload) ParseBodyReturns(result1 error) {
	fake.parseBodyMutex.Lock()
	defer fake.parseBodyMutex.Unlock()
	fake.ParseBodyStub = nil
	fake.parseBodyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventPayload) ParseBodyReturnsOnCall(i int, result1 error) {
	fake.parseBodyMutex.Lock()
	defer fake.parseBodyMutex.Unlock()
	fake.ParseBodyStub = nil
	if fake.parseBodyReturnsOnCall == nil {
		fake.parseBodyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.parseBodyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventPayload) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acknowledgeMutex.RLock()
	defer fake.acknowledgeMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.parseBodyMutex.RLock()
	defer fake.parseBodyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventPayload) recordInvocation(key string, args []interface{}) {
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

var _ broker.EventPayload = new(FakeEventPayload)