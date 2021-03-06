// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package grpc_utils

import (
	"context"
	"google.golang.org/grpc"
	"sync"
)

var (
	lockChatsServiceClientMockCheckChannelPermission sync.RWMutex
	lockChatsServiceClientMockCheckChatPermission    sync.RWMutex
	lockChatsServiceClientMockContains               sync.RWMutex
	lockChatsServiceClientMockCreateChannel          sync.RWMutex
	lockChatsServiceClientMockCreateWorkspace        sync.RWMutex
	lockChatsServiceClientMockDeleteChannel          sync.RWMutex
	lockChatsServiceClientMockDeleteChat             sync.RWMutex
	lockChatsServiceClientMockDeleteWorkspace        sync.RWMutex
	lockChatsServiceClientMockEditChannel            sync.RWMutex
	lockChatsServiceClientMockEditWorkspace          sync.RWMutex
	lockChatsServiceClientMockGetChannelByID         sync.RWMutex
	lockChatsServiceClientMockGetChatByID            sync.RWMutex
	lockChatsServiceClientMockGetChatsByUserID       sync.RWMutex
	lockChatsServiceClientMockGetWorkspaceByID       sync.RWMutex
	lockChatsServiceClientMockGetWorkspacesByUserID  sync.RWMutex
	lockChatsServiceClientMockLogoutFromChannel      sync.RWMutex
	lockChatsServiceClientMockLogoutFromWorkspace    sync.RWMutex
	lockChatsServiceClientMockPutChat                sync.RWMutex
)

// Ensure, that ChatsServiceClientMock does implement ChatsServiceClient.
// If this is not the case, regenerate this file with moq.
var _ ChatsServiceClient = &ChatsServiceClientMock{}

