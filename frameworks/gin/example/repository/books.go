package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pqppq/go-related/frameworks/gin/example/model"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(dbFile string) *BookRepo {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	return &BookRepo{
		db: db,
	}
}

func (r *BookRepo) Get(id int) (model.Book, error) {
	stmt, err := r.db.Prepare("SELECT * from books WHERE id = ?")
	if err != nil {
		return model.Book{}, err
	}

	var book model.Book
	err = stmt.QueryRow(id).Scan(&book.Id, &book.Title)

	return book, err
}

func (r *BookRepo) GetAll() ([]model.Book, error) {
	stmt, err := r.db.Prepare("SELECT * from books")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Title); err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepo) Create(title string) error {
	stmt, err := r.db.Prepare("INSERT INTO books(title) VALUES (?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(title)
	return err
}

func (r *BookRepo) Update(id int, title string) error {
	stmt, err := r.db.Prepare("UPDATE books SET title = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(title, id)
	return err

}

func (r *BookRepo) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM books WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
