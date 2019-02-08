package userrepo

import "pavan/gohttpexamples/sample4/domain"

type Reader interface {
	GetAll() ([]*domain.User, error)
	GetByID(ID string) (*domain.User, error)
}

type Writer interface {
	Create(*domain.User) (string, error)
	Update(*domain.User) error
	Delete(ID string) error
	//Archive(*domain.User) error
}

type Repository interface {
	Reader
	Writer
}
