package app

import (
	"database/sql"
)

type news struct {
	ID    int     `json:"id"`
	Link  string  `json:"link"`
	ReturnLink string `json:"returnlink"`
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

type account struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email string `json:"email"`
}

func (u *account) getNews(db *sql.DB) error {
	return db.QueryRow("SELECT username, email FROM account WHERE username=$1", u.Username).Scan(&u.Password, &u.Email)
}

func (u *account) updateNews(db *sql.DB) error {
	_, err := db.Exec("UPDATE account SET password=$1, email=$2 WHERE username=$3", u.Password, u.Email, u.Username)
	return err
}

func (u *account) deleteNews(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM account WHERE username=$1", u.Username)
	return err
}

func (u *account) createNews(db *sql.DB) error {
	// postgres doesn't return the last inserted Username so this is the workaround
	err := db.QueryRow(
		"INSERT INTO account(username, password, email) VALUES($1, $2) RETURNING username",
		u.Password, u.Email).Scan(&u.Username)
	return err
}