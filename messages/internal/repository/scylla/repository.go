package scylla

import "github.com/alserov/chatter/messages/internal/repository"

var _ repository.Repository = &Scylla{}

func NewRepository() *Scylla {
	return &Scylla{}
}

type Scylla struct {
}
