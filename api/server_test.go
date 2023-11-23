package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	host = "http://localhost"
)

type ActualResponse struct {
	Handler http.Handler
	Method  string
	Path    string
}

func (ac ActualResponse) PerformRequest() *httptest.ResponseRecorder {
	req, _ := http.NewRequest(ac.Method, ac.Path, nil)
	w := httptest.NewRecorder()
	ac.Handler.ServeHTTP(w, req)
	return w
}

type ApiTest struct {
	suite.Suite
	router *gin.Engine
}

func TestApiTest(t *testing.T) {
	suite.Run(t, &ApiTest{})
}

func (ap *ApiTest) SetupSuite() {
	server := Server{}
	ap.router = server.SetupRouter()
}

func (ap *ApiTest) TestPing() {
	ac := ActualResponse{
		Handler: ap.router,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("%s%s/ping", host, port),
	}.PerformRequest()

	ap.Equal(http.StatusOK, ac.Code)
	ap.Equal("pong", ac.Body.String())
}
