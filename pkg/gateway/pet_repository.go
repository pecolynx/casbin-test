package gateway

import (
	"context"
	"time"

	"gorm.io/gorm"

	casbinquery "github.com/pecolynx/casbin-query"
	"github.com/pecolynx/casbin-test/pkg/domain"
)

type petEntity struct {
	ID        uint
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (e *petEntity) TableName() string {
	return "pet"
}

func (e *petEntity) toModel() domain.Pet {
	return domain.NewPet(e.ID, e.Version, e.CreatedAt, e.UpdatedAt, e.Name)
}

type PetRepository interface {
	FindPets(ctx context.Context, name string) ([]domain.Pet, error)
}

type petRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) PetRepository {
	return &petRepository{
		db: db,
	}
}

func (r *petRepository) FindPets(ctx context.Context, name string) ([]domain.Pet, error) {
	objectColumnName := "name"
	subQuery, err := casbinquery.QueryObject(r.db, objectColumnName, "user_"+name, "read")
	if err != nil {
		return nil, err
	}

	petEntities := []petEntity{}
	if result := r.db.Model(&petEntity{}).
		Joins("inner join (?) as t3 on `pet`.`name`= t3."+objectColumnName, subQuery).
		Scan(&petEntities); result.Error != nil {
		return nil, result.Error
	}

	petModels := make([]domain.Pet, len(petEntities))
	for i, e := range petEntities {
		petModels[i] = e.toModel()
	}

	return petModels, nil
}
