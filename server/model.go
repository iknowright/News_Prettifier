package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type news struct {
	ID    int     `json:"id"`
	Link  string  `json:"link"`
	ReturnLink string `json:"returnlink"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1", p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", p.Name, p.Price, p.ID)
	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	return err
}

func (p *product) createProduct(db *sql.DB) error {
	// postgres doesn't return the last inserted ID so this is the workaround
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)
	return err
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (n *news) getNews(db *sql.DB) error {
	return db.QueryRow("SELECT link, returnlink FROM news WHERE id=$1", n.ID).Scan(&n.Link, &n.ReturnLink)
}

func (n *news) updateNews(db *sql.DB) error {
	_, err := db.Exec("UPDATE news SET link=$1, returnlink=$2 WHERE id=$3", n.Link, n.ReturnLink, n.ID)
	return err
}

func (n *news) deleteNews(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM news WHERE id=$1", n.ID)
	return err
}

func (n *news) createNews(db *sql.DB) error {
	// postgres doesn't return the last inserted ID so this is the workaround
	err := db.QueryRow(
		"INSERT INTO news(link, returnlink) VALUES($1, $2) RETURNING id",
		n.Link, n.ReturnLink).Scan(&n.ID)
	return err
}

func getMultipleNews(db *sql.DB, start, count int) ([]news, error) {
	rows, err := db.Query("SELECT id, link, returnlink FROM news LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	multipleNews := []news{}

	for rows.Next() {
		var n news
		if err := rows.Scan(&n.ID, &n.Link, &n.ReturnLink); err != nil {
			return nil, err
		}
		multipleNews = append(multipleNews, n)
	}

	return multipleNews, nil
}