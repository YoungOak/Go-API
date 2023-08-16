package data

import (
	"sync"

	"github.com/YoungOak/GoAPI/internal/car"
)

type Manager interface {
	Add(car.Record) error
	Get(carID string) (car.Record, error)
	List() []car.Record
	Update(car.Record) error
}

type manager struct {
	records map[string]car.Record
	mu      *sync.RWMutex
}

func NewManager() Manager {
	return &manager{
		records: make(map[string]car.Record),
		mu:      &sync.RWMutex{},
	}
}

func (s *manager) Add(record car.Record) error {
	err := record.Validate()
	if err != nil {
		return err
	}

	if s.recordExists(record.ID) {
		return ErrorAlreadyExists{record.ID}
	}

	s.saveRecord(record)
	return nil
}

func (s *manager) Get(recordID string) (car.Record, error) {
	if !s.recordExists(recordID) {
		return car.Record{}, ErrorRecordNotFound{recordID}
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.records[recordID], nil
}

func (s *manager) List() []car.Record {
	var list = make([]car.Record, 0, len(s.records))

	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, record := range s.records {
		list = append(list, record)
	}

	return list
}

func (s *manager) Update(record car.Record) error {
	err := record.Validate()
	if err != nil {
		return err
	}

	if !s.recordExists(record.ID) {
		return ErrorRecordNotFound{record.ID}
	}

	// Update will overwrite whole object
	s.saveRecord(record)
	return nil
}

func (s *manager) recordExists(id string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.records[id]
	return exists
}

func (s *manager) saveRecord(record car.Record) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records[record.ID] = record
}
