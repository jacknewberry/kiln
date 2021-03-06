// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"path/filepath"
	"sync"
)

type Filesystem struct {
	CreateStub        func(path string) (io.WriteCloser, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		path string
	}
	createReturns struct {
		result1 io.WriteCloser
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 io.WriteCloser
		result2 error
	}
	OpenStub        func(path string) (io.ReadCloser, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		path string
	}
	openReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	openReturnsOnCall map[int]struct {
		result1 io.ReadCloser
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
	RemoveStub        func(path string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		path string
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Filesystem) Create(path string) (io.WriteCloser, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Create", []interface{}{path})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *Filesystem) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Filesystem) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].path
}

func (fake *Filesystem) CreateReturns(result1 io.WriteCloser, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *Filesystem) CreateReturnsOnCall(i int, result1 io.WriteCloser, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 io.WriteCloser
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *Filesystem) Open(path string) (io.ReadCloser, error) {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Open", []interface{}{path})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(path)
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
	return fake.openArgsForCall[i].path
}

func (fake *Filesystem) OpenReturns(result1 io.ReadCloser, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *Filesystem) OpenReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 io.ReadCloser
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

func (fake *Filesystem) Remove(path string) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Remove", []interface{}{path})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeReturns.result1
}

func (fake *Filesystem) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *Filesystem) RemoveArgsForCall(i int) string {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].path
}

func (fake *Filesystem) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Filesystem) RemoveReturnsOnCall(i int, result1 error) {
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

func (fake *Filesystem) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
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
