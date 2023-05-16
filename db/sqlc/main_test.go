package db

import (
	"database/sql"
	"jwt-paseto-token-in-golang/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	testDB, err := sql.Open(config.DBDriver, config.DBSourceLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// fmt.Println(conn)

	testQueries = New(testDB)

	os.Exit(m.Run())
}
