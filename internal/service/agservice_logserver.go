// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	logs "ag-layout-demo/api/logs"
	"context"
)

// LogServerImpl defines the service implementation for LogServerImpl.
type LogServerImpl struct {
	// TODO 这里需要根据业务场景，添加依赖
}

// NewLogServerImpl creates and returns a new LogServerImpl instance.
// @param TODO inject dependencies
// @return *LogServerImpl
func NewLogServerImpl( /* TODO inject dependencies */ ) *LogServerImpl {
	return &LogServerImpl{}
}

// NewLogServerImplI new LogServerImpl instance, and return interface logs.LogServer.
// @param TODO inject dependencies
// @return logs.LogServer
func NewLogServerImplI( /* TODO inject dependencies */ ) logs.LogServer {
	return NewLogServerImpl( /* TODO inject dependencies */ )
}

// BatchLog TODO:DESCRIBE
func (c *LogServerImpl) BatchLog(ctx context.Context, req *logs.BatchLogRequest) (*logs.BatchLogResponse, error) {
	var resp *logs.BatchLogResponse

	// TODO 添加业务处理逻辑

	return resp, nil
}

// StreamLog TODO:DESCRIBE
// ClientStreaming
func (c *LogServerImpl) StreamLog(stream logs.LogServer_StreamLogServer) error {
	// TODO
	return nil
}
