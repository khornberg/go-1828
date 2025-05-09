package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"errors"
)

type Word struct {
	Word       string
	Definition string
}

func formatArgument(arg string) (string, error) {
	if len(arg) < 1 {
		return arg, nil
	}
	formatted_arg := strings.ToUpper(string(arg[0])) + strings.ToLower(string(arg[1:]))
	return formatted_arg, nil
}

func find(arg string) (Word, error) {
	db, err := gorm.Open(sqlite.Open("dictionary.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("help!")
	}
	var word Word
	formatedArg, _ := formatArgument(arg)
	db_err := db.First(&word, "word = ?", formatedArg).Error
	if errors.Is(db_err, gorm.ErrRecordNotFound) {
		return word, nil
	}
	return word, nil
}

func main() {
	arg := os.Args[1]
	word, err := find(arg)
	if err != nil {
		fmt.Println("There was an error", err)
	}
	PrintHeader("Webster's 1828 Dictionary\n")
	PrintWord(word.Definition, "cyan")

	done := make(chan bool, 1)
	go thesaurus(arg, done)
	<-done
}
