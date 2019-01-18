// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import container "github.com/docker/docker/api/types/container"
import context "context"
import filters "github.com/docker/docker/api/types/filters"
import io "io"
import mock "github.com/stretchr/testify/mock"
import network "github.com/docker/docker/api/types/network"
import time "time"
import types "github.com/docker/docker/api/types"

// ContainerAPIClient is an autogenerated mock type for the ContainerAPIClient type
type ContainerAPIClient struct {
	mock.Mock
}

// ContainerAttach provides a mock function with given fields: ctx, _a1, options
func (_m *ContainerAPIClient) ContainerAttach(ctx context.Context, _a1 string, options types.ContainerAttachOptions) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerAttachOptions) types.HijackedResponse); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerAttachOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerCommit provides a mock function with given fields: ctx, _a1, options
func (_m *ContainerAPIClient) ContainerCommit(ctx context.Context, _a1 string, options types.ContainerCommitOptions) (types.IDResponse, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.IDResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerCommitOptions) types.IDResponse); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.IDResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerCommitOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerCreate provides a mock function with given fields: ctx, config, hostConfig, networkingConfig, containerName
func (_m *ContainerAPIClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	ret := _m.Called(ctx, config, hostConfig, networkingConfig, containerName)

	var r0 container.ContainerCreateCreatedBody
	if rf, ok := ret.Get(0).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) container.ContainerCreateCreatedBody); ok {
		r0 = rf(ctx, config, hostConfig, networkingConfig, containerName)
	} else {
		r0 = ret.Get(0).(container.ContainerCreateCreatedBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) error); ok {
		r1 = rf(ctx, config, hostConfig, networkingConfig, containerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerDiff provides a mock function with given fields: ctx, _a1
func (_m *ContainerAPIClient) ContainerDiff(ctx context.Context, _a1 string) ([]types.ContainerChange, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []types.ContainerChange
	if rf, ok := ret.Get(0).(func(context.Context, string) []types.ContainerChange); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ContainerChange)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecAttach provides a mock function with given fields: ctx, execID, config
func (_m *ContainerAPIClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error) {
	ret := _m.Called(ctx, execID, config)

	var r0 types.HijackedResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecConfig) types.HijackedResponse); ok {
		r0 = rf(ctx, execID, config)
	} else {
		r0 = ret.Get(0).(types.HijackedResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecConfig) error); ok {
		r1 = rf(ctx, execID, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecCreate provides a mock function with given fields: ctx, _a1, config
func (_m *ContainerAPIClient) ContainerExecCreate(ctx context.Context, _a1 string, config types.ExecConfig) (types.IDResponse, error) {
	ret := _m.Called(ctx, _a1, config)

	var r0 types.IDResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecConfig) types.IDResponse); ok {
		r0 = rf(ctx, _a1, config)
	} else {
		r0 = ret.Get(0).(types.IDResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ExecConfig) error); ok {
		r1 = rf(ctx, _a1, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecInspect provides a mock function with given fields: ctx, execID
func (_m *ContainerAPIClient) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error) {
	ret := _m.Called(ctx, execID)

	var r0 types.ContainerExecInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerExecInspect); ok {
		r0 = rf(ctx, execID)
	} else {
		r0 = ret.Get(0).(types.ContainerExecInspect)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, execID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerExecResize provides a mock function with given fields: ctx, execID, options
func (_m *ContainerAPIClient) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error {
	ret := _m.Called(ctx, execID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ResizeOptions) error); ok {
		r0 = rf(ctx, execID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerExecStart provides a mock function with given fields: ctx, execID, config
func (_m *ContainerAPIClient) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error {
	ret := _m.Called(ctx, execID, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ExecStartCheck) error); ok {
		r0 = rf(ctx, execID, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerExport provides a mock function with given fields: ctx, _a1
func (_m *ContainerAPIClient) ContainerExport(ctx context.Context, _a1 string) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerInspect provides a mock function with given fields: ctx, _a1
func (_m *ContainerAPIClient) ContainerInspect(ctx context.Context, _a1 string) (types.ContainerJSON, error) {
	ret := _m.Called(ctx, _a1)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ContainerJSON); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerInspectWithRaw provides a mock function with given fields: ctx, _a1, getSize
func (_m *ContainerAPIClient) ContainerInspectWithRaw(ctx context.Context, _a1 string, getSize bool) (types.ContainerJSON, []byte, error) {
	ret := _m.Called(ctx, _a1, getSize)

	var r0 types.ContainerJSON
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) types.ContainerJSON); ok {
		r0 = rf(ctx, _a1, getSize)
	} else {
		r0 = ret.Get(0).(types.ContainerJSON)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) []byte); ok {
		r1 = rf(ctx, _a1, getSize)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, bool) error); ok {
		r2 = rf(ctx, _a1, getSize)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ContainerKill provides a mock function with given fields: ctx, _a1, signal
func (_m *ContainerAPIClient) ContainerKill(ctx context.Context, _a1 string, signal string) error {
	ret := _m.Called(ctx, _a1, signal)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, _a1, signal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerList provides a mock function with given fields: ctx, options
func (_m *ContainerAPIClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.Container
	if rf, ok := ret.Get(0).(func(context.Context, types.ContainerListOptions) []types.Container); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ContainerListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerLogs provides a mock function with given fields: ctx, _a1, options
func (_m *ContainerAPIClient) ContainerLogs(ctx context.Context, _a1 string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerPause provides a mock function with given fields: ctx, _a1
func (_m *ContainerAPIClient) ContainerPause(ctx context.Context, _a1 string) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerRemove provides a mock function with given fields: ctx, _a1, options
func (_m *ContainerAPIClient) ContainerRemove(ctx context.Context, _a1 string, options types.ContainerRemoveOptions) error {
	ret := _m.Called(ctx, _a1, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerRemoveOptions) error); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerRename provides a mock function with given fields: ctx, _a1, newContainerName
func (_m *ContainerAPIClient) ContainerRename(ctx context.Context, _a1 string, newContainerName string) error {
	ret := _m.Called(ctx, _a1, newContainerName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, _a1, newContainerName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerResize provides a mock function with given fields: ctx, _a1, options
func (_m *ContainerAPIClient) ContainerResize(ctx context.Context, _a1 string, options types.ResizeOptions) error {
	ret := _m.Called(ctx, _a1, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ResizeOptions) error); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerRestart provides a mock function with given fields: ctx, _a1, timeout
func (_m *ContainerAPIClient) ContainerRestart(ctx context.Context, _a1 string, timeout *time.Duration) error {
	ret := _m.Called(ctx, _a1, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *time.Duration) error); ok {
		r0 = rf(ctx, _a1, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerStart provides a mock function with given fields: ctx, _a1, options
func (_m *ContainerAPIClient) ContainerStart(ctx context.Context, _a1 string, options types.ContainerStartOptions) error {
	ret := _m.Called(ctx, _a1, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerStartOptions) error); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerStatPath provides a mock function with given fields: ctx, _a1, path
func (_m *ContainerAPIClient) ContainerStatPath(ctx context.Context, _a1 string, path string) (types.ContainerPathStat, error) {
	ret := _m.Called(ctx, _a1, path)

	var r0 types.ContainerPathStat
	if rf, ok := ret.Get(0).(func(context.Context, string, string) types.ContainerPathStat); ok {
		r0 = rf(ctx, _a1, path)
	} else {
		r0 = ret.Get(0).(types.ContainerPathStat)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, _a1, path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerStats provides a mock function with given fields: ctx, _a1, stream
func (_m *ContainerAPIClient) ContainerStats(ctx context.Context, _a1 string, stream bool) (types.ContainerStats, error) {
	ret := _m.Called(ctx, _a1, stream)

	var r0 types.ContainerStats
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) types.ContainerStats); ok {
		r0 = rf(ctx, _a1, stream)
	} else {
		r0 = ret.Get(0).(types.ContainerStats)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, _a1, stream)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerStop provides a mock function with given fields: ctx, _a1, timeout
func (_m *ContainerAPIClient) ContainerStop(ctx context.Context, _a1 string, timeout *time.Duration) error {
	ret := _m.Called(ctx, _a1, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *time.Duration) error); ok {
		r0 = rf(ctx, _a1, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerTop provides a mock function with given fields: ctx, _a1, arguments
func (_m *ContainerAPIClient) ContainerTop(ctx context.Context, _a1 string, arguments []string) (types.ContainerProcessList, error) {
	ret := _m.Called(ctx, _a1, arguments)

	var r0 types.ContainerProcessList
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) types.ContainerProcessList); ok {
		r0 = rf(ctx, _a1, arguments)
	} else {
		r0 = ret.Get(0).(types.ContainerProcessList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, _a1, arguments)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerUnpause provides a mock function with given fields: ctx, _a1
func (_m *ContainerAPIClient) ContainerUnpause(ctx context.Context, _a1 string) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContainerUpdate provides a mock function with given fields: ctx, _a1, updateConfig
func (_m *ContainerAPIClient) ContainerUpdate(ctx context.Context, _a1 string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error) {
	ret := _m.Called(ctx, _a1, updateConfig)

	var r0 container.ContainerUpdateOKBody
	if rf, ok := ret.Get(0).(func(context.Context, string, container.UpdateConfig) container.ContainerUpdateOKBody); ok {
		r0 = rf(ctx, _a1, updateConfig)
	} else {
		r0 = ret.Get(0).(container.ContainerUpdateOKBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, container.UpdateConfig) error); ok {
		r1 = rf(ctx, _a1, updateConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerWait provides a mock function with given fields: ctx, _a1
func (_m *ContainerAPIClient) ContainerWait(ctx context.Context, _a1 string) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainersPrune provides a mock function with given fields: ctx, pruneFilters
func (_m *ContainerAPIClient) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error) {
	ret := _m.Called(ctx, pruneFilters)

	var r0 types.ContainersPruneReport
	if rf, ok := ret.Get(0).(func(context.Context, filters.Args) types.ContainersPruneReport); ok {
		r0 = rf(ctx, pruneFilters)
	} else {
		r0 = ret.Get(0).(types.ContainersPruneReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, filters.Args) error); ok {
		r1 = rf(ctx, pruneFilters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyFromContainer provides a mock function with given fields: ctx, _a1, srcPath
func (_m *ContainerAPIClient) CopyFromContainer(ctx context.Context, _a1 string, srcPath string) (io.ReadCloser, types.ContainerPathStat, error) {
	ret := _m.Called(ctx, _a1, srcPath)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, string) io.ReadCloser); ok {
		r0 = rf(ctx, _a1, srcPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 types.ContainerPathStat
	if rf, ok := ret.Get(1).(func(context.Context, string, string) types.ContainerPathStat); ok {
		r1 = rf(ctx, _a1, srcPath)
	} else {
		r1 = ret.Get(1).(types.ContainerPathStat)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, _a1, srcPath)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CopyToContainer provides a mock function with given fields: ctx, _a1, path, content, options
func (_m *ContainerAPIClient) CopyToContainer(ctx context.Context, _a1 string, path string, content io.Reader, options types.CopyToContainerOptions) error {
	ret := _m.Called(ctx, _a1, path, content, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader, types.CopyToContainerOptions) error); ok {
		r0 = rf(ctx, _a1, path, content, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
