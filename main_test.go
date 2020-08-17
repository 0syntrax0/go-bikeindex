package gobikeindex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
)

// creates test server
func createTestServer(testString string) *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testString)
		}))
	defer srv.Close()
	return srv
}

// ***********************************************************
// TODO: DELETE BEFORE FINALAZING
// ***********************************************************
func appID() string {
	return os.Getenv("BIKEINDEX_APPID")
}
func appSecret() string {
	return os.Getenv("BIKEINDEX_SECRET")
}
