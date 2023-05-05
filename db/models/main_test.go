package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

// globals for all tests
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// this will come from a config
	var err error
	testDB, err = sql.Open("postgres", "postgresql://genomeMaster@genomefileflipper.postgres.database.azure.com:Hackathon23Team1Cat@genomefileflipper:5432/fileflipperdb?sslmode=disable")
	if err != nil {
		log.Fatalln("cannot connect to db", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
