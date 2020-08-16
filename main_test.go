package gobikeindex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
