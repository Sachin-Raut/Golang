package bookrepository

import (
	"database/sql"
	"05-Golang-REST-API-Mike-M/03-books-list-refactored/model"
)

//BookRepository is
type BookRepository struct {}


//GetBooks is
func (b BookRepository) GetBooks(db *sql.DB, book model.Book, books []model.Book) ([]model.Book, error){
	rows, err := db.Query("select * from booksTable")

	if err != nil {
		return []model.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}



//GetBook is
func (b BookRepository) GetBook(db *sql.DB, book model.Book, id int) (model.Book, error) {
	rows := db.QueryRow("select * from booksTable where id = $1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil {
		return model.Book{}, err
	}

	return book, err
}

//AddBook is
func (b BookRepository) AddBook(db *sql.DB, book model.Book) (int, error) {
	err := db.QueryRow("insert into booksTable(title, author, year) values($1, $2, $3) returning id;", book.Title, book.Author, book.Year).Scan(&book.ID)
	if err != nil {
		return 0, err
	}
	return book.ID, err
}

//UpdateBook is
func (b BookRepository) UpdateBook(db *sql.DB, book model.Book) (int, error) {

	result, err := db.Exec("update booksTable set title = $1, author = $2, year = $3 where id = $4 returning id", &book.Title, &book.Author, &book.Year, &book.ID)
	if err != nil {
		return 0, err
	}
	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	//lets return the number of rowsUpdated
	return int(rowsUpdated), err
}

//RemoveBook is
func (b BookRepository) RemoveBook(db *sql.DB, id int) (int, error) {
	
	result, err := db.Exec("delete from booksTable where id = $1", id)
	
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), err
}