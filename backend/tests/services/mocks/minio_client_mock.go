package mocks_test

import (
	"github.com/stretchr/testify/mock"
)

type MockMinioClient struct {
	mock.Mock
}

func (m *MockMinioClient) Upload(bucketName, objectName string, content []byte) error {
	args := m.Called(bucketName, objectName, content)
	return args.Error(0)
}

func (m *MockMinioClient) GetObject(bucketName, objectName string) ([]byte, error) {
	args := m.Called(bucketName, objectName)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockMinioClient) MakeBucket(bucketName string) error {
	args := m.Called(bucketName)
	return args.Error(0)
}
