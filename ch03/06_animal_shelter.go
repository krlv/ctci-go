package ch03

// Animal represents an animal in the shelter
type Animal interface {
	// ID returns the ID of the animal (the order in which they arrived at the shelter)
	ID() int

	// Name returns the name of the animal
	Name() string

	// IsOlder returns true if the animal spent more time in the shelter than the given animal
	IsOlder(Animal) bool
}

// animal is a base struct for Dog and Cat
type animal struct {
	id   int
	name string
}

// ID returns the ID of the animal (the order in which they arrived at the shelter)
func (a animal) ID() int {
	return a.id
}

// Name returns the name of the animal
func (a animal) Name() string {
	return a.name
}

// IsOlder returns true if the animal spent more time in the shelter than the given animal
func (a animal) IsOlder(o Animal) bool {
	return a.id < o.ID()
}

// Dog represents a dog in the shelter
type Dog struct{ animal }

// Cat represents a cat in the shelter
type Cat struct{ animal }

// Shelter represents an animal shelter
//
// It holds two queues: one for cats and one for dogs
// Each animal has an ID that represents the order in which they arrived at the shelter.
//
// The shelter operates using a "first in, first out" approach.
// People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
// or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type)
type Shelter struct {
	id   int
	cats Queue[Cat]
	dogs Queue[Dog]
}

// NewShelter returns new Shelter
func NewShelter() *Shelter {
	return &Shelter{
		cats: Queue[Cat]{},
		dogs: Queue[Dog]{},
	}
}

// Enqueue adds an animal to the shelter
func (s *Shelter) Enqueue(a Animal) {
	s.id++

	switch a.(type) {
	case Cat:
		cat := a.(Cat)
		cat.id = s.id
		s.cats.Enqueue(cat)
	case Dog:
		dog := a.(Dog)
		dog.id = s.id
		s.dogs.Enqueue(dog)
	}
}

// DequeueAny returns the oldest animal from the shelter
func (s *Shelter) DequeueAny() (Animal, error) {
	cat, err := s.cats.Peek()
	if err != nil {
		return s.dogs.Dequeue()
	}

	dog, err := s.dogs.Peek()
	if err != nil {
		return s.cats.Dequeue()
	}

	if cat.IsOlder(dog) {
		return s.cats.Dequeue()
	}

	return s.dogs.Dequeue()
}

// DequeueCat returns the oldest cat from the shelter
func (s *Shelter) DequeueCat() (Cat, error) {
	return s.cats.Dequeue()
}

// DequeueDog returns the oldest dog from the shelter
func (s *Shelter) DequeueDog() (Dog, error) {
	return s.dogs.Dequeue()
}
