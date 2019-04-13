package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/cache"
)

type Storage struct {
	listMut   sync.Mutex
	latestMut sync.Mutex
	cache.Cache
}

func NewStorage() *Storage {
	return &Storage{
		Cache: cache.NewMemoryCache(),
	}
}

func (s *Storage) list(

	mod string,
	putter func(mod string) ([]string, error),
) ([]string, error) {
	key := fmt.Sprintf("list-%s", mod)
	s.listMut.Lock()
	defer s.listMut.Unlock()
	lstIface := s.Get(key)
	if lstIface != nil {
		lst, ok := lstIface.([]string)
		if ok {
			return lst, nil
		}
	}
	lst, err := putter(mod)
	if err != nil {
		return nil, err
	}
	s.Put(key, lst, 5*time.Second)
	return lst, nil
}

func (s *Storage) latest(
	mod string,
	putter func(mod string) (string, error),
) (string, error) {
	key := fmt.Sprintf("latest-%s", mod)
	s.latestMut.Lock()
	defer s.latestMut.Unlock()
	latIface := s.Get(key)
	if latIface != nil {
		latest, ok := latIface.(string)
		if ok {
			return latest, nil
		}
	}
	lat, err := putter(mod)
	if err != nil {
		return "", err
	}
	s.Put(key, lat, 5*time.Second)
	return lat, nil
}
