package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/maulana48/backend_master_class/simplebank/db/sqlc"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	// config := util.Config{
	// 	TokenSymmetricKey:   util.RandomString(32),
	// 	AccessTokenDuration: time.Minute,
	// }

	server := NewServer(store)
	// require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
