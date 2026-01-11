package main

import (
	"errors"
	"studyProject/model"
	"sync"
)

type Storage interface {
	Insert(e *model.Employee)
	Get(id int) (model.Employee, error)
	Update(id int, e model.Employee)
	Delete(id int)
	GetAll() []model.Employee
}
type MemoryStorage struct {
	counter int
	data    map[int]model.Employee
	sync.Mutex
}

func newMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]model.Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *model.Employee) {
	s.Lock()

	e.ID = s.counter
	s.data[e.ID] = *e
	s.counter++
	s.Unlock()
}

func (s *MemoryStorage) Update(id int, e model.Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (model.Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee, ok := s.data[id]

	if !ok {
		return employee, errors.New("employee not found")
	}

	return employee, nil
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
}

func (s *MemoryStorage) GetAll() []model.Employee {
	s.Lock()
	defer s.Unlock()

	employees := make([]model.Employee, len(s.data))

	for _, key := range s.data {
		employees = append(employees, key)
	}

	return employees
}
