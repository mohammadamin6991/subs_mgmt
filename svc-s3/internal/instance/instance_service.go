package instance

import (
    "context"
)

type InstanceService struct {
    InstanceRepository *InstanceRepository
}

func NewInstanceService(instanceRepository *InstanceRepository) *InstanceService {
    return &InstanceService{
        InstanceRepository: instanceRepository,
    }
}

func (s *InstanceService) CreateInstance(ctx context.Context, i *Instance) (*Instance, error) {
    return s.InstanceRepository.CreateInstance(ctx, i)
}

func (s *InstanceService) GetInstance(ctx context.Context, id int) (*Instance, error) {
    return s.InstanceRepository.GetInstance(ctx, id)
}

func (s *InstanceService) UpdateInstance(ctx context.Context, i *Instance) (*Instance, error) {
    return s.InstanceRepository.UpdateInstance(ctx, i)
}

func (s *InstanceService) DeleteInstance(ctx context.Context, id int) error {
    return s.InstanceRepository.DeleteInstance(ctx, id)
}

func (s *InstanceService) GetAllInstances(ctx context.Context) ([]Instance, error) {
    return s.InstanceRepository.GetAllInstances(ctx)
}
