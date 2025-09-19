// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	hzw "ag-layout-demo/api/hzw"
	"context"
)

// HzwHello2Impl defines the service implementation for HzwHello2Impl.
type HzwHello2Impl struct {
	// TODO 这里需要根据业务场景，添加依赖
}

// NewHzwHello2Impl creates and returns a new HzwHello2Impl instance.
// @param TODO inject dependencies
// @return *HzwHello2Impl
func NewHzwHello2Impl( /* TODO inject dependencies */ ) *HzwHello2Impl {
	return &HzwHello2Impl{}
}

// NewHzwHello2ImplI new HzwHello2Impl instance, and return interface hzw.HzwHello2.
// @param TODO inject dependencies
// @return hzw.HzwHello2
func NewHzwHello2ImplI( /* TODO inject dependencies */ ) hzw.HzwHello2 {
	return NewHzwHello2Impl( /* TODO inject dependencies */ )
}

// HelloHzw2 TODO:DESCRIBE
func (c *HzwHello2Impl) HelloHzw2(ctx context.Context, req *hzw.HzwRequest2) (*hzw.HzwReply2, error) {
	var resp *hzw.HzwReply2

	// TODO 添加业务处理逻辑

	return resp, nil
}

// HelloHzwStream2 TODO:DESCRIBE
// ClientStreaming and ServerStreaming
func (c *HzwHello2Impl) HelloHzwStream2(stream hzw.HzwHello2_HelloHzwStream2Server) error {
	// TODO
	return nil
}
