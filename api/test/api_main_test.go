package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maulana48/backend_master_class/simplebank/api"
	db "github.com/maulana48/backend_master_class/simplebank/db/sqlc"
	"github.com/maulana48/backend_master_class/simplebank/util"
)

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

	os.Exit(m.Run())
}
