package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Word struct {
	word       string
	definition string
}

func main() {
	db, err := gorm.Open(sqlite.Open("dictionary.db"), &gorm.Config{})
	var word Word
	db.First(&word)
	fmt.Println(word)
}
