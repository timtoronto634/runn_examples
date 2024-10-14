package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k1LoW/runn"
)

func TestSample(t *testing.T) {
	// Database connection parameters
	dbUser := "root"
	dbPassword := "example"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "sample"

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	runnDsn := fmt.Sprintf("mysql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	ctx := context.Background()
	opts := []runn.Option{
		runn.T(t),
		runn.Runner("db", runnDsn),
	}
	o, err := runn.Load("*.yml", opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(ctx); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Successfully connected to the MySQL database!")
}
