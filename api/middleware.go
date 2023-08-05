package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maulana48/backend_master_class/simplebank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationType       = "bearer"
	authorizationPayloadKey = "authorization"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
