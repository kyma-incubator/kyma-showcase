package main

import "github.com/matoous/go-nanoid/v2"

type idGenerator struct{}

func NewIdGenerator() *idGenerator {
	return &idGenerator{}
}

func (g *idGenerator) NewID() (string, error) {
	return gonanoid.New()
}
