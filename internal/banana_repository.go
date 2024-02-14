package internal

import (
	"context"

	"github.com/vctaragao/assertdb/internal/entity"
)

type BananaRepository interface {
	CreateBanana(ctx context.Context, banana *entity.Banana) error
}
