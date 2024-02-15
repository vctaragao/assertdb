package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Banana struct {
	bun.BaseModel `bun:"table:bananas,alias:b"`
	ID            uuid.UUID `bun:"type:uuid,pk"`
	Name          string    `bun:"type:text"`
	Color         string    `bun:"type:text"`
	Weight        int       `bun:"type:integer,nullzero"`
}

func (b *Banana) SetUUID() {
	b.ID = uuid.New()
}
