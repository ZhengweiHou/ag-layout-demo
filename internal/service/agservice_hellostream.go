// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	hzw "ag-layout-demo/api/hzw"
	"context"
)

// HelloStreamImpl defines the service implementation for HelloStreamImpl.
type HelloStreamImpl struct {
	// TODO 这里需要根据业务场景，添加依赖
}

// NewHelloStreamImpl creates and returns a new HelloStreamImpl instance.
// @param TODO inject dependencies
// @return *HelloStreamImpl
func NewHelloStreamImpl( /* TODO inject dependencies */ ) *HelloStreamImpl {
	return &HelloStreamImpl{}
}

// NewHelloStreamImplI new HelloStreamImpl instance, and return interface hzw.HelloStream.
// @param TODO inject dependencies
// @return hzw.HelloStream
func NewHelloStreamImplI( /* TODO inject dependencies */ ) hzw.HelloStream {
	return NewHelloStreamImpl( /* TODO inject dependencies */ )
}

// NoStreaming TODO:DESCRIBE
func (c *HelloStreamImpl) NoStreaming(ctx context.Context, req *hzw.Request) (*hzw.Response, error) {
	var resp *hzw.Response

	// TODO 添加业务处理逻辑

	return resp, nil
}

// ClientSideStreaming TODO:DESCRIBE
// ClientStreaming
func (c *HelloStreamImpl) ClientSideStreaming(stream hzw.HelloStream_ClientSideStreamingServer) error {
	// TODO
	return nil
}

// ServerSideStreaming TODO:DESCRIBE
// ServerStreaming
func (c *HelloStreamImpl) ServerSideStreaming(req *hzw.Request, stream hzw.HelloStream_ServerSideStreamingServer) error {
	// TODO
	return nil
}

// BidiSideStreaming TODO:DESCRIBE
// ClientStreaming and ServerStreaming
func (c *HelloStreamImpl) BidiSideStreaming(stream hzw.HelloStream_BidiSideStreamingServer) error {
	// TODO
	return nil
}
