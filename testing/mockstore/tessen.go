package mockstore

import (
	"context"

	v2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/backend/store"
)

// CreateOrUpdateTessenConfig ...
func (s *MockStore) CreateOrUpdateTessenConfig(ctx context.Context, config *v2.TessenConfig) error {
	args := s.Called(ctx, config)
	return args.Error(0)
}

// GetTessenConfig ...
func (s *MockStore) GetTessenConfig(ctx context.Context) (*v2.TessenConfig, error) {
	args := s.Called(ctx)
	return args.Get(0).(*v2.TessenConfig), args.Error(1)
}

// GetTessenConfigWatcher ...
func (s *MockStore) GetTessenConfigWatcher(ctx context.Context) <-chan store.WatchEventTessenConfig {
	args := s.Called(ctx)
	return args.Get(0).(<-chan store.WatchEventTessenConfig)
}
