// store.go
package store

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/sessions"
)

type Store struct {
	Pool  *redis.Pool
	Store *sessions.CookieStore
}

func NewStore(pool *redis.Pool, store *sessions.CookieStore) *Store {
	return &Store{Pool: pool, Store: store}
}

func (s *Store) GetSession(userID uint) (*sessions.Session, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	sessionID, err := redis.String(conn.Do("GET", userID))
	if err != nil {
		return nil, err
	}

	session, err := s.Store.Get(nil, sessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Store) SaveSession(userID uint, session *sessions.Session) error {
	conn := s.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", userID, session.ID)
	if err != nil {
		return err
	}

	return nil
}