package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p product) get(db *sql.DB) error {
	panic("implement me")
}

func (p product) update(db *sql.DB) error {
	panic("implement me")
}

func (p product) delete(db *sql.DB) error {
	panic("implement me")
}

func (p product) create(db *sql.DB) error {
	panic("implement me")
}

func getProducts(db *sql.DB, start, count int)([]product, error){
	return nil, errors.New("Not implemented")
}
