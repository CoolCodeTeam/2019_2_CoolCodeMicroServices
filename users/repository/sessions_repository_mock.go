// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"sync"
)

var (
	lockSessionRepositoryMockContains sync.RWMutex
	lockSessionRepositoryMockGetID    sync.RWMutex
	lockSessionRepositoryMockPut      sync.RWMutex
	lockSessionRepositoryMockRemove   sync.RWMutex
)

// Ensure, that SessionRepositoryMock does implement SessionRepository.
// If this is not the case, regenerate this file with moq.
var _ SessionRepository = &SessionRepositoryMock{}

// SessionRepositoryMock is a mock implementation of SessionRepository.
//
//     func TestSomethingThatUsesSessionRepository(t *testing.T) {
//
//         // make and configure a mocked SessionRepository
//         mockedSessionRepository := &SessionRepositoryMock{
//             ContainsFunc: func(session string) bool {
// 	               panic("mock out the Contains method")
//             },
//             GetIDFunc: func(session string) (uint64, error) {
// 	               panic("mock out the GetID method")
//             },
//             PutFunc: func(session string, id uint64) error {
// 	               panic("mock out the Put method")
//             },
//             RemoveFunc: func(session string) error {
// 	               panic("mock out the Remove method")
//             },
//         }
//
//         // use mockedSessionRepository in code that requires SessionRepository
//         // and then make assertions.
//
//     }
type SessionRepositoryMock struct {
	// ContainsFunc mocks the Contains method.
	ContainsFunc func(session string) bool

	// GetIDFunc mocks the GetID method.
	GetIDFunc func(session string) (uint64, error)

	// PutFunc mocks the Put method.
	PutFunc func(session string, id uint64) error

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(session string) error

	// calls tracks calls to the methods.
	calls struct {
		// Contains holds details about calls to the Contains method.
		Contains []struct {
			// Session is the session argument value.
			Session string
		}
		// GetID holds details about calls to the GetID method.
		GetID []struct {
			// Session is the session argument value.
			Session string
		}
		// Put holds details about calls to the Put method.
		Put []struct {
			// Session is the session argument value.
			Session string
			// ID is the id argument value.
			ID uint64
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// Session is the session argument value.
			Session string
		}
	}
}

// Contains calls ContainsFunc.
func (mock *SessionRepositoryMock) Contains(session string) bool {
	if mock.ContainsFunc == nil {
		panic("SessionRepositoryMock.ContainsFunc: method is nil but SessionRepository.Contains was just called")
	}
	callInfo := struct {
		Session string
	}{
		Session: session,
	}
	lockSessionRepositoryMockContains.Lock()
	mock.calls.Contains = append(mock.calls.Contains, callInfo)
	lockSessionRepositoryMockContains.Unlock()
	return mock.ContainsFunc(session)
}

// ContainsCalls gets all the calls that were made to Contains.
// Check the length with:
//     len(mockedSessionRepository.ContainsCalls())
func (mock *SessionRepositoryMock) ContainsCalls() []struct {
	Session string
} {
	var calls []struct {
		Session string
	}
	lockSessionRepositoryMockContains.RLock()
	calls = mock.calls.Contains
	lockSessionRepositoryMockContains.RUnlock()
	return calls
}

// GetID calls GetIDFunc.
func (mock *SessionRepositoryMock) GetID(session string) (uint64, error) {
	if mock.GetIDFunc == nil {
		panic("SessionRepositoryMock.GetIDFunc: method is nil but SessionRepository.GetID was just called")
	}
	callInfo := struct {
		Session string
	}{
		Session: session,
	}
	lockSessionRepositoryMockGetID.Lock()
	mock.calls.GetID = append(mock.calls.GetID, callInfo)
	lockSessionRepositoryMockGetID.Unlock()
	return mock.GetIDFunc(session)
}

// GetIDCalls gets all the calls that were made to GetID.
// Check the length with:
//     len(mockedSessionRepository.GetIDCalls())
func (mock *SessionRepositoryMock) GetIDCalls() []struct {
	Session string
} {
	var calls []struct {
		Session string
	}
	lockSessionRepositoryMockGetID.RLock()
	calls = mock.calls.GetID
	lockSessionRepositoryMockGetID.RUnlock()
	return calls
}

// Put calls PutFunc.
func (mock *SessionRepositoryMock) Put(session string, id uint64) error {
	if mock.PutFunc == nil {
		panic("SessionRepositoryMock.PutFunc: method is nil but SessionRepository.Put was just called")
	}
	callInfo := struct {
		Session string
		ID      uint64
	}{
		Session: session,
		ID:      id,
	}
	lockSessionRepositoryMockPut.Lock()
	mock.calls.Put = append(mock.calls.Put, callInfo)
	lockSessionRepositoryMockPut.Unlock()
	return mock.PutFunc(session, id)
}

// PutCalls gets all the calls that were made to Put.
// Check the length with:
//     len(mockedSessionRepository.PutCalls())
func (mock *SessionRepositoryMock) PutCalls() []struct {
	Session string
	ID      uint64
} {
	var calls []struct {
		Session string
		ID      uint64
	}
	lockSessionRepositoryMockPut.RLock()
	calls = mock.calls.Put
	lockSessionRepositoryMockPut.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *SessionRepositoryMock) Remove(session string) error {
	if mock.RemoveFunc == nil {
		panic("SessionRepositoryMock.RemoveFunc: method is nil but SessionRepository.Remove was just called")
	}
	callInfo := struct {
		Session string
	}{
		Session: session,
	}
	lockSessionRepositoryMockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	lockSessionRepositoryMockRemove.Unlock()
	return mock.RemoveFunc(session)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedSessionRepository.RemoveCalls())
func (mock *SessionRepositoryMock) RemoveCalls() []struct {
	Session string
} {
	var calls []struct {
		Session string
	}
	lockSessionRepositoryMockRemove.RLock()
	calls = mock.calls.Remove
	lockSessionRepositoryMockRemove.RUnlock()
	return calls
}