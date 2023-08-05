package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	api "github.com/maulana48/backend_master_class/simplebank/api"
	db "github.com/maulana48/backend_master_class/simplebank/db/sqlc"
	"github.com/maulana48/backend_master_class/simplebank/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func NewTestServer(t *testing.T, store db.Store) (*api.Server, error) {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := api.NewServer(config, store)
	// require.NoError(t, err)

	return server, err
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSourceLocal)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// fmt.Println(conn)

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
