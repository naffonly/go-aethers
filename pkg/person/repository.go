package person

import (
	"go-aether/pkg/entities"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonRepository interface {
	Create(ctx fiber.Ctx, person *entities.Person)
}

type personRepository struct {
	dbPool *pgxpool.Pool
}

// Create implements [PersonRepository].
func (p *personRepository) Create(ctx fiber.Ctx, person *entities.Person) {
	panic("unimplemented")
}

func newRepository(db *pgxpool.Pool) PersonRepository {
	return &personRepository{
		dbPool: db,
	}
}
