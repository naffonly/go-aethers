package person

import (
	"go-aether/pkg/entities"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var (
	ctx fiber.Ctx
)

type Repository interface {
	Create(person *entities.Person) (*entities.Person, error)
	GetAll() (*[]entities.Person, error)
}

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// GetAll implements [PersonRepository].
func (repo *repository) GetAll() (*[]entities.Person, error) {

	var person []entities.Person
	page := 12
	offset := (page - 1) * page
	err := repo.db.Limit(page).Offset(offset).Find(&person).Error

	if err != nil {
		return nil, err
	}

	return &person, nil
}

// Create implements [PersonRepository].
func (repo *repository) Create(person *entities.Person) (*entities.Person, error) {

	err := repo.db.WithContext(ctx).Create(&person).Error

	if err != nil {
		return nil, err
	}
	return person, nil

}
