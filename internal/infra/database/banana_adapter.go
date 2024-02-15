package database

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/vctaragao/assertdb/internal/entity"
)

type BananaAdapter struct {
	db bun.IDB
}

func NewBananaAdapter(db bun.IDB) *BananaAdapter {
	return &BananaAdapter{
		db: db,
	}
}

func (a *BananaAdapter) CreateBanana(ctx context.Context, b *entity.Banana) error {
	_, err := a.db.NewInsert().Model(b).Exec(ctx)
	return err
}
