package repository

import (
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
	"l0/internal/model"
)

type Repository struct {
	DB     *pg.DB
	config *config
}

func NewRepository(path string) (*Repository, error) {
	config, err := newConfig(path)
	if err != nil {
		return nil, err
	}
	db := pg.Connect(&pg.Options{
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})
	logrus.Info("Database is up!")
	if err != nil {
		logrus.Info("Table create!")
	}
	return &Repository{db, config}, nil
}

func (r *Repository) Close() error {
	err := r.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddOrder(order model.Order) error {
	// вставка нового ордера в базу
	if _, err := r.DB.Model(&order).Insert(); err != nil {
		return err
	}
	return nil
}
