package internal

import (
	"context"

	"github.com/vctaragao/assertdb/internal/entity"
)

type CreateBananaService struct {
	repo BananaRepository
}

func NewCreateBananaService(repo BananaRepository) *CreateBananaService {
	return &CreateBananaService{repo: repo}
}

func (s *CreateBananaService) Execute(ctx context.Context, b *entity.Banana) error {
	b.SetUUID()
	return s.repo.CreateBanana(ctx, b)
}
