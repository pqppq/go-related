package repository

import (
	"io/ioutil"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createTableBooks string = `
		CREATE TABLE books(
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL
		);`
	initialRecords           = 5
	insertInitialData string = `
		INSERT INTO books(id, title) VALUES (1, "To Kill a Mockingbird by Harper Lee");
		INSERT INTO books(id, title) VALUES (2, "1984 by George Orwell");
		INSERT INTO books(id, title) VALUES (3, "The Great Gatsby by F. Scott Fitzgerald");
		INSERT INTO books(id, title) VALUES (4, "Pride and Prejudice by Jane Austen");
		INSERT INTO books(id, title) VALUES (5, "The Lord of the Rings by J.R.R. Tolkien");

	`
)

func getRepo(t *testing.T) *BookRepo {
	t.Helper()
	tf, err := ioutil.TempFile("", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	return NewBookRepo(tf.Name())
}

func InitializeRepo(t *testing.T, repo *BookRepo) error {
	t.Helper()
	_, err := repo.db.Exec(createTableBooks)
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(insertInitialData)
	if err != nil {
		return err
	}
	return nil
}

func TestRepo(t *testing.T) {
	repo := getRepo(t)
	err := InitializeRepo(t, repo)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("get all books", func(t *testing.T) {
		books, err := repo.GetAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(books) != initialRecords {
			t.Errorf("Expected len(books) %d, but got %d instead.", initialRecords, len(books))
		}
	})

	t.Run("get book by id", func(t *testing.T) {
		id := 1
		title := "To Kill a Mockingbird by Harper Lee"
		book, err := repo.Get(id)
		if err != nil {
			t.Fatal(err)
		}
		if book.Title != title {
			t.Errorf("Expected books.Title %s, but got %s instead.", title, book.Title)
		}
	})

	t.Run("insert book", func(t *testing.T) {
		title := "The Diary of a Young Girl by Anne Frank"
		err := repo.Create(title)
		if err != nil {
			t.Fatal(err)
		}
		books, err := repo.GetAll()
		if len(books) != initialRecords+1 {
			t.Errorf("Expected records %d, but got %d instead.", initialRecords+1, len(books))
		}
		book, err := repo.Get(initialRecords + 1)
		if book.Title != title {
			t.Errorf("Expected books.Title %s, but got %s instead.", title, book.Title)
		}
	})

	t.Run("update book", func(t *testing.T) {
		id := 1
		title := "The Hobbit by J.R.R. Tolkien"
		err := repo.Update(id, title)
		if err != nil {
			t.Fatal(err)
		}
		book, err := repo.Get(id)
		if book.Title != title {
			t.Errorf("Expected books.Title %s, but got %s instead.", title, book.Title)
		}
	})

	t.Run("delete book", func(t *testing.T) {
		id := 1
		err := repo.Delete(id)
		if err != nil {
			t.Fatal(err)
		}
		books, err := repo.GetAll()
		if len(books) != initialRecords {
			t.Errorf("Expected len(books) %d, but got %d instead.", initialRecords, len(books))
		}
	})

}
