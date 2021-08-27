package api

import "github.com/matoous/go-nanoid/v2"

//go:generate mockery --name=IdGenerator --output=../mocks
// IdGenerator it is an interface used for generation of unique id.
type IdGenerator interface {
	NewID() (string, error)
}

// idGenerator is a struct used for generating unique ids
type idGenerator struct{}

// NewIdGenerator returns pointer to new idGenerator structure
func NewIdGenerator() *idGenerator {
	return &idGenerator{}
}

// NewID return new unique id
func (g *idGenerator) NewID() (string, error) {
	return gonanoid.New()
}
