package main

import (
	"context"
	"fmt"
	"github.com/Threpio/vatdb-ingest/vatdb"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

const url string = "https://data.vatsim.net/v3/vatsim-data.json"

func main() {

	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		DB_USER = "theoa"
	}
	DB_PASSWORD := os.Getenv("DB_USER")
	if DB_PASSWORD == "" {
		DB_PASSWORD = ""
	}
	DB_HOST := os.Getenv("DB_USER")
	if DB_HOST == "" {
		DB_HOST = "localhost"
	}
	DB_PORT := os.Getenv("DB_USER")
	if DB_PORT == "" {
		DB_PORT = "5432"
	}
	DB_NAME := os.Getenv("DB_USER")
	if DB_NAME == "" {
		DB_NAME = "vatdb"
	}

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s ", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	queries := vatdb.New(conn)

	// Main subroutine loop
	for {
		go fetchVatsimData(ctx, queries)
		time.Sleep(30 * time.Second)
	}
}

// fetchVatsimData fetches the data from the VATSIM API and stores it in the database.
func fetchVatsimData(ctx context.Context, queries *vatdb.Queries) {
	fmt.Printf("fetchVatsimData: %s - ", time.Now())
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	_, err = queries.CreateDataInstance(ctx, bodyBytes)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		return
	}
	fmt.Println("Data fetched from the VATSIM API.")
}
