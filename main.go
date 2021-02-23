package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/moritamori/gorm-testing/repository"
)

func main() {
	db, err := gorm.Open("postgres", "dbname=gormtesting password=mypassword")
	if err != nil {
		panic(err)
	}
	bookRepository := repository.BookRepositoryImpl{DB: db}

	bookRepository.Create("Go言語の本", "誰か")
	book, err := bookRepository.Get(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
	}
}