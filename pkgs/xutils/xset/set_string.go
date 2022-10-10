package xset

type SetString struct {
	m map[string]struct{}
}

func NewSetString() *SetString {
	return &SetString{
		m: make(map[string]struct{}),
	}
}

func (s *SetString) Insert(item interface{}) bool {
	i, ok := item.(string)
	if !ok {
		return false
	}
	if s.Contains(item) {
		return false
	}
	s.m[i] = struct{}{}
	return true
}

func (s *SetString) Contains(item interface{}) bool {
	if i, ok := item.(string); ok {
		if _, ok := s.m[i]; ok {
			return true
		}
	}
	return false
}

func (s *SetString) Size() int {
	return len(s.m)
}
