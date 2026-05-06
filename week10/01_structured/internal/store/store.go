package store

import "sync"

// User is the domain type. Exported so server package can use it.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Store is a thread-safe in-memory user store.
//
// Node.js bridge: the event loop serialises all access so you never think about
// this. In Go, goroutines are truly concurrent — RWMutex lets multiple readers
// proceed simultaneously while writes are exclusive.
type Store struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func New() *Store {
	return &Store{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *Store) Create(name string, age int) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := User{ID: s.nextID, Name: name, Age: age}
	s.users[s.nextID] = u
	s.nextID++
	return u
}

func (s *Store) Get(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *Store) Update(id int, name string, age int) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[id]
	if !ok {
		return User{}, false
	}
	u.Name = name
	u.Age = age
	s.users[id] = u
	return u, true
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}
