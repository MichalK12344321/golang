// Code generated by counterfeiter. DO NOT EDIT.
package brokerfakes

import (
	"lca/internal/pkg/broker"
	"sync"
)

type FakePublisherFactoryProvider struct {
	GetPublisherFactoryStub        func() broker.PublisherFactory
	getPublisherFactoryMutex       sync.RWMutex
	getPublisherFactoryArgsForCall []struct {
	}
	getPublisherFactoryReturns struct {
		result1 broker.PublisherFactory
	}
	getPublisherFactoryReturnsOnCall map[int]struct {
		result1 broker.PublisherFactory
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePublisherFactoryProvider) GetPublisherFactory() broker.PublisherFactory {
	fake.getPublisherFactoryMutex.Lock()
	ret, specificReturn := fake.getPublisherFactoryReturnsOnCall[len(fake.getPublisherFactoryArgsForCall)]
	fake.getPublisherFactoryArgsForCall = append(fake.getPublisherFactoryArgsForCall, struct {
	}{})
	stub := fake.GetPublisherFactoryStub
	fakeReturns := fake.getPublisherFactoryReturns
	fake.recordInvocation("GetPublisherFactory", []interface{}{})
	fake.getPublisherFactoryMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublisherFactoryProvider) GetPublisherFactoryCallCount() int {
	fake.getPublisherFactoryMutex.RLock()
	defer fake.getPublisherFactoryMutex.RUnlock()
	return len(fake.getPublisherFactoryArgsForCall)
}

func (fake *FakePublisherFactoryProvider) GetPublisherFactoryCalls(stub func() broker.PublisherFactory) {
	fake.getPublisherFactoryMutex.Lock()
	defer fake.getPublisherFactoryMutex.Unlock()
	fake.GetPublisherFactoryStub = stub
}

func (fake *FakePublisherFactoryProvider) GetPublisherFactoryReturns(result1 broker.PublisherFactory) {
	fake.getPublisherFactoryMutex.Lock()
	defer fake.getPublisherFactoryMutex.Unlock()
	fake.GetPublisherFactoryStub = nil
	fake.getPublisherFactoryReturns = struct {
		result1 broker.PublisherFactory
	}{result1}
}

func (fake *FakePublisherFactoryProvider) GetPublisherFactoryReturnsOnCall(i int, result1 broker.PublisherFactory) {
	fake.getPublisherFactoryMutex.Lock()
	defer fake.getPublisherFactoryMutex.Unlock()
	fake.GetPublisherFactoryStub = nil
	if fake.getPublisherFactoryReturnsOnCall == nil {
		fake.getPublisherFactoryReturnsOnCall = make(map[int]struct {
			result1 broker.PublisherFactory
		})
	}
	fake.getPublisherFactoryReturnsOnCall[i] = struct {
		result1 broker.PublisherFactory
	}{result1}
}

func (fake *FakePublisherFactoryProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPublisherFactoryMutex.RLock()
	defer fake.getPublisherFactoryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePublisherFactoryProvider) recordInvocation(key string, args []interface{}) {
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

var _ broker.PublisherFactoryProvider = new(FakePublisherFactoryProvider)
