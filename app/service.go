package app

import (
	"github.com/acerquetti/ports/domain"
)

type Service interface {
	Create(domain.PortRaw) error
}

type serviceImpl struct {
	repository domain.Repository
}

func NewService(repository domain.Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) Create(portRaw domain.PortRaw) error {
	port, err := domain.NewPortFromRaw(portRaw)
	if err != nil {
		return err
	}

	return s.repository.Save(*port)
}