// ChatsServiceClientMock is a mock implementation of ChatsServiceClient.
//
//     func TestSomethingThatUsesChatsServiceClient(t *testing.T) {
//
//         // make and configure a mocked ChatsServiceClient
//         mockedChatsServiceClient := &ChatsServiceClientMock{
//             CheckChannelPermissionFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the CheckChannelPermission method")
//             },
//             CheckChatPermissionFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the CheckChatPermission method")
//             },
//             ContainsFunc: func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the Contains method")
//             },
//             CreateChannelFunc: func(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the CreateChannel method")
//             },
//             CreateWorkspaceFunc: func(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the CreateWorkspace method")
//             },
//             DeleteChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the DeleteChannel method")
//             },
//             DeleteChatFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the DeleteChat method")
//             },
//             DeleteWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the DeleteWorkspace method")
//             },
//             EditChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the EditChannel method")
//             },
//             EditWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the EditWorkspace method")
//             },
//             GetChannelByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the GetChannelByID method")
//             },
//             GetChatByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the GetChatByID method")
//             },
//             GetChatsByUserIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the GetChatsByUserID method")
//             },
//             GetWorkspaceByIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the GetWorkspaceByID method")
//             },
//             GetWorkspacesByUserIDFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the GetWorkspacesByUserID method")
//             },
//             LogoutFromChannelFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the LogoutFromChannel method")
//             },
//             LogoutFromWorkspaceFunc: func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
// 	               panic("mock out the LogoutFromWorkspace method")
//             },
//             PutChatFunc: func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*ResponseMessage, error) {
// 	               panic("mock out the PutChat method")
//             },
//         }
//
//         // use mockedChatsServiceClient in code that requires ChatsServiceClient
//         // and then make assertions.
//
//     }
type ChatsServiceClientMock struct {
	// CheckChannelPermissionFunc mocks the CheckChannelPermission method.
	CheckChannelPermissionFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// CheckChatPermissionFunc mocks the CheckChatPermission method.
	CheckChatPermissionFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// ContainsFunc mocks the Contains method.
	ContainsFunc func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*EmptyChats, error)

	// CreateChannelFunc mocks the CreateChannel method.
	CreateChannelFunc func(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*ResponseMessage, error)

	// CreateWorkspaceFunc mocks the CreateWorkspace method.
	CreateWorkspaceFunc func(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (*ResponseMessage, error)

	// DeleteChannelFunc mocks the DeleteChannel method.
	DeleteChannelFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// DeleteChatFunc mocks the DeleteChat method.
	DeleteChatFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// DeleteWorkspaceFunc mocks the DeleteWorkspace method.
	DeleteWorkspaceFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// EditChannelFunc mocks the EditChannel method.
	EditChannelFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// EditWorkspaceFunc mocks the EditWorkspace method.
	EditWorkspaceFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// GetChannelByIDFunc mocks the GetChannelByID method.
	GetChannelByIDFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// GetChatByIDFunc mocks the GetChatByID method.
	GetChatByIDFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// GetChatsByUserIDFunc mocks the GetChatsByUserID method.
	GetChatsByUserIDFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// GetWorkspaceByIDFunc mocks the GetWorkspaceByID method.
	GetWorkspaceByIDFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// GetWorkspacesByUserIDFunc mocks the GetWorkspacesByUserID method.
	GetWorkspacesByUserIDFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)

	// LogoutFromChannelFunc mocks the LogoutFromChannel method.
	LogoutFromChannelFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// LogoutFromWorkspaceFunc mocks the LogoutFromWorkspace method.
	LogoutFromWorkspaceFunc func(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error)

	// PutChatFunc mocks the PutChat method.
	PutChatFunc func(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*ResponseMessage, error)

	// calls tracks calls to the methods.
	calls struct {
		// CheckChannelPermission holds details about calls to the CheckChannelPermission method.
		CheckChannelPermission []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// CheckChatPermission holds details about calls to the CheckChatPermission method.
		CheckChatPermission []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Contains holds details about calls to the Contains method.
		Contains []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *Chat
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// CreateChannel holds details about calls to the CreateChannel method.
		CreateChannel []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *Channel
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// CreateWorkspace holds details about calls to the CreateWorkspace method.
		CreateWorkspace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *Workspace
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// DeleteChannel holds details about calls to the DeleteChannel method.
		DeleteChannel []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// DeleteChat holds details about calls to the DeleteChat method.
		DeleteChat []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// DeleteWorkspace holds details about calls to the DeleteWorkspace method.
		DeleteWorkspace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// EditChannel holds details about calls to the EditChannel method.
		EditChannel []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// EditWorkspace holds details about calls to the EditWorkspace method.
		EditWorkspace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// GetChannelByID holds details about calls to the GetChannelByID method.
		GetChannelByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// GetChatByID holds details about calls to the GetChatByID method.
		GetChatByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// GetChatsByUserID holds details about calls to the GetChatsByUserID method.
		GetChatsByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// GetWorkspaceByID holds details about calls to the GetWorkspaceByID method.
		GetWorkspaceByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// GetWorkspacesByUserID holds details about calls to the GetWorkspacesByUserID method.
		GetWorkspacesByUserID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// LogoutFromChannel holds details about calls to the LogoutFromChannel method.
		LogoutFromChannel []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// LogoutFromWorkspace holds details about calls to the LogoutFromWorkspace method.
		LogoutFromWorkspace []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RequestMessage
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// PutChat holds details about calls to the PutChat method.
		PutChat []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *Chat
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
	}
}

