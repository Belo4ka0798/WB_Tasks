package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"l0/internal/model"
	"l0/internal/repository"
)

type Server struct {
	db     *repository.Repository
	config *config
	routes *chi.Mux
	sc     stan.Conn
	sub    stan.Subscription
	cache  map[string]model.Order
}

func NewServer(path string) (*Server, error) {
	db, err := repository.NewRepository(path)
	if err != nil {
		return nil, err
	}
	conf, err := newConfig(path)
	if err != nil {
		return nil, err
	}
	server := &Server{
		db:     db,
		config: conf,
		cache:  make(map[string]model.Order),
		routes: chi.NewRouter(),
	}
	return server, nil
}

func (s *Server) Up() error {
	if err := s.setCache(); err != nil {
		logrus.Warning(err)
		return err
	}
	if err := s.connectToStream(); err != nil {
		logrus.Warning(err)
		return err
	}
	s.startRouting()
	return nil
}

func (s *Server) Down() {
	logrus.Info("Server is down")
	s.db.Close()
	s.sub.Unsubscribe()
	s.sc.Close()
}

func (s *Server) connectToStream() error {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}
	sub, err := sc.Subscribe(s.config.SubscribeSubject, s.handleRequest)
	if err != nil {
		return err
	}
	s.sc, s.sub = sc, sub
	return nil
}

func (s *Server) handleRequest(m *stan.Msg) {
	data := model.Order{}
	err := json.Unmarshal(m.Data, &data)
	if err != nil {
		return
	}
	if ok := s.addToCache(data); ok {
		logrus.Info("Add to cache")
		s.db.AddOrder(data)
	}
}

func (s *Server) addToCache(data model.Order) bool {
	_, ok := s.cache[data.OrderUid]
	if ok {
		return false
	}
	s.cache[data.OrderUid] = data
	for key := range s.cache {
		fmt.Printf("%s ", key)
	}
	fmt.Println()
	return true
}

func (s *Server) setCache() error {
	orders := make([]model.Order, 0)
	err := s.db.DB.Model(&orders).Select()
	if err != nil {
		return err
	}
	for _, order := range orders {
		s.cache[order.OrderUid] = order
	}
	return nil
}
