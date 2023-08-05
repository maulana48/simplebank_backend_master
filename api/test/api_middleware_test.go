package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	api "github.com/maulana48/backend_master_class/simplebank/api"
	"github.com/maulana48/backend_master_class/simplebank/token"
	"github.com/stretchr/testify/require"
)

// add bearer token to request
func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(api.AuthorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name string
		// setup the request header authorization with token maker
		setupAuth func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		// check the response from middleware
		chechResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T,
				request *http.Request,
				tokenMaker token.Maker,
			) {
				addAuthorization(t, request, tokenMaker, api.AllowedType, "user", time.Minute)
			},
			chechResponse: func(t *testing.T,
				recorder *httptest.ResponseRecorder,
			) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "No Authorization",
			setupAuth: func(t *testing.T,
				request *http.Request,
				tokenMaker token.Maker,
			) {
			},
			chechResponse: func(t *testing.T,
				recorder *httptest.ResponseRecorder,
			) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Unsupported Authorization",
			setupAuth: func(t *testing.T,
				request *http.Request,
				tokenMaker token.Maker,
			) {
				addAuthorization(t, request, tokenMaker, "unsupported", "user", time.Minute)
			},
			chechResponse: func(t *testing.T,
				recorder *httptest.ResponseRecorder,
			) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Invalid Authorization Format",
			setupAuth: func(t *testing.T,
				request *http.Request,
				tokenMaker token.Maker,
			) {
				addAuthorization(t, request, tokenMaker, "", "user", time.Minute)
			},
			chechResponse: func(t *testing.T,
				recorder *httptest.ResponseRecorder,
			) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Expired Token",
			setupAuth: func(t *testing.T,
				request *http.Request,
				tokenMaker token.Maker,
			) {
				addAuthorization(t, request, tokenMaker, api.AllowedType, "user", -time.Minute)
			},
			chechResponse: func(t *testing.T,
				recorder *httptest.ResponseRecorder,
			) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server, err := NewTestServer(t, nil)
			if err != nil {
				fmt.Println(err.Error())
			}

			authPath := "/auth"
			server.Router.GET(
				authPath,
				api.AuthMiddleware(server.TokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			// record the http request call
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.chechResponse(t, recorder)
		})
	}
}
