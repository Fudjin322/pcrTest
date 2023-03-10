package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Book struct {
	id       int
	title    string
	author   string
	numPages int
	rating   float64
}

func main() {
	const connStr = "postgres://admin:admin@localhost:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	//query := `
	//
	//  create table if not exists bmg_request(
	//      id integer primary key GENERATED BY DEFAULT AS IDENTITY,
	//      iin text,
	//      phone_number text
	//
	//  );
	//`
	query := `
	
	  create table if not exists bmg_response(
	      id integer primary key GENERATED BY DEFAULT AS IDENTITY,
	      iin text,
	      status text
	      
	  );
	`
	query = `
	
	  create table if not exists pcr_request(
	      id integer primary key GENERATED BY DEFAULT AS IDENTITY,
	      iin text
	      
	  );
	`
	query = `
	
	  create table if not exists pcr_response(
	      id integer primary key GENERATED BY DEFAULT AS IDENTITY,
	      iin text, 
	      credit_rating text,
	      risk_class text,
	      risk_code text
	      
	  );
	`

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	//	query = `
	//   insert into books(title, author, num_pages, rating)
	//   values ($1, $2, $3, $4)
	//`
	//
	//	data := [][]any{
	//		{"The Catcher in the Rye", "J.D. Salinger", 277, 3.8},
	//		{"The Fellowship of the Ring", "J.R.R. Tolkin", 398, 4.36},
	//		{"The Giver", "Lois Lowry", 208, 4.13},
	//		{"The Da Vinci Code", "Dan Brown", 489, 3.84},
	//		{"The Alchemist", "Paulo Coelho", 197, 3.86},
	//	}
	//
	//for _, vals := range data {
	//	res, err := db.Exec(query, vals...) // (1)
	//	if err != nil {
	//		panic(err)
	//	}
	//	bookID, err := res.LastInsertId() // (2)
	//	fmt.Printf("added new book: id=%d, error=%v\n", bookID, err)
	//}

	/*query := "select * from books"
	rows, err := db.Query(query) // (1)
	if err != nil {
		panic(err)
	}
	defer rows.Close() // (2)

	for rows.Next() { // (3)
		var book Book
		err := rows.Scan(&book.id, &book.title, &book.author, &book.numPages, &book.rating) // (4)
		if err != nil {
			panic(err)
		}
		fmt.Println(book)
	}

	if err := rows.Err(); err != nil { // (5)
		panic(err)
	}*/

}
