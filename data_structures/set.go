package data_structures

type Setter[T comparable] interface {
	Add(element T)
	Remove(element T)
	IsSubset(other Setter[T]) bool
}

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable](elements map[T]struct{}) *Set[T] {
	return &Set[T]{elements: elements}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) IsSubset(other Setter[T]) bool {
	otherSet, ok := other.(*Set[T])
	if !ok {
		return false
	}

	for element := range s.elements {
		if _, exists := otherSet.elements[element]; !exists {
			return false
		}
	}
	return true
}
