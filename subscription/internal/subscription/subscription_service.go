package subscription

import "context"

type SubscriptionService struct {
    Repo *SubscriptionRepository
}

func NewSubscriptionService(repo *SubscriptionRepository) *SubscriptionService {
    return &SubscriptionService{Repo: repo}
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context, sub *Subscription) (*Subscription, error) {
    return s.Repo.CreateSubscription(ctx, sub)
}

func (s *SubscriptionService) GetSubscription(ctx context.Context, id int) (*Subscription, error) {
    return s.Repo.GetSubscription(ctx, id)
}

func (s *SubscriptionService) UpdateSubscription(ctx context.Context, sub *Subscription) (*Subscription, error) {
    return s.Repo.UpdateSubscription(ctx, sub)
}

func (s *SubscriptionService) DeleteSubscription(ctx context.Context, id int) error {
    return s.Repo.DeleteSubscription(ctx, id)
}

func (s *SubscriptionService) GetAllSubscriptions(ctx context.Context) ([]Subscription, error) {
    return s.Repo.GetAllSubscriptions(ctx)
}
