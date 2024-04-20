package main

import (
	"context"
	"fmt"
	"github.com/Threpio/vatdb-ingest/vatdb"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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
		DB_PASSWORD = "theoa"
	}
	DB_HOST := os.Getenv("DB_USER")
	if DB_HOST == "" {
		DB_HOST = "127.0.0.1"
	}
	DB_PORT := os.Getenv("DB_USER")
	if DB_PORT == "" {
		DB_PORT = "5432"
	}
	DB_NAME := os.Getenv("DB_USER")
	if DB_NAME == "" {
		DB_NAME = "vatdb"
	}
	SLEEP_TIME := os.Getenv("SLEEP_TIME")
	if SLEEP_TIME == "" {
		SLEEP_TIME = strconv.Itoa(30)
	}
	s, err := strconv.Atoi(SLEEP_TIME)
	if err != nil {
		panic(err)
	}

	// Let the Postgress container wake up
	// This line just reduces logs
	
	time.Sleep(10 * time.Second)

	connString := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable ", "postgres", DB_USER, DB_PASSWORD, DB_PORT, DB_NAME)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}

	schema := filepath.Join("schema.sql")

	c, err := ioutil.ReadFile(schema)
	if err != nil {
		panic(err)
	}
	sql := string(c)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		fmt.Println("Error implementing schema")
		panic(err)
	}

	defer conn.Close(ctx)
	queries := vatdb.New(conn)

	// Main subroutine loop
	for {
		go fetchVatsimData(ctx, queries)
		time.Sleep(time.Duration(s) * time.Second)
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
