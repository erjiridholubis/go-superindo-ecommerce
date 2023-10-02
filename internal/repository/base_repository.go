package repository

import "database/sql"

type postgreProductRepository struct {
	ConnDB *sql.DB
}

func NewPostgreProductRepository(connDB *sql.DB) ProductRepository {
	return &postgreProductRepository{connDB}
}

type ProductRepository interface {

}