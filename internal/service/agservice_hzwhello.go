// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	hzw "ag-layout-demo/api/hzw"
	"context"
)

// HzwHelloImpl defines the service implementation for HzwHelloImpl.
type HzwHelloImpl struct {
	// TODO 这里需要根据业务场景，添加依赖
}

// NewHzwHelloImpl creates and returns a new HzwHelloImpl instance.
// @param TODO inject dependencies
// @return *HzwHelloImpl
func NewHzwHelloImpl( /* TODO inject dependencies */ ) *HzwHelloImpl {
	return &HzwHelloImpl{}
}

// NewHzwHelloImplI new HzwHelloImpl instance, and return interface hzw.HzwHello.
// @param TODO inject dependencies
// @return hzw.HzwHello
func NewHzwHelloImplI( /* TODO inject dependencies */ ) hzw.HzwHello {
	return NewHzwHelloImpl( /* TODO inject dependencies */ )
}

// HelloHzw TODO:DESCRIBE
func (c *HzwHelloImpl) HelloHzw(ctx context.Context, req *hzw.HzwRequest) (*hzw.HzwReply, error) {
	var resp *hzw.HzwReply

	// TODO 添加业务处理逻辑

	return resp, nil
}

// HelloHzwStream TODO:DESCRIBE
// ClientStreaming and ServerStreaming
func (c *HzwHelloImpl) HelloHzwStream(stream hzw.HzwHello_HelloHzwStreamServer) error {
	// TODO
	return nil
}