// CheckChannelPermission calls CheckChannelPermissionFunc.
func (mock *ChatsServiceClientMock) CheckChannelPermission(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.CheckChannelPermissionFunc == nil {
		panic("ChatsServiceClientMock.CheckChannelPermissionFunc: method is nil but ChatsServiceClient.CheckChannelPermission was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockCheckChannelPermission.Lock()
	mock.calls.CheckChannelPermission = append(mock.calls.CheckChannelPermission, callInfo)
	lockChatsServiceClientMockCheckChannelPermission.Unlock()
	return mock.CheckChannelPermissionFunc(ctx, in, opts...)
}

// CheckChannelPermissionCalls gets all the calls that were made to CheckChannelPermission.
// Check the length with:
//     len(mockedChatsServiceClient.CheckChannelPermissionCalls())
func (mock *ChatsServiceClientMock) CheckChannelPermissionCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockCheckChannelPermission.RLock()
	calls = mock.calls.CheckChannelPermission
	lockChatsServiceClientMockCheckChannelPermission.RUnlock()
	return calls
}

// CheckChatPermission calls CheckChatPermissionFunc.
func (mock *ChatsServiceClientMock) CheckChatPermission(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.CheckChatPermissionFunc == nil {
		panic("ChatsServiceClientMock.CheckChatPermissionFunc: method is nil but ChatsServiceClient.CheckChatPermission was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockCheckChatPermission.Lock()
	mock.calls.CheckChatPermission = append(mock.calls.CheckChatPermission, callInfo)
	lockChatsServiceClientMockCheckChatPermission.Unlock()
	return mock.CheckChatPermissionFunc(ctx, in, opts...)
}

// CheckChatPermissionCalls gets all the calls that were made to CheckChatPermission.
// Check the length with:
//     len(mockedChatsServiceClient.CheckChatPermissionCalls())
func (mock *ChatsServiceClientMock) CheckChatPermissionCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockCheckChatPermission.RLock()
	calls = mock.calls.CheckChatPermission
	lockChatsServiceClientMockCheckChatPermission.RUnlock()
	return calls
}

// Contains calls ContainsFunc.
func (mock *ChatsServiceClientMock) Contains(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.ContainsFunc == nil {
		panic("ChatsServiceClientMock.ContainsFunc: method is nil but ChatsServiceClient.Contains was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *Chat
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockContains.Lock()
	mock.calls.Contains = append(mock.calls.Contains, callInfo)
	lockChatsServiceClientMockContains.Unlock()
	return mock.ContainsFunc(ctx, in, opts...)
}

// ContainsCalls gets all the calls that were made to Contains.
// Check the length with:
//     len(mockedChatsServiceClient.ContainsCalls())
func (mock *ChatsServiceClientMock) ContainsCalls() []struct {
	Ctx  context.Context
	In   *Chat
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *Chat
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockContains.RLock()
	calls = mock.calls.Contains
	lockChatsServiceClientMockContains.RUnlock()
	return calls
}

// CreateChannel calls CreateChannelFunc.
func (mock *ChatsServiceClientMock) CreateChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.CreateChannelFunc == nil {
		panic("ChatsServiceClientMock.CreateChannelFunc: method is nil but ChatsServiceClient.CreateChannel was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *Channel
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockCreateChannel.Lock()
	mock.calls.CreateChannel = append(mock.calls.CreateChannel, callInfo)
	lockChatsServiceClientMockCreateChannel.Unlock()
	return mock.CreateChannelFunc(ctx, in, opts...)
}

// CreateChannelCalls gets all the calls that were made to CreateChannel.
// Check the length with:
//     len(mockedChatsServiceClient.CreateChannelCalls())
func (mock *ChatsServiceClientMock) CreateChannelCalls() []struct {
	Ctx  context.Context
	In   *Channel
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *Channel
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockCreateChannel.RLock()
	calls = mock.calls.CreateChannel
	lockChatsServiceClientMockCreateChannel.RUnlock()
	return calls
}

// CreateWorkspace calls CreateWorkspaceFunc.
func (mock *ChatsServiceClientMock) CreateWorkspace(ctx context.Context, in *Workspace, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.CreateWorkspaceFunc == nil {
		panic("ChatsServiceClientMock.CreateWorkspaceFunc: method is nil but ChatsServiceClient.CreateWorkspace was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *Workspace
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockCreateWorkspace.Lock()
	mock.calls.CreateWorkspace = append(mock.calls.CreateWorkspace, callInfo)
	lockChatsServiceClientMockCreateWorkspace.Unlock()
	return mock.CreateWorkspaceFunc(ctx, in, opts...)
}

// CreateWorkspaceCalls gets all the calls that were made to CreateWorkspace.
// Check the length with:
//     len(mockedChatsServiceClient.CreateWorkspaceCalls())
func (mock *ChatsServiceClientMock) CreateWorkspaceCalls() []struct {
	Ctx  context.Context
	In   *Workspace
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *Workspace
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockCreateWorkspace.RLock()
	calls = mock.calls.CreateWorkspace
	lockChatsServiceClientMockCreateWorkspace.RUnlock()
	return calls
}

// DeleteChannel calls DeleteChannelFunc.
func (mock *ChatsServiceClientMock) DeleteChannel(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.DeleteChannelFunc == nil {
		panic("ChatsServiceClientMock.DeleteChannelFunc: method is nil but ChatsServiceClient.DeleteChannel was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockDeleteChannel.Lock()
	mock.calls.DeleteChannel = append(mock.calls.DeleteChannel, callInfo)
	lockChatsServiceClientMockDeleteChannel.Unlock()
	return mock.DeleteChannelFunc(ctx, in, opts...)
}

// DeleteChannelCalls gets all the calls that were made to DeleteChannel.
// Check the length with:
//     len(mockedChatsServiceClient.DeleteChannelCalls())
func (mock *ChatsServiceClientMock) DeleteChannelCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockDeleteChannel.RLock()
	calls = mock.calls.DeleteChannel
	lockChatsServiceClientMockDeleteChannel.RUnlock()
	return calls
}

// DeleteChat calls DeleteChatFunc.
func (mock *ChatsServiceClientMock) DeleteChat(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.DeleteChatFunc == nil {
		panic("ChatsServiceClientMock.DeleteChatFunc: method is nil but ChatsServiceClient.DeleteChat was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockDeleteChat.Lock()
	mock.calls.DeleteChat = append(mock.calls.DeleteChat, callInfo)
	lockChatsServiceClientMockDeleteChat.Unlock()
	return mock.DeleteChatFunc(ctx, in, opts...)
}

// DeleteChatCalls gets all the calls that were made to DeleteChat.
// Check the length with:
//     len(mockedChatsServiceClient.DeleteChatCalls())
func (mock *ChatsServiceClientMock) DeleteChatCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockDeleteChat.RLock()
	calls = mock.calls.DeleteChat
	lockChatsServiceClientMockDeleteChat.RUnlock()
	return calls
}

// DeleteWorkspace calls DeleteWorkspaceFunc.
func (mock *ChatsServiceClientMock) DeleteWorkspace(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.DeleteWorkspaceFunc == nil {
		panic("ChatsServiceClientMock.DeleteWorkspaceFunc: method is nil but ChatsServiceClient.DeleteWorkspace was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockDeleteWorkspace.Lock()
	mock.calls.DeleteWorkspace = append(mock.calls.DeleteWorkspace, callInfo)
	lockChatsServiceClientMockDeleteWorkspace.Unlock()
	return mock.DeleteWorkspaceFunc(ctx, in, opts...)
}

// DeleteWorkspaceCalls gets all the calls that were made to DeleteWorkspace.
// Check the length with:
//     len(mockedChatsServiceClient.DeleteWorkspaceCalls())
func (mock *ChatsServiceClientMock) DeleteWorkspaceCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockDeleteWorkspace.RLock()
	calls = mock.calls.DeleteWorkspace
	lockChatsServiceClientMockDeleteWorkspace.RUnlock()
	return calls
}

// EditChannel calls EditChannelFunc.
func (mock *ChatsServiceClientMock) EditChannel(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.EditChannelFunc == nil {
		panic("ChatsServiceClientMock.EditChannelFunc: method is nil but ChatsServiceClient.EditChannel was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockEditChannel.Lock()
	mock.calls.EditChannel = append(mock.calls.EditChannel, callInfo)
	lockChatsServiceClientMockEditChannel.Unlock()
	return mock.EditChannelFunc(ctx, in, opts...)
}

// EditChannelCalls gets all the calls that were made to EditChannel.
// Check the length with:
//     len(mockedChatsServiceClient.EditChannelCalls())
func (mock *ChatsServiceClientMock) EditChannelCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockEditChannel.RLock()
	calls = mock.calls.EditChannel
	lockChatsServiceClientMockEditChannel.RUnlock()
	return calls
}

// EditWorkspace calls EditWorkspaceFunc.
func (mock *ChatsServiceClientMock) EditWorkspace(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.EditWorkspaceFunc == nil {
		panic("ChatsServiceClientMock.EditWorkspaceFunc: method is nil but ChatsServiceClient.EditWorkspace was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockEditWorkspace.Lock()
	mock.calls.EditWorkspace = append(mock.calls.EditWorkspace, callInfo)
	lockChatsServiceClientMockEditWorkspace.Unlock()
	return mock.EditWorkspaceFunc(ctx, in, opts...)
}

// EditWorkspaceCalls gets all the calls that were made to EditWorkspace.
// Check the length with:
//     len(mockedChatsServiceClient.EditWorkspaceCalls())
func (mock *ChatsServiceClientMock) EditWorkspaceCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockEditWorkspace.RLock()
	calls = mock.calls.EditWorkspace
	lockChatsServiceClientMockEditWorkspace.RUnlock()
	return calls
}

// GetChannelByID calls GetChannelByIDFunc.
func (mock *ChatsServiceClientMock) GetChannelByID(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.GetChannelByIDFunc == nil {
		panic("ChatsServiceClientMock.GetChannelByIDFunc: method is nil but ChatsServiceClient.GetChannelByID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockGetChannelByID.Lock()
	mock.calls.GetChannelByID = append(mock.calls.GetChannelByID, callInfo)
	lockChatsServiceClientMockGetChannelByID.Unlock()
	return mock.GetChannelByIDFunc(ctx, in, opts...)
}

// GetChannelByIDCalls gets all the calls that were made to GetChannelByID.
// Check the length with:
//     len(mockedChatsServiceClient.GetChannelByIDCalls())
func (mock *ChatsServiceClientMock) GetChannelByIDCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockGetChannelByID.RLock()
	calls = mock.calls.GetChannelByID
	lockChatsServiceClientMockGetChannelByID.RUnlock()
	return calls
}

// GetChatByID calls GetChatByIDFunc.
func (mock *ChatsServiceClientMock) GetChatByID(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.GetChatByIDFunc == nil {
		panic("ChatsServiceClientMock.GetChatByIDFunc: method is nil but ChatsServiceClient.GetChatByID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockGetChatByID.Lock()
	mock.calls.GetChatByID = append(mock.calls.GetChatByID, callInfo)
	lockChatsServiceClientMockGetChatByID.Unlock()
	return mock.GetChatByIDFunc(ctx, in, opts...)
}

// GetChatByIDCalls gets all the calls that were made to GetChatByID.
// Check the length with:
//     len(mockedChatsServiceClient.GetChatByIDCalls())
func (mock *ChatsServiceClientMock) GetChatByIDCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockGetChatByID.RLock()
	calls = mock.calls.GetChatByID
	lockChatsServiceClientMockGetChatByID.RUnlock()
	return calls
}

// GetChatsByUserID calls GetChatsByUserIDFunc.
func (mock *ChatsServiceClientMock) GetChatsByUserID(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.GetChatsByUserIDFunc == nil {
		panic("ChatsServiceClientMock.GetChatsByUserIDFunc: method is nil but ChatsServiceClient.GetChatsByUserID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockGetChatsByUserID.Lock()
	mock.calls.GetChatsByUserID = append(mock.calls.GetChatsByUserID, callInfo)
	lockChatsServiceClientMockGetChatsByUserID.Unlock()
	return mock.GetChatsByUserIDFunc(ctx, in, opts...)
}

// GetChatsByUserIDCalls gets all the calls that were made to GetChatsByUserID.
// Check the length with:
//     len(mockedChatsServiceClient.GetChatsByUserIDCalls())
func (mock *ChatsServiceClientMock) GetChatsByUserIDCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockGetChatsByUserID.RLock()
	calls = mock.calls.GetChatsByUserID
	lockChatsServiceClientMockGetChatsByUserID.RUnlock()
	return calls
}

// GetWorkspaceByID calls GetWorkspaceByIDFunc.
func (mock *ChatsServiceClientMock) GetWorkspaceByID(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.GetWorkspaceByIDFunc == nil {
		panic("ChatsServiceClientMock.GetWorkspaceByIDFunc: method is nil but ChatsServiceClient.GetWorkspaceByID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockGetWorkspaceByID.Lock()
	mock.calls.GetWorkspaceByID = append(mock.calls.GetWorkspaceByID, callInfo)
	lockChatsServiceClientMockGetWorkspaceByID.Unlock()
	return mock.GetWorkspaceByIDFunc(ctx, in, opts...)
}

// GetWorkspaceByIDCalls gets all the calls that were made to GetWorkspaceByID.
// Check the length with:
//     len(mockedChatsServiceClient.GetWorkspaceByIDCalls())
func (mock *ChatsServiceClientMock) GetWorkspaceByIDCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockGetWorkspaceByID.RLock()
	calls = mock.calls.GetWorkspaceByID
	lockChatsServiceClientMockGetWorkspaceByID.RUnlock()
	return calls
}

// GetWorkspacesByUserID calls GetWorkspacesByUserIDFunc.
func (mock *ChatsServiceClientMock) GetWorkspacesByUserID(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.GetWorkspacesByUserIDFunc == nil {
		panic("ChatsServiceClientMock.GetWorkspacesByUserIDFunc: method is nil but ChatsServiceClient.GetWorkspacesByUserID was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockGetWorkspacesByUserID.Lock()
	mock.calls.GetWorkspacesByUserID = append(mock.calls.GetWorkspacesByUserID, callInfo)
	lockChatsServiceClientMockGetWorkspacesByUserID.Unlock()
	return mock.GetWorkspacesByUserIDFunc(ctx, in, opts...)
}

// GetWorkspacesByUserIDCalls gets all the calls that were made to GetWorkspacesByUserID.
// Check the length with:
//     len(mockedChatsServiceClient.GetWorkspacesByUserIDCalls())
func (mock *ChatsServiceClientMock) GetWorkspacesByUserIDCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockGetWorkspacesByUserID.RLock()
	calls = mock.calls.GetWorkspacesByUserID
	lockChatsServiceClientMockGetWorkspacesByUserID.RUnlock()
	return calls
}

// LogoutFromChannel calls LogoutFromChannelFunc.
func (mock *ChatsServiceClientMock) LogoutFromChannel(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.LogoutFromChannelFunc == nil {
		panic("ChatsServiceClientMock.LogoutFromChannelFunc: method is nil but ChatsServiceClient.LogoutFromChannel was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockLogoutFromChannel.Lock()
	mock.calls.LogoutFromChannel = append(mock.calls.LogoutFromChannel, callInfo)
	lockChatsServiceClientMockLogoutFromChannel.Unlock()
	return mock.LogoutFromChannelFunc(ctx, in, opts...)
}

// LogoutFromChannelCalls gets all the calls that were made to LogoutFromChannel.
// Check the length with:
//     len(mockedChatsServiceClient.LogoutFromChannelCalls())
func (mock *ChatsServiceClientMock) LogoutFromChannelCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockLogoutFromChannel.RLock()
	calls = mock.calls.LogoutFromChannel
	lockChatsServiceClientMockLogoutFromChannel.RUnlock()
	return calls
}

// LogoutFromWorkspace calls LogoutFromWorkspaceFunc.
func (mock *ChatsServiceClientMock) LogoutFromWorkspace(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*EmptyChats, error) {
	if mock.LogoutFromWorkspaceFunc == nil {
		panic("ChatsServiceClientMock.LogoutFromWorkspaceFunc: method is nil but ChatsServiceClient.LogoutFromWorkspace was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockLogoutFromWorkspace.Lock()
	mock.calls.LogoutFromWorkspace = append(mock.calls.LogoutFromWorkspace, callInfo)
	lockChatsServiceClientMockLogoutFromWorkspace.Unlock()
	return mock.LogoutFromWorkspaceFunc(ctx, in, opts...)
}

// LogoutFromWorkspaceCalls gets all the calls that were made to LogoutFromWorkspace.
// Check the length with:
//     len(mockedChatsServiceClient.LogoutFromWorkspaceCalls())
func (mock *ChatsServiceClientMock) LogoutFromWorkspaceCalls() []struct {
	Ctx  context.Context
	In   *RequestMessage
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RequestMessage
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockLogoutFromWorkspace.RLock()
	calls = mock.calls.LogoutFromWorkspace
	lockChatsServiceClientMockLogoutFromWorkspace.RUnlock()
	return calls
}

// PutChat calls PutChatFunc.
func (mock *ChatsServiceClientMock) PutChat(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*ResponseMessage, error) {
	if mock.PutChatFunc == nil {
		panic("ChatsServiceClientMock.PutChatFunc: method is nil but ChatsServiceClient.PutChat was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *Chat
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	lockChatsServiceClientMockPutChat.Lock()
	mock.calls.PutChat = append(mock.calls.PutChat, callInfo)
	lockChatsServiceClientMockPutChat.Unlock()
	return mock.PutChatFunc(ctx, in, opts...)
}

// PutChatCalls gets all the calls that were made to PutChat.
// Check the length with:
//     len(mockedChatsServiceClient.PutChatCalls())
func (mock *ChatsServiceClientMock) PutChatCalls() []struct {
	Ctx  context.Context
	In   *Chat
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *Chat
		Opts []grpc.CallOption
	}
	lockChatsServiceClientMockPutChat.RLock()
	calls = mock.calls.PutChat
	lockChatsServiceClientMockPutChat.RUnlock()
	return calls
}
