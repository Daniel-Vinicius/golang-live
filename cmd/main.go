package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Daniel-Vinicius/golang-live/adapter/repository"
	"github.com/Daniel-Vinicius/golang-live/usecases/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    100,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}
