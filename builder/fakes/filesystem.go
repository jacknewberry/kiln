// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"path/filepath"
	"sync"
)

type Filesystem struct {
	OpenStub        func(name string) (io.ReadWriteCloser, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		name string
	}
	openReturns struct {
		result1 io.ReadWriteCloser
		result2 error
	}
	openReturnsOnCall map[int]struct {
		result1 io.ReadWriteCloser
		result2 error
	}
	WalkStub        func(root string, walkFn filepath.WalkFunc) error
	walkMutex       sync.RWMutex
	walkArgsForCall []struct {
		root   string
		walkFn filepath.WalkFunc
	}
	walkReturns struct {
		result1 error
	}
	walkReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Filesystem) Open(name string) (io.ReadWriteCloser, error) {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Open", []interface{}{name})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openReturns.result1, fake.openReturns.result2
}

func (fake *Filesystem) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *Filesystem) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].name
}

func (fake *Filesystem) OpenReturns(result1 io.ReadWriteCloser, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 io.ReadWriteCloser
		result2 error
	}{result1, result2}
}

func (fake *Filesystem) OpenReturnsOnCall(i int, result1 io.ReadWriteCloser, result2 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 io.ReadWriteCloser
			result2 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 io.ReadWriteCloser
		result2 error
	}{result1, result2}
}

func (fake *Filesystem) Walk(root string, walkFn filepath.WalkFunc) error {
	fake.walkMutex.Lock()
	ret, specificReturn := fake.walkReturnsOnCall[len(fake.walkArgsForCall)]
	fake.walkArgsForCall = append(fake.walkArgsForCall, struct {
		root   string
		walkFn filepath.WalkFunc
	}{root, walkFn})
	fake.recordInvocation("Walk", []interface{}{root, walkFn})
	fake.walkMutex.Unlock()
	if fake.WalkStub != nil {
		return fake.WalkStub(root, walkFn)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.walkReturns.result1
}

func (fake *Filesystem) WalkCallCount() int {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return len(fake.walkArgsForCall)
}

func (fake *Filesystem) WalkArgsForCall(i int) (string, filepath.WalkFunc) {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return fake.walkArgsForCall[i].root, fake.walkArgsForCall[i].walkFn
}

func (fake *Filesystem) WalkReturns(result1 error) {
	fake.WalkStub = nil
	fake.walkReturns = struct {
		result1 error
	}{result1}
}

func (fake *Filesystem) WalkReturnsOnCall(i int, result1 error) {
	fake.WalkStub = nil
	if fake.walkReturnsOnCall == nil {
		fake.walkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.walkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Filesystem) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Filesystem) recordInvocation(key string, args []interface{}) {
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
