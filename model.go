package main

import (
	"database/sql"
	"errors"
)

func getProducts(db *sql.DB, lang string) ([]product, error) {

	if lang != "en" && lang != "fr" {
		err := errors.New("Invalid Language")
		return nil, err
	}

	query := "SELECT id, code FROM SourceEn"
	if lang == "fr" {
		query = "SELECT id, code FROM SourceFr"
	}

	rows, err := db.Query(
		query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.Id, &p.Code); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func deleteProductsEn(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM SourceEn")

	if err != nil {
		return err
	}

	return nil
}

func deleteProductsFr(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM SourceFr")

	if err != nil {
		return err
	}

	return nil
}

func (p *productEn) create(db *sql.DB) error {
	query := getInsertQueryEn(p)
	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (p *productFr) create(db *sql.DB) error {
	query := getInsertQueryFr(p)
	_, err := db.Exec(query)

	if err != nil {
		return errors.New(err.Error() + "\r\n" + query)
	}

	return nil
}
