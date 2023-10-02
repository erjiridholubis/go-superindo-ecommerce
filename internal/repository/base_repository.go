package repository

import "database/sql"

type postgreRepository struct {
	ConnDB *sql.DB
}

func NewPostgreRepository(connDB *sql.DB) PostgreRepository {
	return &postgreRepository{connDB}
}

type PostgreRepository interface {

}