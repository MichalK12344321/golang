// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	"lca/internal/pkg/storage"
	"sync"

	"github.com/google/uuid"
)

type FakeStorage struct {
	CreateFileStub        func(uuid.UUID, string, string) (string, error)
	createFileMutex       sync.RWMutex
	createFileArgsForCall []struct {
		arg1 uuid.UUID
		arg2 string
		arg3 string
	}
	createFileReturns struct {
		result1 string
		result2 error
	}
	createFileReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetFileStub        func(uuid.UUID) ([]byte, error)
	getFileMutex       sync.RWMutex
	getFileArgsForCall []struct {
		arg1 uuid.UUID
	}
	getFileReturns struct {
		result1 []byte
		result2 error
	}
	getFileReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ListFilesStub        func(uuid.UUID) ([]string, error)
	listFilesMutex       sync.RWMutex
	listFilesArgsForCall []struct {
		arg1 uuid.UUID
	}
	listFilesReturns struct {
		result1 []string
		result2 error
	}
	listFilesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) CreateFile(arg1 uuid.UUID, arg2 string, arg3 string) (string, error) {
	fake.createFileMutex.Lock()
	ret, specificReturn := fake.createFileReturnsOnCall[len(fake.createFileArgsForCall)]
	fake.createFileArgsForCall = append(fake.createFileArgsForCall, struct {
		arg1 uuid.UUID
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateFileStub
	fakeReturns := fake.createFileReturns
	fake.recordInvocation("CreateFile", []interface{}{arg1, arg2, arg3})
	fake.createFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorage) CreateFileCallCount() int {
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	return len(fake.createFileArgsForCall)
}

func (fake *FakeStorage) CreateFileCalls(stub func(uuid.UUID, string, string) (string, error)) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = stub
}

func (fake *FakeStorage) CreateFileArgsForCall(i int) (uuid.UUID, string, string) {
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	argsForCall := fake.createFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStorage) CreateFileReturns(result1 string, result2 error) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = nil
	fake.createFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) CreateFileReturnsOnCall(i int, result1 string, result2 error) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = nil
	if fake.createFileReturnsOnCall == nil {
		fake.createFileReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createFileReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) GetFile(arg1 uuid.UUID) ([]byte, error) {
	fake.getFileMutex.Lock()
	ret, specificReturn := fake.getFileReturnsOnCall[len(fake.getFileArgsForCall)]
	fake.getFileArgsForCall = append(fake.getFileArgsForCall, struct {
		arg1 uuid.UUID
	}{arg1})
	stub := fake.GetFileStub
	fakeReturns := fake.getFileReturns
	fake.recordInvocation("GetFile", []interface{}{arg1})
	fake.getFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorage) GetFileCallCount() int {
	fake.getFileMutex.RLock()
	defer fake.getFileMutex.RUnlock()
	return len(fake.getFileArgsForCall)
}

func (fake *FakeStorage) GetFileCalls(stub func(uuid.UUID) ([]byte, error)) {
	fake.getFileMutex.Lock()
	defer fake.getFileMutex.Unlock()
	fake.GetFileStub = stub
}

func (fake *FakeStorage) GetFileArgsForCall(i int) uuid.UUID {
	fake.getFileMutex.RLock()
	defer fake.getFileMutex.RUnlock()
	argsForCall := fake.getFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStorage) GetFileReturns(result1 []byte, result2 error) {
	fake.getFileMutex.Lock()
	defer fake.getFileMutex.Unlock()
	fake.GetFileStub = nil
	fake.getFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) GetFileReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getFileMutex.Lock()
	defer fake.getFileMutex.Unlock()
	fake.GetFileStub = nil
	if fake.getFileReturnsOnCall == nil {
		fake.getFileReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getFileReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) ListFiles(arg1 uuid.UUID) ([]string, error) {
	fake.listFilesMutex.Lock()
	ret, specificReturn := fake.listFilesReturnsOnCall[len(fake.listFilesArgsForCall)]
	fake.listFilesArgsForCall = append(fake.listFilesArgsForCall, struct {
		arg1 uuid.UUID
	}{arg1})
	stub := fake.ListFilesStub
	fakeReturns := fake.listFilesReturns
	fake.recordInvocation("ListFiles", []interface{}{arg1})
	fake.listFilesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStorage) ListFilesCallCount() int {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return len(fake.listFilesArgsForCall)
}

func (fake *FakeStorage) ListFilesCalls(stub func(uuid.UUID) ([]string, error)) {
	fake.listFilesMutex.Lock()
	defer fake.listFilesMutex.Unlock()
	fake.ListFilesStub = stub
}

func (fake *FakeStorage) ListFilesArgsForCall(i int) uuid.UUID {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	argsForCall := fake.listFilesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStorage) ListFilesReturns(result1 []string, result2 error) {
	fake.listFilesMutex.Lock()
	defer fake.listFilesMutex.Unlock()
	fake.ListFilesStub = nil
	fake.listFilesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) ListFilesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listFilesMutex.Lock()
	defer fake.listFilesMutex.Unlock()
	fake.ListFilesStub = nil
	if fake.listFilesReturnsOnCall == nil {
		fake.listFilesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listFilesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	fake.getFileMutex.RLock()
	defer fake.getFileMutex.RUnlock()
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
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

var _ storage.Storage = new(FakeStorage)