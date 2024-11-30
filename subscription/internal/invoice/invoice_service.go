package invoice

import "context"

type InvoiceService struct {
    Repo *InvoiceRepository
}

func NewInvoiceService(repo *InvoiceRepository) *InvoiceService {
    return &InvoiceService{Repo: repo}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, inv *Invoice) (*Invoice, error) {
    return s.Repo.CreateInvoice(ctx, inv)
}

func (s *InvoiceService) GetInvoice(ctx context.Context, id int) (*Invoice, error) {
    return s.Repo.GetInvoice(ctx, id)
}

func (s *InvoiceService) UpdateInvoice(ctx context.Context, inv *Invoice) (*Invoice, error) {
    return s.Repo.UpdateInvoice(ctx, inv)
}

func (s *InvoiceService) DeleteInvoice(ctx context.Context, id int) error {
    return s.Repo.DeleteInvoice(ctx, id)
}

func (s *InvoiceService) GetAllInvoices(ctx context.Context) ([]Invoice, error) {
    return s.Repo.GetAllInvoices(ctx)
}
