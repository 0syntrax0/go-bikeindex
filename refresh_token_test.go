package gobikeindex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRefreshToken(t *testing.T) {
	// create test server
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testSearchCountBikes)
		}))
	defer srv.Close()

	bi := NewBikeIndexClient(appID(), appSecret())
	excp, err := bi.RefreshToken()
	if err != nil {
		t.Fatalf("Error refreshing token: %+v", err)
	}
	if excp != nil {
		t.Fatalf("Exception refreshin token: %+v", excp)
	}
}
