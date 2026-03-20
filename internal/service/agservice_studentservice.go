// TODO 这里需要根据业务场景，自行修改实现

package service

import (
	student "ag-layout-demo/api/student"
	"context"
)

// StudentServiceImpl defines the service implementation for StudentServiceImpl.
type StudentServiceImpl struct {
	// TODO 这里需要根据业务场景，添加依赖
}

// NewStudentServiceImpl creates and returns a new StudentServiceImpl instance.
// @param TODO inject dependencies
// @return *StudentServiceImpl
func NewStudentServiceImpl() *StudentServiceImpl {
	return &StudentServiceImpl{}
}

// CreateStudent TODO:DESCRIBE
func (c *StudentServiceImpl) CreateStudent(ctx context.Context, req *student.CreateStudentRequest) (*student.Student, error) {
	var resp *student.Student
	resp = &student.Student{}
	// TODO 添加业务处理逻辑

	return resp, nil
}

// GetStudent TODO:DESCRIBE
func (c *StudentServiceImpl) GetStudent(ctx context.Context, req *student.GetStudentRequest) (*student.Student, error) {
	var resp *student.Student
	resp = &student.Student{}
	// TODO 添加业务处理逻辑

	return resp, nil
}
