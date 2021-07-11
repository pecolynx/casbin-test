package domain

import "time"

type Pet interface {
	GetName() string
}

type pet struct {
	id        uint
	version   int
	createdAt time.Time
	updatedAt time.Time
	name      string
}

func NewPet(id uint, version int, createdAt time.Time, updatedAt time.Time, name string) Pet {
	return &pet{
		id:        id,
		version:   version,
		createdAt: createdAt,
		updatedAt: updatedAt,
		name:      name,
	}
}

func (m *pet) GetName() string {
	return m.name
}
