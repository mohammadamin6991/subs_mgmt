package plan

import (
    "context"
)

type PlanService struct {
    PlanRepository *PlanRepository
}

func NewPlanService(planRepository *PlanRepository) *PlanService {
    return &PlanService{
        PlanRepository: planRepository,
    }
}

func (s *PlanService) CreatePlan(ctx context.Context, p *Plan) (*Plan, error) {
    return s.PlanRepository.CreatePlan(ctx, p)
}

func (s *PlanService) GetPlan(ctx context.Context, id int) (*Plan, error) {
    return s.PlanRepository.GetPlan(ctx, id)
}

func (s *PlanService) UpdatePlan(ctx context.Context, p *Plan) (*Plan, error) {
    return s.PlanRepository.UpdatePlan(ctx, p)
}

func (s *PlanService) DeletePlan(ctx context.Context, id int) error {
    return s.PlanRepository.DeletePlan(ctx, id)
}

func (s *PlanService) GetAllPlans(ctx context.Context) ([]Plan, error) {
    return s.PlanRepository.GetAllPlans(ctx)
}
