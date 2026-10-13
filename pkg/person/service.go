package person

import "go-aether/pkg/entities"

type Service interface {
	InsertPerson(person *entities.Person) (*entities.Person, error)
	FetchPerson() (*[]entities.Person, error)
}
type service struct {
	repository Repository
}

// FetchPerson implements [Service].
func (s *service) FetchPerson() (*[]entities.Person, error) {
	rs, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return rs, nil
}

// InsertPerson implements [Service].
func (s *service) InsertPerson(person *entities.Person) (*entities.Person, error) {

	rs, err := s.repository.Create(person)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
