package singleton

type SingletonCounter struct {
	count int
}

var instance *SingletonCounter

func GetInstance() *SingletonCounter {
	if instance == nil {
		instance = &SingletonCounter{}
	}
	return instance
}

func (s *SingletonCounter) AddOne() int {
	s.count++
	return s.count
}
