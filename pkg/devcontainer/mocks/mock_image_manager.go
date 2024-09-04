package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/artmoskvin/hide/pkg/devcontainer"
)

var _ devcontainer.ImageManager = (*MockImageManager)(nil)

type MockImageManager struct {
	mock.Mock
}

func (m *MockImageManager) PullImage(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *MockImageManager) BuildImage(ctx context.Context, workingDir string, config devcontainer.Config) (string, error) {
	args := m.Called(ctx, workingDir, config)
	return args.String(0), args.Error(1)
}