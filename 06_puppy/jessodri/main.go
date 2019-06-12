package main

import "sync"

// Puppy - Implement a `Puppy` struct containing `ID`, `Breed`, `Colour`, `Value`
type Puppy struct {
	ID    string
	Breed string
	Color string
	Value float64
}

// Storer - Create `Storer` interface with [crud](https://en.wikipedia.org/wiki/Create,\_read,\_update_and_delete) methods for `Puppy`
type Storer interface {
	CreatePuppy(p Puppy) *Puppy
	ReadPuppy(id string) *Puppy
	UpdatePuppy(id string)
	DeletePuppy(id string)
}

// MapStore - Write a `MapStore` implementation of `Storer` backed by a `map`
type MapStore map[string]*Puppy

// SyncStore - Write a `SyncStore` implementation of `Storer` backed by a [sync.Map](https://golang.org/pkg/sync/#Map)
type SyncStore *sync.Map

func (m MapStore) CreatePuppy(p Puppy) *Puppy {

	return &Puppy{}
}

func (m MapStore) ReadPuppy(id string) *Puppy {

	return &Puppy{}
}

func (m MapStore) UpdatePuppy(id string) {

}

func (m MapStore) DeletePuppy(id string) {

}


func main() {

}
