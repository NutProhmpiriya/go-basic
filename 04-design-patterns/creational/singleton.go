// Singleton Pattern ensures a class has only one instance and provides a global point
// of access to that instance.
//
// Use cases:
// - Database connections
// - Configuration settings
// - Logging services

package creational

import (
	"sync"
)

type Singleton struct {
	count int
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance returns the single instance of Singleton
// Thread-safe due to sync.Once
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{count: 0}
	})
	return instance
}

// IncrementCount increases the counter
func (s *Singleton) IncrementCount() int {
	s.count++
	return s.count
}

// GetCount returns the current count
func (s *Singleton) GetCount() int {
	return s.count
}
