package scylla

import "github.com/alserov/chatter/chat/internal/repository"

var _ repository.Repository = &Scylla{}

func NewRepository() *Scylla {
	return &Scylla{}
}

type Scylla struct {
}
