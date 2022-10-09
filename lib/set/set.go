package set

type Set[T comparable] struct {
	cache map[T]any
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		cache: make(map[T]any),
	}
}

func (s *Set[T]) Add(args ...T) {
	for _, arg := range args {
		s.cache[arg] = struct{}{}
	}
}

func (s *Set[T]) IsExist(arg T) bool {
	if _, ok := s.cache[arg]; ok {
		return true
	}
	return false
}
