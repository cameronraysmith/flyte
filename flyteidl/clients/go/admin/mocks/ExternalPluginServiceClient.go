// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	service "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
)

// ExternalPluginServiceClient is an autogenerated mock type for the ExternalPluginServiceClient type
type ExternalPluginServiceClient struct {
	mock.Mock
}

type ExternalPluginServiceClient_CreateTask struct {
	*mock.Call
}

func (_m ExternalPluginServiceClient_CreateTask) Return(_a0 *service.TaskCreateResponse, _a1 error) *ExternalPluginServiceClient_CreateTask {
	return &ExternalPluginServiceClient_CreateTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExternalPluginServiceClient) OnCreateTask(ctx context.Context, in *service.TaskCreateRequest, opts ...grpc.CallOption) *ExternalPluginServiceClient_CreateTask {
	c_call := _m.On("CreateTask", ctx, in, opts)
	return &ExternalPluginServiceClient_CreateTask{Call: c_call}
}

func (_m *ExternalPluginServiceClient) OnCreateTaskMatch(matchers ...interface{}) *ExternalPluginServiceClient_CreateTask {
	c_call := _m.On("CreateTask", matchers...)
	return &ExternalPluginServiceClient_CreateTask{Call: c_call}
}

// CreateTask provides a mock function with given fields: ctx, in, opts
func (_m *ExternalPluginServiceClient) CreateTask(ctx context.Context, in *service.TaskCreateRequest, opts ...grpc.CallOption) (*service.TaskCreateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.TaskCreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.TaskCreateRequest, ...grpc.CallOption) *service.TaskCreateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TaskCreateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.TaskCreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ExternalPluginServiceClient_DeleteTask struct {
	*mock.Call
}

func (_m ExternalPluginServiceClient_DeleteTask) Return(_a0 *service.TaskDeleteResponse, _a1 error) *ExternalPluginServiceClient_DeleteTask {
	return &ExternalPluginServiceClient_DeleteTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExternalPluginServiceClient) OnDeleteTask(ctx context.Context, in *service.TaskDeleteRequest, opts ...grpc.CallOption) *ExternalPluginServiceClient_DeleteTask {
	c_call := _m.On("DeleteTask", ctx, in, opts)
	return &ExternalPluginServiceClient_DeleteTask{Call: c_call}
}

func (_m *ExternalPluginServiceClient) OnDeleteTaskMatch(matchers ...interface{}) *ExternalPluginServiceClient_DeleteTask {
	c_call := _m.On("DeleteTask", matchers...)
	return &ExternalPluginServiceClient_DeleteTask{Call: c_call}
}

// DeleteTask provides a mock function with given fields: ctx, in, opts
func (_m *ExternalPluginServiceClient) DeleteTask(ctx context.Context, in *service.TaskDeleteRequest, opts ...grpc.CallOption) (*service.TaskDeleteResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.TaskDeleteResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.TaskDeleteRequest, ...grpc.CallOption) *service.TaskDeleteResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TaskDeleteResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.TaskDeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ExternalPluginServiceClient_GetTask struct {
	*mock.Call
}

func (_m ExternalPluginServiceClient_GetTask) Return(_a0 *service.TaskGetResponse, _a1 error) *ExternalPluginServiceClient_GetTask {
	return &ExternalPluginServiceClient_GetTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExternalPluginServiceClient) OnGetTask(ctx context.Context, in *service.TaskGetRequest, opts ...grpc.CallOption) *ExternalPluginServiceClient_GetTask {
	c_call := _m.On("GetTask", ctx, in, opts)
	return &ExternalPluginServiceClient_GetTask{Call: c_call}
}

func (_m *ExternalPluginServiceClient) OnGetTaskMatch(matchers ...interface{}) *ExternalPluginServiceClient_GetTask {
	c_call := _m.On("GetTask", matchers...)
	return &ExternalPluginServiceClient_GetTask{Call: c_call}
}

// GetTask provides a mock function with given fields: ctx, in, opts
func (_m *ExternalPluginServiceClient) GetTask(ctx context.Context, in *service.TaskGetRequest, opts ...grpc.CallOption) (*service.TaskGetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.TaskGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.TaskGetRequest, ...grpc.CallOption) *service.TaskGetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.TaskGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.TaskGetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}