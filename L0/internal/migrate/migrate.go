package migrate

import (
	"github.com/go-pg/pg/orm"
	"l0/internal/model"
	"l0/internal/repository"
)

func CreateSchema(db *repository.Repository) error {
	models := []interface{}{
		(*model.Order)(nil),
	}

	for _, model := range models {
		op := orm.CreateTableOptions{}
		err := db.DB.Model(model).CreateTable(&op)
		if err != nil {
			return err
		}
	}
	return nil
}
