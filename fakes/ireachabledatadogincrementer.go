// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/naqvijafar91/go-health/v2/checkers"
)

type FakeReachableDatadogIncrementer struct {
	IncrStub        func(name string, tags []string, rate float64) error
	incrMutex       sync.RWMutex
	incrArgsForCall []struct {
		name string
		tags []string
		rate float64
	}
	incrReturns struct {
		result1 error
	}
	incrReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReachableDatadogIncrementer) Incr(name string, tags []string, rate float64) error {
	var tagsCopy []string
	if tags != nil {
		tagsCopy = make([]string, len(tags))
		copy(tagsCopy, tags)
	}
	fake.incrMutex.Lock()
	ret, specificReturn := fake.incrReturnsOnCall[len(fake.incrArgsForCall)]
	fake.incrArgsForCall = append(fake.incrArgsForCall, struct {
		name string
		tags []string
		rate float64
	}{name, tagsCopy, rate})
	fake.recordInvocation("Incr", []interface{}{name, tagsCopy, rate})
	fake.incrMutex.Unlock()
	if fake.IncrStub != nil {
		return fake.IncrStub(name, tags, rate)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.incrReturns.result1
}

func (fake *FakeReachableDatadogIncrementer) IncrCallCount() int {
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	return len(fake.incrArgsForCall)
}

func (fake *FakeReachableDatadogIncrementer) IncrArgsForCall(i int) (string, []string, float64) {
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	return fake.incrArgsForCall[i].name, fake.incrArgsForCall[i].tags, fake.incrArgsForCall[i].rate
}

func (fake *FakeReachableDatadogIncrementer) IncrReturns(result1 error) {
	fake.IncrStub = nil
	fake.incrReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReachableDatadogIncrementer) IncrReturnsOnCall(i int, result1 error) {
	fake.IncrStub = nil
	if fake.incrReturnsOnCall == nil {
		fake.incrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReachableDatadogIncrementer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReachableDatadogIncrementer) recordInvocation(key string, args []interface{}) {
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

var _ checkers.ReachableDatadogIncrementer = new(FakeReachableDatadogIncrementer)
