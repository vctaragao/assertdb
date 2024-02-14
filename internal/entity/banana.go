package entity

import "github.com/google/uuid"

type Banana struct {
	ID     uuid.UUID
	Name   string
	Color  string
	Weight int
}

func (b *Banana) SetUUID() {
	b.ID = uuid.New()
}
